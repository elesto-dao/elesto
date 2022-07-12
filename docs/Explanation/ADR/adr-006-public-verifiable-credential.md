# ADR 006: Verifiable Credentials

## Changelog

- 2022-07-04: Initial draft

## Status

DRAFT

## Abstract

Credentials are part of our daily interactions; driver's licenses are used to assert that we are capable of operating a motor vehicle, university degrees can be used to assert our level of education, and concert tickets are used to assert that we can attend an event.

Many credential types are used to store [personal identifiable information](https://en.wikipedia.org/wiki/Personal_data) (PII) and therefore should be kept private. However, there are many cases where it is more practical to publish credentials. For example, business licenses are used to assert that a business is legitimate and has the authorization to operate.

In this context, a credential is defined as a set of one or more claims made by an issuer and a [verifiable credential](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) (VC) as a tamper-evident credential that has authorship that can be cryptographically verified.

This document specifies support for VCs for private credentials and public credentials and support for credential definitions.

## Context

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has gained momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has captured the attention of the private and public sectors.

The SSI approach relies on two building blocks: decentralized identifiers (DID) and VCs. This architecture decision record (ADR) describes the implementation for supporting VCs in a Cosmos SDK-based blockchain.

This ADR aims to define a foundation for the components required to realize the Elesto objectives while ensuring the VC implementation is fully compliant with the W3C recommendations. **Successive iterations will address API ergonomics and standard compatibility issues.**

## Decision

The Elesto implementation for VCs will follow the [Verificable Credentials Data Model W3C Recommendations](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) to maximize compatibility with third-party tools and projects.

This ADR introduces two structures: 

- Credential Definition (CD) 
- Public Verifiable Credential (PVC)

### Credential Definition (CD)

CDs are used to describe the model of a credential, that is, its structure, fields, properties, and behavior. CDs consist of the following fields:

- `id` - the credential definition did
- `publisherId` - the DID of the publisher of the credential
- `name` - the human-readable name of the credential 
- `description` - the description of the credential usage
- `isPublic` - the credential can be issued on-chain
- `isActive` - a hint that indicates the credentials based on this schema should not be issued
- `supersededBy` - id of a new credential definition that supersedes this definition 
- `schema` - the credential [data schema](#credential-schema) 
- `vocab` - the credential [JSON-LD vocabulary](#credential-json-ld-vocabulary)

A CD cannot be deleted from the state. 

#### Credential Schema

A credential [data schema](https://www.w3.org/TR/vc-data-model/#data-schemas) is a machine-readable definition of fields and data types of a credential. 

The credential module supports the `JsonSchemaValidator2018` JSON-LD type schema format as described in the W3C [Verifiable Credentials JSON Schema Specification](https://w3c-ccg.github.io/vc-json-schemas/v1/index.html). The schema is stored as an uncompressed byte slice.

#### Credential JSON-LD vocabulary

A credential [JSON-LD context](https://www.w3.org/TR/json-ld11/#the-context) is a semantic vocabulary for the credential field descriptions. The vocabulary, in JSON format, is stored as an uncompressed byte slice.
  

### Public Verifiable Credential (PVC)

With PVCs, a VC is stored on-chain. 

To publish a VC on-chain, the sender must provide the credential definition that describes the credential. A credential is published only if:

- The credential definition has been published on-chain
- The credential definition allows publication of the credential (the field `isPublic` is true)
- The credential definition status is active (the field `isActive` is true)
- The credential conforms to the schema
- The proof of the credential can be positively verified

The `supersededBy` is used only for communication purposes and does not influence the publication of a credential. 
  
PVCs are intended for advertising information that is public domain and leverages the tamper-resistant capability of the blockchain. A PVC can be deleted from the state.

## Privacy Considerations

There is a risk that a credential definition allows publication of PII since the failsafe mechanism provided by the `isPublic` field can be misused by the credential publisher.   

## Security Considerations

The credential schema and the vocabulary must be stored uncompressed. Compression will make the node vulnerable to a [zip bomb](https://en.wikipedia.org/wiki/Zip_bomb) attack. 

## Consequences

The Elesto chain will provide support for credential schemas and vocabularies. At the same time, support for PVCs provides a strong foundation for a network of trust.
  
### Backward Compatibility

The credential module is a new module so backward compatibility is not a concern.

### Positive

- The ADR implementation improves compatibility with the SSI identity model.
- Closely following the W3C standard gives the best chance of successful interoperability with third-party components.

### Negative

N/A

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section summarizes issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.

Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

N/A

## References

- [W3C Recommendation Verificable Credentials Data Model v1.1](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/)
- [W3C Verifiable Credentials Data Model v1.1](https://www.w3.org/TR/vc-data-model/#data-schemas)
- [W3C Recommendation JSON-LD v1.1](https://www.w3.org/TR/2020/REC-json-ld11-20200716/)
- [W3C Verifiable Credentials JSON Schema Specification](https://w3c-ccg.github.io/vc-json-schemas/v1/index.html)
- [Personally Identifiable Information (PII)](https://en.wikipedia.org/wiki/Personal_data)
- [General Data Protection Regulation (GDPR)](https://eur-lex.europa.eu/eli/reg/2016/679/oj)
