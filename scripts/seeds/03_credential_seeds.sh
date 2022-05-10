#!/bin/bash

echo "Publish a credential definition"
elestod tx credentials publish-credential-definition revocation-list-2020 RevocationList2020 03_schema.json 03_vocab.json \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

 echo "Query credential"

elestod query credentials credential-definition did:cosmos:elesto:revocation-list-2020 --output json | jq

echo "Create a DID doc for the issuer (by the regulator account)"
elestod tx did create-did credential-issuer \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query the issuer DID"
elestod query did did did:cosmos:elesto:credential-issuer --output json | jq

echo "Register a credential issuer"
elestod tx credentials register-issuer credential-issuer \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query the credential issuer DID"
elestod query credentials issuer did:cosmos:elesto:credential-issuer --output json | jq


echo "Add issuance credential"

echo "Issue public credential"

