#!/bin/bash

OPTS="--node http://65.108.254.247:26657 --chain-id elesto-canary-1 -y --broadcast-mode block"
FAUCET=http://cns.cosmos.prototyp.xyz:8000

### Set up issuer 1
R=$(($RANDOM%10000))
ISSUER=issuer$R

### Add keys and get tokens from the faucet
elestod keys add $ISSUER
curl $FAUCET?address\=$(elestod keys show $ISSUER -a)

### Wait 6s for next block
sleep 6

elestod tx credential publish-credential-definition \
  revocation-list-2020 \
  RevocationList2020 \
  credentials/revocationList2020.schema.json \
  credentials/revocationList2020.context.json \
 --public \
 --gas auto \
 --gas-adjustment 1.5 \
 --from $ISSUER \
 $OPTS

### Wait 6s for next block
sleep 6

echo "Publish a credential definition for an employee (private credential)"
elestod tx credential publish-credential-definition \
  organization-role \
  OrganizationRole \
  credentials/organization-role.schema.json \
  credentials/organization-role.context.json \
 --gas auto \
 --from $ISSUER \
 $OPTS

