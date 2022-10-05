# State

This document describes the state pertaining to:

1. [Credential Definition](./02_state.md#credential-definition)
2. [Public Verifiable Credential](./02_state.md#public-verifiable-credential)



## Credential Definition 
```go 

type CredentialDefinition struct {
	// the credential definition did
	Id string
	// the did of the publisher of the credential
	PublisherId string 
	// The credential json-ld schema.
	// The schema can be big and verbose (and expensive), but should not be compressed since
	// it is used by the msg_server to verify public credentials and. if zipped, will open
	// the node to a zip bomb attack vector
	Schema []byte 
	// the credential vocabulary
	// The vocabulary can be big and verbose (and expensive), but should not be compressed since
	// it is used by the msg_server to verify public credentials and. if zipped, will open
	// the node to a zip bomb attack vector
	Vocab []byte 
	// the human readable name of the credential, should be included
	// in the type of the issued credential
	Name string
	// the description of the credential, such as it's purpose
	Description string 
	// wherever the credential is intended for public use (on-chain) or not (off-chain)
	// if the value is false then the module will forbid the issuance of the credential on chain
	IsPublic bool
	// did of the credential should not be used anymore in favour of something else
	SupersededBy string `protobuf:"bytes,8,opt,name=supersededBy,proto3" json:"supersededBy,omitempty"`
	// the credential can be de-activated
	IsActive bool
}
```

### CredentialStatus
```go
type CredentialStatus struct {
    Id                       string 
    Type                     string 
    RevocationListIndex      int32  
    RevocationListCredential string 
}
```


### Proof
```go
type Proof struct {
	Created            string 
	Type               string 
	ProofPurpose       string 
	VerificationMethod string 
	Signature          string 
}
```

#### Source
- [credential.proto](../../../proto/credential/v1/credential.proto)

## Public Verifiable Credential
```go
type PublicVerifiableCredential struct {
    // json-ld context
    Context []string 
    // the credential id
    Id string 
    // the credential types
    Type []string
    // the DID of the issuer
    Issuer string 
    // the date-time of issuance
    IssuanceDate *time.Time
    // the date-time of expiration
    ExpirationDate *time.Time 
    // credential status for the revocation lists
    CredentialStatus *CredentialStatus 
    // the subject of the credential
    // the preferred way to handle the subject will be to use the Struct type
    // but at the moment is not supported
    // google.protobuf.Struct credentialSubject = 7;
    CredentialSubject []byte
    // One or more cryptographic proofs that can be used to detect tampering
    // and verify the authorship of a credential or presentation. The specific
    // method used for an embedded proof MUST be included using the type property.
    Proof *Proof 
}
```

#### Source
- [credential.proto](../../../proto/credential/v1/credential.proto)




