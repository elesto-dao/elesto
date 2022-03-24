# State Transitions

This document describes the state transitions pertaining a [DidDocument](02_state.md#diddocument) according to the [did operations](https://www.w3.org/TR/did-core/#method-operations):

1. [Create](03_state_transitions.md#Create)
2. [Resolve](03_state_transitions.md#Resolve)
3. [Update](03_state_transitions.md#Update)
4. [Deactivate](03_state_transitions.md#Deactivate)

### Create

[DidDocument](02_state.md#diddocument) are created via the rpc method [CreateDidDocument](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L13) that accepts a [MsgCreateDidDocument](./04_messages.md#MsgCreateDidDocument) messages as parameter.

The operation will fail if:
- the signer account has insufficient funds
- the did is malformed
- a did document with the same did exists
- verifications
  - the verification method is invalid (according to the verification method specifications)
  - there is more than one verification method with the same id
  - relationships are empty
  - relationships contain unsupported values (according to the did method specifications)
- services are invalid (according to the services specifications)

Example:

<!--

elestod tx did create-did \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet
-->

```javascript
/* gRPC message */
CreateDidDocument(
    MsgCreateDidDocument(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        [], // controller
        [   // verifications
            {
                "relationships": ["authentication"],
                {
                    "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
                    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
                    "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
                    "type": "EcdsaSecp256k1VerificationKey2019"
                },
                [],
            },
        ],
        [], // services
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
}

```

##### Implementation Source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L27
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L43

### Resolve

[DidDocument](02_state.md#diddocument) are resolved via the rpc method [QueryDidDocument](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/query.proto#L17) that accepts a [QueryDidDocumentRequest](./04_messages.md#QueryDidDocumentRequest) messages as parameter.


The operation will fail if:
- the did does not exists

Example:

<!--
elestod query did did did:cosmos:cash:900d82bc-2bfe-45a7-ab22-a8d11773568e \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet \
 --output=json | jq
-->

```javascript
/* gRPC message */
QueryDidDocument(
    QueryDidDocumentRequest(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e"
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
}

```

##### Implementation Source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/grpc_query.go#L32
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/query.go#L70

### Update

[DidDocument](02_state.md#diddocument) are updated via the rpc methods:

- [UpdateDidDocument](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L15)
- [AddVerification](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L18)
- [RevokeVerification](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L20)
- [SetVerificationRelationships](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L22)
- [AddService](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L25)
- [DeleteService](https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L27)

All the operations will fail if:

- the signer account has insufficient funds
- the signer account address doesn't match the verification method listed in the `Authorization` verification relationships
- the target did does not exists

The following sections provide specific details for each method invocation.

#### UpdateDidDocument

The  `UpdateDidDocument` method will is used to **overwrite** the  [DidDocument](02_state.md#diddocument). It accepts a [MsgUpdateDidDocument](./04_messages.md#MsgUpdateDidDocument) as a parameter.

The operation will fail if:

- any integrity and sanity checks on the DID document fail
- the DID document is not found

<!--

elestod tx did update-did-document \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8 \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet
-->

```javascript
/* gRPC message */
UpdateDidDocument(
    MsgUpdateDidDocument(
        {
          "context": [
            "https://www.w3.org/ns/did/v1"
          ],
          "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
          "controller": [
             "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
          ],
          "verificationMethod": [
            {
              "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
              "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
              "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
              "type": "EcdsaSecp256k1VerificationKey2019"
            }
          ],
          "service": [],
          "authentication": [
            "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
          ],
          "assertionMethod": [],
          "keyAgreement": [],
          "capabilityInvocation": [],
          "capabilityDelegation": []
        },
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    )
)

/* Resolved DID Document */

{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
        "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
}
```

##### Implementation Source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L65
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L277

#### AddVerification

The `AddVerification` method is used to add new [verification methods](https://w3c.github.io/did-core/#verification-methods) and [verification relationships](https://w3c.github.io/did-core/#verification-relationships) to a [DidDocument](02_state.md#diddocument). It accepts a [MsgAddVerification](./04_messages.md#MsgAddVerification) as a parameter.

The operation will fail if:

- the verification method is invalid (according to the verification method specifications)
- the verification method id already exists for the did document
- the verification relationships are empty
- the verification relationships contain unsupported values (according to the did method specification)

<!--

elestod tx did add-verification-method \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 cosmospub1addwnpepqduxp90pt6ez3a8p26fwmfhqvparz0xqsxk4f4564hg46527xpzeq82cerm \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet
-->

```javascript
/* gRPC message */
AddVerification(
    MsgAddVerification(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        {
            "relationships": ["authentication"],
            {
                "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
                "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
                "publicKeyHex": "03786095e15eb228f4e15692eda6e0607a313cc081ad54d69aadd15d515e304590",
                "type": "EcdsaSecp256k1VerificationKey2019"
            },
            [],
        },
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */

{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      },
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
        "publicKeyHex": "03786095e15eb228f4e15692eda6e0607a313cc081ad54d69aadd15d515e304590",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
}

```

##### Implementation Source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L107
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L101

#### RevokeVerification

The `RevokeVerification` method is used to remove existing [verification methods](https://w3c.github.io/did-core/#verification-methods) and [verification relationships](https://w3c.github.io/did-core/#verification-relationships) from a [DidDocument](02_state.md#diddocument). It accepts a [MsgRevokeVerification](./04_messages.md#MsgRevokeVerification) as a parameter.

The operation will fail if:

- the verification method id is not found


<!--

elestod tx did revoke-verification-method \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2 \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
RevokeVerification(
    MsgRevokeVerification(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L202
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L201


#### SetVerificationRelationships


The `SetVerificationRelationships` method is used to **overwrite** existing [verification relationships](https://w3c.github.io/did-core/#verification-relationships) for a [verification methods](https://w3c.github.io/did-core/#verification-methods) in a [DidDocument](02_state.md#diddocument). It accepts a [MsgSetVerificationRelationships](./04_messages.md#MsgSetVerificationRelationships) as a parameter.

The operation will fail if:

- the verification method id is not found for the target did document
- the verification relationships are empty
- the verification relationships contain unsupported values (according to the did method specification)

<!--

elestod tx did set-verification-relationships \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0 \
 --relationship capabilityInvocation \
 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
SetVerificationRelationships(
    MsgSetVerificationRelationships(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        ["authentication", "capabilityInvocation"]
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L287
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L319


#### AddService


The `AddService` method is used to add a [service](https://w3c.github.io/did-core/#services) in a [DidDocument](02_state.md#diddocument). It accepts a [MsgAddService](./04_messages.md#MsgAddService) as a parameter.

The operation will fail if:

- a service with the same id already present in the did document
- the service definition is invalid (according to the did services specification)

<!--

elestod tx did add-service \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \

TODO

 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
AddService(
    MsgAddService(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        {
            "agent:xyz",
            "DIDCommMessaging",
            "https://agent.xyz/1234",
        }
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [
        {
            "agent:xyz",
            "DIDCommMessaging",
            "https://agent.xyz/1234",
        }
    ],
    "authentication": [
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L150
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L154

#### DeleteService


The `DeleteService` method is used to remove a [service](https://w3c.github.io/did-core/#services) from a [DidDocument](02_state.md#diddocument). It accepts a [MsgDeleteService](./04_messages.md#MsgDeleteService) as a parameter.

The operation will fail if:

- the service id does not match any service in the did document
<!--

elestod tx did add-service \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \

TODO

 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
DeleteService(
    MsgDeleteService(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "agent:xyz",
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/keeper/msg_server.go#L150
- client: https://github.com/elesto-dao/elesto/blob/v1.0.0/x/did/client/cli/tx.go#L154


### AddController

The `AddController` method is used to add a [controller](https://w3c.github.io/did-core/#did-controller) to a [DidDocument](02_state.md#diddocument). It accepts a [MsgAddContoller](./04_messages.md#MsgAddController) as a parameter.

The operation will fail if:

- the controller DID is not of type `key`
<!--

elestod tx did add-controller \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \

TODO

 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
AddController(
    MsgAddController(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: --
- client: --

### DeleteController

The `DeleteController` method is used to delete a [controller](https://w3c.github.io/did-core/#did-controller) from a [DidDocument](02_state.md#diddocument). It accepts a [MsgDeleteContoller](./04_messages.md#MsgDeleteController) as a parameter.

The operation will fail if:

- the DID document is not found
<!--

elestod tx did add-controller \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \

TODO

 --from vasp --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet

-->

```javascript
/* gRPC message */
DeleteController(
    MsgDeleteController(
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
        "cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved DID Document */
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [],
    "authentication": [
        "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [
      "did:cosmos:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "capabilityDelegation": []
  },
}

```

##### Implementation source

- server: --
- client: --

### Deactivate

Currently not supported
