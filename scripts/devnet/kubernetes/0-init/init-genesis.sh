#!/usr/local/bin/bash

GENESIS_AMOUNT=${GENESIS_AMOUNT:-100000000}

# Init node
if [[ ! -f /home/config/node_key.json ]]; then
  echo "Init ${MONIKER} on ${CHAIN_ID}"
  elestod init ${MONIKER} --home=/home/ --chain-id=${CHAIN_ID}
  sed -i 's|keyring-backend = ".*"|keyring-backend = "test"|' /home/config/client.toml
fi

# Create key if it does not exists
if [ "$(elestod keys list --output json --home /home/ | jq -r --arg MONIKER "${MONIKER}" '.[] | select(.name == $MONIKER).name')" != "${MONIKER}" ]; then
  echo "Create ${MONIKER} key"
  elestod keys add ${MONIKER} --home=/home/ --output json | jq -r '.mnemonic' > /home/config/${MONIKER}_mnemonic
  elestod add-genesis-account $(elestod keys show ${MONIKER} --home=/home/ -a) ${GENESIS_AMOUNT}stake --home=/home/
fi

# Create regulator
if [ "$(elestod keys list --output json --home /home/ | jq -r '.[] | select(.name == "regulator").name')" != "regulator" ]; then
  echo "Create regulator key"
  echo ${REGULATOR_KEY} | elestod keys add regulator --home=/home/ --output json --recover
  elestod add-genesis-account $(elestod keys show regulator --home=/home/ -a) 20000000stake --home=/home/
fi

# Create emti (e-money token issuer)
if [ "$(elestod keys list --output json --home /home/ | jq -r '.[] | select(.name == "emti").name')" != "emti" ]; then
  echo "Create emti key"
  echo ${EMTI_KEY} | elestod keys add emti --home=/home/ --output json --recover
  elestod add-genesis-account $(elestod keys show emti --home=/home/ -a) 20000000stake --home=/home/
fi

# Create arti (asset-referenced token issuer)
if [ "$(elestod keys list --output json --home /home/ | jq -r '.[] | select(.name == "arti").name')" != "arti" ]; then
  echo "Create arti key"
  echo ${ARTI_KEY} | elestod keys add arti --home=/home/ --output json --recover
  elestod add-genesis-account $(elestod keys show arti --home=/home/ -a) 20000000stake --home=/home/
fi

if [ "$(jq -r '.app_state.genutil.gen_txs | length' /home/config/genesis.json)" == "0" ]; then
  # Generate a transaction & save it to the Genesis file
  echo "Generate transaction and collect to Genesis on chain ${CHAIN_ID}"
  elestod gentx ${MONIKER} 70000000stake --home=/home/ --chain-id=${CHAIN_ID}
  elestod collect-gentxs --home=/home/
fi

# Update Genesis
jq '.app_state["staking"]["params"]["unbonding_time"] = "240s"' /home/config/genesis.json > /home/config/genesis.json.tmp  && mv /home/config/genesis.json.tmp /home/config/genesis.json
jq '.app_state["gov"]["voting_params"]["voting_period"] = "60s"' /home/config/genesis.json > /home/config/genesis.json.tmp  && mv /home/config/genesis.json.tmp /home/config/genesis.json

# Create/Update genesis configmap
echo "Update configmap and secret"
kubectl create configmap ${MONIKER%-*} --dry-run=client --output=yaml --from-file /home/config/genesis.json | kubectl apply --force=true --filename=-
kubectl create secret generic ${MONIKER} --dry-run=client --output=yaml --from-file MNEMONIC=/home/config/${MONIKER}_mnemonic | kubectl apply --force=true --filename=-
