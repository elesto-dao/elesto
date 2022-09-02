#!/bin/bash

acme_uuid=$(uuidgen)

echo "Create an Organization"

echo "Create ACME DDO"
elestod tx did create-did $acme_uuid \
 --from alice \
 --chain-id elesto -y --broadcast-mode block

# Manual task - issue a credential for the acme organization
elestod query credential prepare-credential organization
# https://business-register.eu/org/acme
# did:cosmos:elesto:business-register
#
# isvc: 11. Financial and insurance activities
# iso:
#
# id: did:cosmos:elesto:acme
#

echo "Issue organization credential"
elestod tx credential issue-public-credential \
 https://exmaple.id/organization \
 credential.organization.acme.json \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Check credential status"
elestod query credential public-credential-status https://business-register.eu/org/acme

echo "Revoke the credential"
elestod tx credential update-revocation-list \
 https://business-register.id/list/001 \
 --revoke 1 \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Check credential status"
elestod query credential public-credential-status https://business-register.eu/org/acme | jq

echo "Reset the credential status"
elestod tx credential update-revocation-list \
 https://business-register.id/list/001 \
 --reset 1 \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Create Organization public verifiable credential"
elestod tx credential create-revocation-list https://acme-corp.id/list/001 \
 --issuer did:cosmos:elesto:acme \
 --from alice \
 --chain-id elesto -y --broadcast-mode block


elestod query credential make-credential did:cosmos:elesto:organization --output json | jq
