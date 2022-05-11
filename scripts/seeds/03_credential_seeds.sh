#!/bin/bash

echo "Publish a credential definition (for a public credential)"
elestod tx credentials publish-credential-definition revocation-list-2020 RevocationList2020 03_schema.json 03_vocab.json \
 --public \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential"

elestod query credentials credential-definition did:cosmos:elesto:revocation-list-2020 --output json | jq

echo "Issue public credential"

