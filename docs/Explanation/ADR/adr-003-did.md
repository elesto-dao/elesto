# ADR 003: DID

## Changelog

- 2022-02-14: Moved to last call
- 2022-02-14: Renamed to ADR002
- 2021-09-23: Added security and privacy considerations
- 2021-08-02: Initial draft

## Status

LAST CALL 2022-02-28

## Abstract

[Decentralized identifiers](https://www.w3.org/TR/did-core/) (DIDs) are a type of identifier that enables verifiable, decentralized digital identity. A DID refer to any subject (for example, a person, organization, thing, data model, abstract entity, and so on) as determined by the controller of the DID.

This document specifies the DID method for a Cosmos SDK-based implementation of the W3C recommendation, its properties, operations, and an explanation of the process to resolve DIDs to the resources they represent.

## Context

The Elesto project aims to provide a state-of-the-art platform for the hosting of collateralized stable coins that is compliant with:

 - EU regulations such as General Data Protection Regulation (GDPR) and Markets in Crypto-Assets (MiCA)
 - International recommendations such as the Financial Action Task Force (FATF) "Travel Rule"
 - Local anti-money laundering (AML) regulations

The Elesto platform is based on the following principles:

- Open financial infrastructure is a public good
- Money laundering prevention also benefits society
- Users benefit from a strict privacy-respecting approach (GDPR)

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has gained momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has captured the attention of both the private and public sectors.

The SSI approach relies on two building blocks: decentralized identifiers (DID) and verifiable credentials (VC). This architecture decision record (ADR) describes the DID implementation in a Cosmos SDK-based blockchain.

This ADR aims to define a foundation for the necessary components to realize the Elesto objectives while ensuring the implementation of the DID is fully compliant with the W3C specifications. **Successive iterations will address API ergonomics and standard compatibility issues.**

## Decision


The Elesto implementation for DIDs will follow the [DID W3C core recommendations](https://www.w3.org/TR/2021/PR-did-core-20210803/) to maximize compatibility with 3rd party tools and projects.


### DID Method Name

The namestring that shall identify the Elesto DID method is: `cosmos`.

A DID that uses the Elesto method MUST begin with the following prefix: `did:cosmos`. Per the [W3C DID specification](https://www.w3.org/TR/2021/PR-did-core-20210803/), this prefix string MUST be in lowercase. The remainder of the DID, after the prefix, is as follows:

#### Method Specific Identifier


The namespace specific identifier is defined by the following ABNF:

```ABNF
cosmos-did                = "did:cosmos:" identifier-type
identifier-type           = cosmos-key / unique-identifier
cosmos-key                = "key:" 1*255id-char "1" 20*255HEXDIG
unique-identifier         = 38*255id-char
id-char                   = ALPHA / DIGIT / (ALPHA "-") / (DIGIT "-") / (ALPHA ":") / (DIGIT ":")
```

For the `unique-identifier` it is RECOMMENDED to use a UUID.

The `identifier-type` distinguishes between two DID types:

1. The `key` type, inspired from the [`did:key`](https://w3c-ccg.github.io/did-method-key/) method
2. The `unique-identifier` type that identifies the DID and the originating network of the DID

DIDs of `key` type are ephemeral and immutable. DIDs of `key` type are always generated from the Cosmos blockchain address they refer to. For example, see this DID of `key` type:

- `did:cosmos:key:elesto1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`

DIDs of `net` type are persistent and mutable. DIDs of `net` type are stored in the node database and can be created and updated according to rules described in the [DID Operations](#did-operations) section. For example, see this DID of `net` type:

- `did:cosmos:elesto:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb`


##### DID Operations

DID and associated DID documents are managed by a Cosmos SDK module that uses the gRPC communication protocol. See [Method operations](https://www.w3.org/TR/2021/PR-did-core-20210803//#method-operations) for details on how the create, read, update and delete (CRUD) operations are handed in a Cosmos DID.

###### Create

To create and publish a DID document use the message

```golang
MsgCreateDidDocument(id string, signerPubKey string)
```

The message parameters are the DID to be created and the `signerPubKey`. The `signerPubKey` MUST be the public key of the account that signs the transaction. The public key MUST be used to attach a verification method of type `EcdsaSecp256k1VerificationKey2019` with the value of `publicKeyMultibase` that contains the public key encoded according to the Multibase Data Format [Hexadecimal upper-case encoding](https://datatracker.ietf.org/doc/html/draft-multiformats-multibase#appendix-B.1).

The verification method controller MUST be one of the following:

- The DID of the document
- The DID of `key` type of the address that signs the transaction

The verification method id SHOULD be generated as:

```
{verificationMethodController}#{CosmosAddressFromPubKey}
```

The verification method id MUST be listed in the `authentication` relationships.

If the input DID is not a valid DID for the Cosmos method, or if the DID already exists on-chain, the message returns an error.

Contextually with the creation of a DID document, a [DID document metadata](#did-document-metadata) MUST be created with the following values:

- The hash of the transaction as `versionId`
- The block time for the `created` and `updated` fields
- The `deactivated` field is `false`


To address privacy concerns:

- Use an id that is different from the blockchain account address
- Isolate the verification methods to the DID subject (for example, during key rotation)


> **Note:** A more fine-grained DID creation method can be implemented to save gas by executing a single transaction in a complex DID scenario.


###### Resolve and Verify

The underlying blockchain protocol guarantees the integrity of the DID documents stored on the ledger.

A DID can be resolved using the gRPC message:

```golang
QueryDidDocumentRequest(did string)
```

This example shows a DID document that was resolved using the gRPC interface:

```javascript
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:elesto:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:elesto:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:elesto:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyMultibase": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "authentication": [
      "did:cosmos:elesto:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ]
  },
}
```


> Note: the DID document resolution following the [W3C DID core recommendations](https://www.w3.org/TR/2021/PR-did-core-20210803/) and pass the [DID Core Specification Test Suite](https://w3c.github.io/did-test-suite/) SHOULD be managed outside the node implementation, in a dedicated project.


###### Update

A DID document can be updated only if it is persisted on-chain on-chain and must it is of subtype `net` and it's not in `deactivated` metadata property is set to `false`. The possible updates are the following:

- Add and remove controllers
- Add and remove services
- Add and revoke verification method
- Set verification method relationships
- Replace the DID document content

The constraints on updating a DID document are **one of** the following:

- The `signerAccount` must exist as a verification method in a verification relationship of type `authentication`
- The `signerAccount` `key` DID is listed as a DID controller.

Each update operation MUST update the `versionId` and the `updated` field of the associated [DID document metadata](#did-document-metadata) with the transaction hash and the block time respectively.

**Add DID document controller**

Add DID document controllers using the gRPC message:

```golang
AddController(
   MsgAddController{ Id string, ControllerDid []string, SignerAccount string}
)
```

The parameters are as follows:

 - `Id` identifies the DID document
 - `Controller` the DID to be added as a controller
 - `SignerAccount` is the account address that signs the transaction

A controller of a DID document MUST be a DID of subtype `key`.

**Remove a DID document controller**

Remove DID document controllers using the gRPC message:

```golang
DeleteController(
   MsgDeleteController{ Id string, ControllerDid []string, SignerAccount string}
)
```

The parameters are:

 - `Id` identifies the DID document
 - `Controller` the DID to be removed from the controllers list
 - `SignerAccount` is the account address that signs the transaction

**Add a service**

A service MUST be an entity with the following properties:

- `id`: a valid RFC3986 URI string.
- `type`: a non empty string.
- `serviceEndpoint`: a valid RFC3986 URI string.

A service can be added using the gRPC method:

```golang
AddService(
   MsgAddService{Id string, ServiceData Service, SignerAccount string}
)
```

The parameters are:

 - `Id` identifies the DID document
 - `ServiceData` contains the service definition as described above
 - `SignerAccount` is the account address that signs the transaction


**Remove a service**

A service can be deleted using the gRPC message:

```golang
DeleteService(
   MsgDeleteService{Id string, ServiceID string, SignerAccount string}
)
```

The parameters are as follows:

 - `Id` identifies the DID document
 - `ServiceID` the id of the service to remove
 - `SignerAccount` is the account address that signs the transaction



**Verification methods and relationships**


A verification method and its relationships are manipulated via the `Verification` object with the following properties:

- `method` a verification method as described in the [W3C specification](https://www.w3.org/TR/2021/NOTE-did-spec-registries-20211102/#verificationmethod)
- `relationships` a non-empty list of relationships associated to the verification method
- `context` a list of additional json-ld contexts to be added to the did document


**Add a verification method**

Add a new verification method using the gRPC message:

```golang
AddVefication(
   MsgAddVerification{
      Id string,
      Verification Verification,
      SignerAccount string
   }
)
```

The fields of the `MsgAddVerification` message are:

 - `Id` identifies the did document
 - `Verification` the verification as described before
 - `SignerAccount` is the account address that is signing the transaction

The list of relationships must contain only valid [verification relationships](#verification-relationships) names as defined in the DID document and MUST be non-empty.

The verification method id must be unique within the DID document.

**Remove a verification method**

Remove a verification method using the gRPC message:

```golang
RevokeVerification(
   MsgRevokeVerification {
      Id string
      MethodId string
      SignerAccount string
   }
)
```

The fields of the `MsgAddVerification` message are:

 - `Id` identifies the did document
 - `MethodId` the verification method id to remove
 - `SignerAccount` is the account address that is signing the transaction

**Set verification method relationships**

Set the relationships of a verification method using the gRPC message:

```golang
SetVerificationRelationships(
   MsgSetVerificationRelationships{
      Id string,
      MethodId string,
      Relationships []string,
      SignerAccount string
   }
)
```

The fields of the `MsgSetVerificationRelationships` message are:

 - `Id` identifies the did document
 - `MethodId` identifies the method id to change the relationships for
 - `Relationships` the list of relationships to set
 - `SignerAccount` is the account address that is signing the transaction

The list of relationships must contain only valid [verification relationships](#verification-relationships) names as defined in the DID document and MUST be non-empty.


**Overwrite an existing DID document**

To completely replace the content of a DID document use the message:

```golang
UpdateDidDocument(
   MsgUpdateDidDocument {
      Doc DidDocument
      SignerAccount string
   }
)
```

The fields of the `MsgSetVerificationRelationships` message are:

 - `Doc` the updated DID document
 - `SignerAccount` is the account address that is signing the transaction

The DID of the new DID document must match the existing document. All the validation rules also apply while replacing the document.

###### Deactivate

Deactivation of a DID document is not currently supported

### Method-Specific Properties

#### DID Core Verification Material

The [Verification Material](https://www.w3.org/TR/2021/PR-did-core-20210803//#verification-material) type MUST support:

- Type `EcdsaSecp256k1VerificationKey2019` with `pubKeyMultibase` to encode a Cosmos account public key in hexadecimal format
- Type `CosmosAccountAddress` with `blockchainAccountID` to represent a Cosmos account


Support for other verification materials can be added.

#### Verification Relationships

The DID document MUST support the following [verification relationships](https://www.w3.org/TR/2021/PR-did-core-20210803//#verification-relationships):

- [`authentication`](https://www.w3.org/TR/2021/PR-did-core-20210803//#authentication) - authorizes amends to the DID document
- [`assertionMethod`](https://www.w3.org/TR/2021/PR-did-core-20210803//#assertion)
- [`keyAgreement`](https://www.w3.org/TR/2021/PR-did-core-20210803//#key-agreement)
- [`capabilityInvocation`](https://www.w3.org/TR/2021/PR-did-core-20210803//#capability-invocation)
- [`capabilityDelegation`](https://www.w3.org/TR/2021/PR-did-core-20210803//#capability-delegation)


#### DID Document Metadata

The implementation for [DID document metadata](https://www.w3.org/TR/2021/PR-did-core-20210803//#did-document-metadata) MUST report the following properties for a DID document:

- `created`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the creation date that is the UTC date associated with the block height when the DID document was submitted the first time
- `updated`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the last update date that is the UTC date associated with the block height when the DID document was submitted the last time
- `deactivated`: a boolean field that indicates if the DID document is [deactivated](#Deactivate)
- `versionId`: a hex-encoded BLAKE2b hash of the transaction that created or updated the DID

#### DID Resolution Metadata

The [DID Resolution Metadata](https://www.w3.org/TR/2021/PR-did-core-20210803//#did-resolution-metadata) is outside the scope of the gRPC interface and is not covered in this ADR.

#### DID URL Syntax

The [DID URL Syntax](https://www.w3.org/TR/2021/PR-did-core-20210803//#did-url-syntax) is outside the scope of the gRPC interface and is not covered in this ADR.

#### DID Query Parameters

The [DID Query parameters](https://www.w3.org/TR/2021/PR-did-core-20210803//#did-parameters) URL is outside the scope of the gRPC interface and is not covered in this ADR.

<!--

The implementation MUST support the following query parameters:

- `versionId` - to retrieve a DID document with a specific version
- `versionTime` - to retrieve the version of the DID document valid at a specific time, the parameter MUST be a valid [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime).

The format for the queries is:
-->

## Privacy Considerations

When any data (for example, W3C Verifiable Credentials) is associated with Cosmos DIDs, sharing that data would also impose sharing the on-chain data graph (for example, transaction history) of the blockchain account that controls the DID.

Using personally identifiable information as DID Method-specific identifiers (for example, account name Alice) discloses personal information every time the DID is shared with a counterparty. This specification DOES NOT endorse the use of identifiers that correlate to human beings or other sensible subjects.

## Security Considerations

Ephemeral DIDs (`did:cosmos:key` type) are generated based on blockchain addresses. If access to the authoritative keys for an account is lost, the DID's control and verifiable data issued by the DID is lost.

## Consequences

The Cosmos ecosystem WILL HAVE a DID module compatible with the W3C standard and offers a high chance of compatibility with third-party components such as cloud and edge agents, resolvers, etc.

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

- The ADR implementation provides the foundation for interoperability with the DID standard and the SSI identity approach.
- Closely following the W3C standard gives the best chances of successful interoperability with third-party components.

### Negative

- The implementation rigidly follows the W3C specification which leaves little room for extensibility. This approach might become an issue for wider adoption.

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section summarizes issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

- The `did:key` method specifies a key format that is different from the one used in this ADR. This ADR needs to be amended or follow a different approach.
- The approach proposed is somewhat locked into the current implementation and will have to be revised in successive iterations.
## Test Cases [optional]

N/A

## References

- [DID Core v1.0](https://www.w3.org/TR/2021/PR-did-core-20210803/)
- [DID Specification Registries (11 February 2022)](https://w3c.github.io/did-spec-registries)
