#!/usr/local/bin/bash

INIT_FROM=${INIT_FROM:-'genesis-0'}

#Wait for pods
kubectl wait pod -l app.kubernetes.io/name=genesis --for condition=Ready

# Init node
if [[ ! -f /home/config/node_key.json ]]; then
  echo "Init ${MONIKER} on ${CHAIN_ID}"
  elestod init ${MONIKER} --home=/home/ --chain-id=${CHAIN_ID}
  sed -i 's|keyring-backend = ".*"|keyring-backend = "test"|' /home/config/client.toml
fi

# Update configmap & secret
echo "Update configmap and secret"
INIT_FROM_ID=$(wget -O - -o /dev/null http://${INIT_FROM%-*}:26657/status | jq -r '.result.node_info.id' | tr -d '\n')
cat /dev/null > /home/config/seeds-env
echo "P2P_PERSISTENT_PEERS=${INIT_FROM_ID}@${INIT_FROM%-*}:26656" >> /home/config/seeds-env
echo "P2P_PRIVATE_PEER_IDS=${INIT_FROM_ID}" >> /home/config/seeds-env
kubectl create configmap ${MONIKER%-*}s-env --dry-run=client --output=yaml --from-env-file=/home/config/seeds-env | kubectl apply --force=true --filename=-
