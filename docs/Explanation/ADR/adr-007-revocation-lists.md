# ADR 007: Revocation Lists

## Changelog

- 2022-07-04: Initial draft

## Status

DRAFT

## Abstract

Credentials are part of our daily interactions; driver's licenses are used to assert that we are capable of operating a motor vehicle, university degrees can be used to assert our level of education, and concert tickets are used to assert that we can attend an event.

It is often useful for a credential issuer of a [verfiable credential](https://www.w3.org/TR/2022/REC-vc-data-model-20220303/) to link to a location where a verifier can check to see if a credential has been revoked.

This document specifies the support for revocation lists, based on the [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report.

## Context

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has gained momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has captured the attention of both the private and public sectors.

The SSI approach relies on two building blocks: decentralized identifiers (DID) and verifiable credentials (VCs). This architecture decision record (ADR) describes a method to leverage public verifiable credentials in a Cosmos SDK-based blockchain to provide support for revocation lists.

## Decision

We introduce a credential definition that can be used by a credential issuer to publish one or more revocation lists. Each revocation lists is encoded in a [public verifiable credential](adr-006-public-verifiable-credential.md#public-verifiable-credential-pvc).

The credential issuer will be responsible to track the index of the credential that it issues. 

The revocation list model is an implementation of the [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report.


The credential schema is the following:

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
          "description": "value should be: RevocationList2020 ",
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

An example revocation list based on the above schema is:

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

that is explained by the following diagram:

![](../../assets/diagrams/out/revocationList2020.example.svg)

## Privacy Considerations

Refer to the privacy consideration of the [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report.

## Security Considerations

Refer to the security consideration of the [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/) report.

## Consequences

By leveraging the public verifiable credentials, the Elesto node offers natively the support for revocation lists.
  
### Backwards Compatibility

There are no concerns related to backward compatibility

### Positive


- The ADR implementation improves the compatibility with the SSI identity model.


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

- [Revocation List 2020](https://w3c-ccg.github.io/vc-status-rl-2020/)
