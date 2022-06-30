#!/bin/bash


# Step #1 public credential definitions
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
elestod tx did create-did business-register \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block


elestod query did did did:cosmos:elesto:business-register -o json | jq

## Use the special command to create the revocation list credential
elestod tx credential create-revocation-list https://business-register.id/list/001 \
 --issuer did:cosmos:elesto:business-register \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block


echo "Query credential"
elestod query credential public-credential  'https://business-register.id/list/001' --output json | jq

echo "Query credential - native format"
elestod query credential public-credential  'https://business-register.id/list/001' --native --output json | jq

echo "Query credentials issued by issuer"
elestod query credential public-credentials-by-issuer did:cosmos:elesto:business-register | jq

echo "Query credentials issued by issuer - native format"
elestod query credential public-credentials-by-issuer did:cosmos:elesto:example-credential-issuer --native --output json | jq

## Step 2 - Publish a public verifiable credential definition

elestod query credential credential-definition did:cosmos:elesto:organization-role --output json | jq

## Step 3 - Register did for organization
echo "Create an Organization"

echo "Create ACME DDO"
elestod tx did create-did acme \
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
