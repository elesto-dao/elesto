#!/bin/bash

echo "Publish a credential definition (for a public credential)"
elestod tx credential publish-credential-definition revocation-list-2020 RevocationList2020 03_schema.json 03_vocab.json \
 --public \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential"

elestod query credential credential-definition did:cosmos:elesto:revocation-list-2020 --output json | jq

echo "Create issuer DDO"

elestod tx did create-did example-credential-issuer \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Sign and issue public credential"

elestod tx credential issue-public-credential revocation-list-2020 03_credential.json \
 --export 03_credential.signed.json \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential"

elestod query credential public-credential  'https://example.com/credentials/status/3' --output json | jq

