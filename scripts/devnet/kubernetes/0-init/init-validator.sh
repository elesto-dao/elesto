#!/usr/local/bin/bash

# TODO param
AMOUNT=${AMOUNT:-10000000} # Maybe check it's below minimum amount?
INIT_FROM=${INIT_FROM:-'genesis-0'}
SEED_FROM=${SEED_FROM:-'seed-0'}

# Wait for pods
kubectl wait pod -l app.kubernetes.io/name=genesis --for condition=Ready
kubectl wait pod -l app.kubernetes.io/name=seed --for condition=Ready

# Init node
if [[ ! -f /home/config/node_key.json ]]; then
  echo "Init ${MONIKER} on ${CHAIN_ID}"
  elestod init ${MONIKER} --home=/home/ --chain-id=${CHAIN_ID}
  sed -i 's|keyring-backend = ".*"|keyring-backend = "test"|' /home/config/client.toml
fi

# Add ${INIT_FROM} private key
if [ "$(elestod keys list --output json --home /home/ | jq -r --arg INIT_FROM "${INIT_FROM}" '.[] | select(.name == $INIT_FROM).name')" != "${INIT_FROM}" ]; then
  echo "Add ${INIT_FROM} key"
  kubectl get secret ${INIT_FROM} --output json | jq -r '.data.MNEMONIC' | base64 -d | elestod keys add ${INIT_FROM} --home=/home/ --output json --recover
fi

# Create a key if it doesn't exist
if [ "$(elestod keys list --output json --home /home/ | jq -r --arg MONIKER "${MONIKER}" '.[] | select(.name == $MONIKER).name')" != "${MONIKER}" ]; then
  echo "Create ${MONIKER} key"
  elestod keys add ${MONIKER} --home=/home/ --output json | jq -r '.mnemonic' > /home/config/${MONIKER}_mnemonic
fi

FROM_ADDRESS=$(elestod keys show ${INIT_FROM} --home=/home/ --address)
TO_ADDRESS=$(elestod keys show ${MONIKER} --home=/home/ --address)

# Send token from ${INIT_FROM} to ${MONIKER}
if [ "$(elestod query bank balances ${TO_ADDRESS} --home=/home/ --chain-id=${CHAIN_ID} --node tcp://${INIT_FROM%-*}:26657 --output json | jq -r '.balances[] | select(.denom="stake").amount')" != "${AMOUNT}" ]; then
  echo "Send ${AMOUNT}stake from ${INIT_FROM} at address ${FROM_ADDRESS} to ${MONIKER} at address ${TO_ADDRESS}"
  elestod tx bank send ${FROM_ADDRESS} ${TO_ADDRESS} ${AMOUNT}stake --home=/home/ --chain-id=${CHAIN_ID} --from ${FROM_ADDRESS} --node tcp://${INIT_FROM%-*}:26657 --output json --yes

  # Wait for transaction to complete & check
  while [ "$(elestod query bank balances ${TO_ADDRESS} --home=/home/ --chain-id=${CHAIN_ID} --node tcp://${INIT_FROM%-*}:26657 --output json | jq -r '.balances[] | select(.denom="stake").amount')" != "${AMOUNT}" ]; do
    echo "tx bank send transaction still not processed (amount not found)"
    sleep 1;
  done
fi

# Create the validator
VALIDATOR_PUBKEY=$(elestod tendermint show-validator --home=/home/)
if [ "$(elestod query staking validators --home=/home/ --chain-id=${CHAIN_ID} --limit 100 --node tcp://${INIT_FROM%-*}:26657 --output json | jq -r --argjson PUBKEY "${VALIDATOR_PUBKEY}" '.validators[] | select(.consensus_pubkey==$PUBKEY).tokens')" != "${AMOUNT}" ]; then
  echo "Create validator with ${AMOUNT}stake from${MONIKER} at address ${FROM_ADDRESS} with pubkey ${VALIDATOR_PUBKEY}"
  elestod tx staking create-validator --home=/home/ --amount=${AMOUNT}stake --from=${MONIKER} --pubkey=${VALIDATOR_PUBKEY} --moniker=${MONIKER} --chain-id=${CHAIN_ID} --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="${AMOUNT}" --node tcp://${INIT_FROM%-*}:26657 --output json --yes

  # Wait for transaction to complete & check
  while [ "$(elestod query staking validators --home=/home/ --chain-id=${CHAIN_ID} --limit 100 --node tcp://${INIT_FROM%-*}:26657 --output json | jq -r --argjson PUBKEY "${VALIDATOR_PUBKEY}" '.validators[] | select(.consensus_pubkey==$PUBKEY).tokens')" != "${AMOUNT}" ]; do
    echo "staking create-validator transaction still not processed (bonded amount not found)"
    sleep 1;
  done
fi

# Update configmap & secret
echo "Update configmap and secret"
INIT_FROM_ID=$(wget -O - -o /dev/null http://${INIT_FROM%-*}:26657/status | jq -r '.result.node_info.id' | tr -d '\n')
SEED_FROM_ID=$(wget -O - -o /dev/null http://${SEED_FROM%-*}:26657/status | jq -r '.result.node_info.id' | tr -d '\n')
cat /dev/null > /home/config/validators-env
echo "P2P_PERSISTENT_PEERS=${INIT_FROM_ID}@${INIT_FROM%-*}:26656" >> /home/config/validators-env
echo "P2P_PRIVATE_PEER_IDS=${INIT_FROM_ID}" >> /home/config/validators-env
echo "P2P_SEEDS=${SEED_FROM_ID}@${SEED_FROM%-*}:26656" >> /home/config/validators-env
kubectl create configmap ${MONIKER%-*}s-env --dry-run=client --output=yaml --from-env-file=/home/config/validators-env | kubectl apply --force=true --filename=-
kubectl create secret generic ${MONIKER} --dry-run=client --output=yaml --from-file MNEMONIC=/home/config/${MONIKER}_mnemonic | kubectl apply --force=true --filename=-
