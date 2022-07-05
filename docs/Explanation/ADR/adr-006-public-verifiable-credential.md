# ADR 006: Credentials

## Changelog

- 2022-07-04: Initial draft

## Status

DRAFT

## Abstract

Credentials are part of our daily interactions; driver's licenses are used to assert that we are capable of operating a motor vehicle, university degrees can be used to assert our level of education, and concert tickets are used to assert that we can attend an event.

Many credential types are used to store [personal identifiable information](https://en.wikipedia.org/wiki/Personal_data) (PII) and therefore should be kept private, but there are many cases where credentials are practical when being published, for example, business licenses can be used to assert that a business is legit and has the authorization to operate.

In this context, we define a credential as a set of one or more claims made by an issuer and a [verifiable credential](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) (VCs) as a tamper-evident credential that has authorship that can be cryptographically verified.

This document specifies the support for verifiable credentials, both for private credentials and public credentials.

## Context

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has gained momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has captured the attention of both the private and public sectors.

The SSI approach relies on two building blocks: decentralized identifiers (DID) and verifiable credentials (VCs). This architecture decision record (ADR) describes the implementation for supporting VCs in a Cosmos SDK-based blockchain.

This ADR aims to define a foundation for the necessary components to realize the Elesto objectives while ensuring the implementation of the VCs is fully compliant with the W3C recommendations. **Successive iterations will address API ergonomics and standard compatibility issues.**

## Decision

The Elesto implementation for VCs will follow the [W3C VC data model recommendations](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) to maximize compatibility with 3rd party tools and projects.


We introduce two structures: 

- Credential Definition (CD) 
- Public Verifiable Credential (PVC)


### Credential Definition (CD)

Credential definitions are used to describe the model of a credential, that is, its structure, fields, properties, and behavior. It consists of the following fields:

- `id` - the credential definition did
- `publisherId` - the did of the publisher of the credential
- `name` - the human-readable name of the credential 
- `description` - a description of the credential usage
- `isPublic` - this field records whenever the credential can be issued on-chain
- `isActive` - this field can be used as a hint to indicate that credentials based on this schema should not be issued
- `supersededBy` - the id of a new credential definition that supersedes this definition 
- `schema` - the credential [data schema](#credential-schema) 
- `vocab` - the credential [JSON-LD vocabulary](#credential-json-ld-vocabulary)

A credential definition cannot be deleted from the state. 

#### Credential Schema

A credential [data schema](https://www.w3.org/TR/vc-data-model/#data-schemas) is a machine-readable definition of fields and data types of a credential. 

The credential module supports the [`JsonSchemaValidator2018`](https://w3c-ccg.github.io/vc-json-schemas/v1/index.html) schema format. The schema is stored as an uncompressed byte slice. 
  
#### Credential JSON-LD vocabulary

A credential [JSON-LD vocabulary or context](https://www.w3.org/TR/json-ld11/#the-context) is a semantic description of the fields of a credential. The vocabulary, in JSON format, is stored as an uncompressed byte slice.
  

### Public Verifiable Credential (PVC)

With public verifiable credential, we define a [verifiable credential](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) that is stored on-chain. 

To publish a verifiable credential on-chain, the sender must provide the credential definition that describes the credential. A credential will be published only if:

- the credential definition has been published on-chain
- the credential definition allows publication of the credential (the field `isPublic` is true)
- the credential definition status is active (the field `isActive` is true)
- the credential is conform to the schema
- the proof of the credential can be positively verified

The `supersededBy` is used as for communication purposes and does not influence the publication of a credential. 
  
Public verifiable credentials are intended for advertising information that is public domain leveraging the tamper-resistant capability of the blockchain. A public verifiable credential can be deleted from the state.

## Privacy Considerations

There is a risk that a credential definition allows publication of PII: the failsafe provided by the `isPublic` field can be misused by the credential publisher.   

## Security Considerations

It is important that both credential schema and vocabulary are stored uncompressed, compression will make the node vulnerable to a [zip bomb](https://en.wikipedia.org/wiki/Zip_bomb) attack. 
## Consequences

The Elesto chain will provide support for credential schemas and vocabularies. At the same time, support for public verifiable credentials provides the 
foundation for a network of trust.
  
### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive


- The ADR implementation improves the compatibility with the SSI identity model.
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

- [Verifiable Credentials Data Model v1.1](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/)
- [JSON-LD 1.1](https://www.w3.org/TR/2020/REC-json-ld11-20200716/)
- [Personal Identifiable Information](https://en.wikipedia.org/wiki/Personal_data)
- [General Data Protection Regulation (GDPR)](https://eur-lex.europa.eu/eli/reg/2016/679/oj)