#!/bin/bash

business_register_uuid=$(uuidgen)

echo "Publish a credential definition for revocation-list (for a public credential)"
elestod tx credential publish-credential-definition \
  https://exmaple.id/revocation-list-2020 \
  RevocationList2020 \
  credentials/revocationList2020.schema.json \
  credentials/revocationList2020.context.json \
 --public \
 --gas auto \
 --gas-adjustment 1.5 \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential definition"
elestod query credential credential-definition https://exmaple.id/revocation-list-2020 --output json | jq

echo "Publish a credential definition for an organization (public credential)"
elestod tx credential publish-credential-definition \
  https://exmaple.id/organization \
  Organization \
  credentials/organization.schema.json \
  credentials/organization.context.json \
 --public \
 --from regulator \
 --gas auto \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential definition"
elestod query credential credential-definition https://exmaple.id/organization --output json | jq

echo "Publish a credential definition for an employee (private credential)"
elestod tx credential publish-credential-definition \
  https://exmaple.id/organization-role \
  OrganizationRole \
  credentials/organization-role.schema.json \
  credentials/organization-role.context.json \
 --from regulator \
 --gas auto \
 --chain-id elesto -y --broadcast-mode block

echo "Query credential definition"
elestod query credential credential-definition https://exmaple.id/organization-role --output json | jq


echo "Create business register DDO"
elestod tx did create-did $business_register_uuid \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block


elestod query did did did:cosmos:elesto:$business_register_uuid -o json | jq

## Use the special command to create the revocation list credential
elestod tx credential create-revocation-list https://business-register.id/list/001 \
 --definition-id https://exmaple.id/revocation-list-2020 \
 --issuer did:cosmos:elesto:$business_register_uuid \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block


echo "Query credential"
elestod query credential public-credential  'https://business-register.id/list/001' --output json | jq

echo "Query credential - native format"
elestod query credential public-credential  'https://business-register.id/list/001' --native --output json | jq

echo "Query credentials issued by issuer"
elestod query credential public-credentials-by-issuer did:cosmos:elesto:$business_register_uuid | jq

