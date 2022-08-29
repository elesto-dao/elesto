---
title: Issuing verifiable credentials
---

This section covers how to issue a verifiable credential (off chain).

There are 3 actors 

- Issuer 
- Holder 
- Verifier

The example goes for a club membership

## Before starting 

load this 3 accounts and create the DID

#### Issuer 

`elestod keys add -i issuer`

```
delay retire mansion suspect merit large pudding amateur blossom cabin swift fold scare wine advice ticket solar effort either target limit august mule side
```

#### Holder 

`elestod keys add -i holder`

```
ski fever flavor satoshi hurry amused raccoon amazing annual valley seek boost neck novel staff service vacant wonder know blossom hobby harsh turtle bone
```

#### Verifier

`elestod keys add -i verifier`

```
shop laugh stool bonus cloud fashion impact jaguar fabric guilt wealth donate act note gas prevent length text gadget plug iron document unaware movie
```


??? Hint "How to get tokens from the faucet"
    ```
    curl -L -X POST -d "{\"address\": \"$(elestod keys show issuer -a)\"}" faucet.elesto-canary-1.elesto.id
    ```


### Create the DID documents (only once!)

#### Issuer

```
elesto tx did create-did issuer \
--from issuer \
--chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```
see it in the resolver:

- [did:cosmos:elesto-canary-1:issuer](https://resolver-driver.canary-1.elesto.id/identifier/did:cosmos:elesto-canary-1:issuer)


#### Holder
```
elesto tx did create-did holder \
 --from holder \
 --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

see it in the resolver: 

- [did:cosmos:elesto-canary-1:holder](https://resolver-driver.canary-1.elesto.id/identifier/did:cosmos:elesto-canary-1:holder)


#### Verifier
```
elesto tx did create-did verifier \
 --from verifier \
 --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```
see it in the resolver: 

- [did:cosmos:elesto-canary-1:verifier](https://resolver-driver.canary-1.elesto.id/identifier/did:cosmos:elesto-canary-1:verifier)


## Scenario #1 - Membership 

Schema number 1: membership 

`did:cosmos:elesto-canary-1:cd-membership`

- [schema](https://resolver-driver.canary-1.elesto.id/schemas/did:cosmos:elesto-canary-1:cd-membership)
- [context](https://resolver-driver.canary-1.elesto.id/context/did:cosmos:elesto-canary-1:cd-membership)


publish the credential definition
```
elestod tx credential publish-credential-definition \
   did:cosmos:elesto-canary-1:cd-membership \
   ClubMembership \
   scripts/seeds/credentials/membership.schema.json \
   scripts/seeds/credentials/membership.context.json \
  --gas auto \
  --description "A simple club membership credential" \
  --from ga --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
  #--public \
  #
```

query the credential definition
```
elestod query credential credential-definition \
 did:cosmos:elesto-canary-1:cd-membership \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443 -o json | jq
```


### Credential exchanges

The **issuer** compiles the credential 
```
elestod query credential prepare-credential \
 did:cosmos:elesto-canary-1:cd-membership \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

- id: https://test-my-credential/1
- issuer: did:cosmos:elesto-canary-1:issuer
- subject id: did:cosmos:elesto-canary-1:holder
- membership.1.json

The **issuer** signs the credential
```
elestod tx credential sign-credential membership.1.json --from issuer
```

The **issuer** ships it to the the holder
```
wormhole send membership.1.signed.json
```

??? Question "What is a wallet?"
    - **Crytpo Wallet**: a key store => recoverable using the mnemonic
    - **SSI Wallet**: the phisical storage of your credentials => not recoverable



The **holder** receives it 
```
wormhole receive xxx-yyyy-zzzz
```


Now the verifier asks for the membership credential from the holder


The **holder** ships the credential to the verifier
```
wormhole send membership.1.signed.json
```


The **verifier** receives it

```
wormhole receive aaa-bbb-ccc
```

The **verifier** verifies that the credential is valid 
```
elestod query credential verify-credential \
 did:cosmos:elesto-canary-1:cd-membership \
 membership.1.signed.json \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

plot twist! 

The **holder** signs the credential with it's own key
```
elestod tx credential sign-credential membership.1.signed.json --from holder
```


### That's great but, problems !

- the verifier does not make statements about the issuer
    - solutions: network of trust
- the verifier only verify the credential, but not who sent it! 
    - solutions: DIDComm & Verfiable Presentations 
- not much transactions there, innit?
    - solutions: stablecoin & defi protocols regulatory friendly
- what if the memebership is revoked??
    - glad you asked: let's see scenario number 2

## Schema number 2: mebership with revocations


### Before we start

`did:cosmos:elesto-canary-1:revocation-list-2020`

- [schema](https://resolver-driver.canary-1.elesto.id/schemas/did:cosmos:elesto-canary-1:revocation-list-2020)
- [context](https://resolver-driver.canary-1.elesto.id/context/did:cosmos:elesto-canary-1:revocation-list-2020)

`did:cosmos:elesto-canary-1:cd-membership-rl`

- [schema](https://resolver-driver.canary-1.elesto.id/schemas/did:cosmos:elesto-canary-1:cd-membership-rl)
- [context](https://resolver-driver.canary-1.elesto.id/context/did:cosmos:elesto-canary-1:cd-membership-rl)


```
elestod tx credential publish-credential-definition \
   did:cosmos:elesto-canary-1:cd-membership-rl \
   ClubMembership \
   scripts/seeds/credentials/membership-with-revocation.schema.json \
   scripts/seeds/credentials/membership-with-revocation.context.json \
  --gas auto \
  --description "A simple club membership credential (with support for revocation lists)" \
  --from ga --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
  # --public \
```

```
elestod query credential credential-definition \
 did:cosmos:elesto-canary-1:cd-membership-rl \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443 -o json | jq
```

```
elestod query credential credential-definition \
 did:cosmos:elesto-canary-1:revocation-list-2020 \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443 -o json | jq
```

the issuer must issue a public verifiable credential!

```
elestod tx credential create-revocation-list \
 https://issuer.id/rl/001 \
 --definition-id did:cosmos:elesto-canary-1:revocation-list-2020 \
 --issuer did:cosmos:elesto-canary-1:issuer
 --from issuer --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

### Credential exchanges

The **issuer** prepares the credential

```
elestod query credential prepare-credential \
 did:cosmos:elesto-canary-1:cd-membership-rl \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

- id: https://test-my-credential/1
- issuer: did:cosmos:elesto-canary-1:issuer
- subject id: did:cosmos:elesto-canary-1:holder
- membership.1.json
- revocation ID 600

The **issuer** sign the credential
```
elestod tx credential sign-credential membership.1.json --from issuer
```

The **issuer** ships it to the the holder
```
wormhole send membership.1.signed.json
```
The **holder** receives it 
```
wormhole receive xxx-yyyy-zzzz
```


Now the verifier asks for the membership credential from the holder


The **holder** ships the credential to the verifier
```
wormhole send membership.1.signed.json
```


The **verifier** receives it

```
wormhole receive aaa-bbb-ccc
```

The **verifier** verifies that the credential is valid 
```
elestod query credential verify-credential \
 did:cosmos:elesto-canary-1:cd-membership-rl \
 membership.1.signed.json \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

The **issuer** meanwhile revokes the credential 

```
elestod tx credential update-revocation-list \
 https://issuer.id/rl/001 \
 --definition-id did:cosmos:elesto-canary-1:revocation-list-2020 \
 --revoke 600 \
 --from issuer \
 --chain-id elesto-canary-1 -y --b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443
```

<!-- 
```
elestod tx credential update-revocation-list \
 https://ga000.id/rl/001 \
 --definition-id did:cosmos:elesto-canary-1:revocation-list-2020 \
 --revoke 600 \
 --from ga --chain-id elesto-canary-1 -y -b block --node https://rpc.harbor.elesto-canary-1.elesto.id:443

```
-->

The **verifier** checks that the credential is not revoked, and ....

```
elestod query credential credential-status \
 membership.1.signed.json \
 --node https://rpc.harbor.elesto-canary-1.elesto.id:443 -o json | jq

```


## The way forward

This interactions needs to be user friendly.

## Questions



