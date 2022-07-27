# ADR 007: Revocation Lists

## Changelog

- 2022-07-04: Initial draft

## Status

DRAFT

## Abstract

Credentials are part of our daily interactions; driver's licenses are used to assert that we are capable of operating a motor vehicle, university degrees can be used to assert our level of education, and concert tickets are used to assert that we can attend an event.

It is helpful for a credential issuer of a [verifiable credential](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) (VC) to provide a location where a verifier can check to see if a credential has been revoked.

This document specifies support for a native implementation of revocation lists that are based on the W3C Credentials Community Group [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report. This report details a strong privacy-preserving, space-efficient, and high-performance mechanism for publishing the revocation status of VCs.

## Context

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has gained momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has captured the attention of the private and public sectors.

The SSI approach relies on two building blocks: decentralized identifiers (DID) and VCs. This architecture decision record (ADR) describes a method to leverage public VCs in Cosmos SDK-based blockchains to provide support for revocation lists. 

## Decision

By leveraging public verifiable credentials (PVCs) and a credential definition schema, the Elesto node will offer native support for revocation lists. Privacy and performance must be considered when designing, publishing, and processing revocation lists.

This ADR introduces a [credential definition](adr-006-public-verifiable-credential.md#credential-definition-cd) for use by a credential issuer to publish one or more revocation lists. Each revocation list is encoded in a [public verifiable credential (PVC)](adr-006-public-verifiable-credential.md#public-verifiable-credential-pvc).

The revocation list model is an implementation of the W3C Credentials Community Group [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) privacy-preserving, space-efficient, and high-performance mechanism for publishing the revocation status of VCs.

In this model, the credential issuer assigns an arbitrary positive number to each credential that it issues: the number is the index in the revocation list encoded in the [public verifiable credential](adr-006-public-verifiable-credential.md#public-verifiable-credential-pvc).

The credential schema, included in the revocation list credential definition, is defined by this JSON-formatted schema:

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://elesto.id/revocation-list2020/1.0/json-schema.json",
  "$metadata": {
    "slug": "revocation-list2020",
    "version": "1.0"
  },
  "title": "RevocationList2020",
  "description": "RevocationList2020 - A privacy-preserving mechanism for revoking Verifiable Credentials",
  "type": "object",
  "required": [
    "@context",
    "type",
    "issuer",
    "issuanceDate",
    "credentialSubject"
  ],
  "properties": {
    "@context": {
      "type": [
        "string",
        "array",
        "object"
      ]
    },
    "id": {
      "type": "string",
      "format": "uri"
    },
    "type": {
      "type": [
        "string",
        "array"
      ],
      "items": {
        "type": "string"
      }
    },
    "issuer": {
      "type": [
        "string",
        "object"
      ],
      "format": "uri",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uri"
        }
      }
    },
    "issuanceDate": {
      "type": "string",
      "format": "date-time"
    },
    "expirationDate": {
      "type": "string",
      "format": "date-time"
    },
    "credentialSubject": {
      "type": "object",
      "required": [
        "id",
        "type",
        "encodedList"
      ],
      "properties": {
        "id": {
          "title": "ID",
          "description": "The revocation list ID",
          "type": "string",
          "format": "uri"
        },
        "type": {
          "title": "Type",
          "description": "value must be: RevocationList2020 ",
          "type": "string"
        },
        "encodedList": {
          "title": "encodedList",
          "description": "base64 endcoded string of the zlib compressed bitstring",
          "type": "string"
        }
      }
    },
    "credentialSchema": {
      "type": "object",
      "required": [
        "id",
        "type"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uri"
        },
        "type": {
          "type": "string"
        }
      }
    }
  }
}
```

### Example

For example, a revocation list based on the credential definition schema looks like:

```json
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://w3id.org/vc-revocation-list-2020/v1"
  ],
  "id": "https://regulator/credentials/status/001",
  "type": ["VerifiableCredential", "RevocationList2020Credential"],
  "issuer": "did:cosmos:elesto:example-credential-issuer",
  "issuanceDate": "2020-04-05T14:27:40Z",
  "credentialSubject": {
    "id": "https://regulator/credentials/status/001#list",
    "type": "RevocationList2020",
    "encodedList": "H4sIAAAAAAAAA-3BMQEAAADCoPVPbQsvoAAAAAAAAAAAAAAAAP4GcwM92tQwAAA"
  }
}
```

As shown in the following diagram:

![](../../assets/diagrams/out/revocationList2020.example.svg)

## Privacy Considerations

Refer to the privacy considerations of the W3C [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/#privacy-considerations) report.

## Security Considerations

Refer to the security considerations of the W3C [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/#security-considerations) report.

## Consequences
  
By leveraging the public verifiable credentials, the Elesto node offers native support for revocation lists. Revocation lists are stored as credentials in the node state, within the credential module keeper. 

### Backward Compatibility

There are no concerns related to backward compatibility.

### Positive

- The revocation list support and implementation improves the compatibility of the Elesto node with the SSI identity model.

### Negative

N/A

### Neutral

The credential issuer will be responsible for creating, maintaining, and tracking an index of the credentials that it issues; it is the credential issuer's responsibility to know what the next unassigned credential index associated with a revocation list is.
## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section summarizes issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.

Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

N/A

## References

- [W3C Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report
- [W3C Recommendation Verifiable Credentials Data Model v1.1](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/)
