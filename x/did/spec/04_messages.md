# Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.


### Verification 

A verification message represent a combination of a verification method and a set of verification relationships. It has the following fields:

- `relationships` - a list of strings identifying the verification relationship for the verification method
- `method` - a [verification method object](02_state.md#verification_method) 
- `context` - a list of strings identifying additional [json ld contexts](https://json-ld.org/spec/latest/json-ld/#the-context)


#### Source 
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L32



### MsgCreateDidDocument

A `MsgCreateDidDocument` is used to create a new DID document, it has the following fields

- `id` - the did string identifying the DID document
- `controller` - a list of did that are controllers of the DID document
- `verifications` - a list of [verification](04_messages.md#verification) for the DID document
- `services` - a list of [services](02_state.md#service) for the DID document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source

https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L45

### MsgUpdateDidDocument

The `MsgUpdateDidDocument` is used to update a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `controller` - a list of did that are controllers of the DID document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L58
### MsgAddVerification

The `MsgAddVerification` is used to add new [verification methods](https://w3c.github.io/did-core/#verification-methods) and [verification relationships](https://w3c.github.io/did-core/#verification-relationships) to a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `verification` - the [verification](04_messages.md#verification) to add to the DID document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L73

### MsgSetVerificationRelationships

The `MsgSetVerificationRelationships` is used to overwrite the [verification relationships](https://w3c.github.io/did-core/#verification-relationships) for a [verification methods](https://w3c.github.io/did-core/#verification-methods) of a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `method_id` - a string containing the unique identifier of the verification method within the DID document.
- `relationships` - a list of strings identifying the verification relationship for the verification method
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L84
### MsgRevokeVerification

The `MsgRevokeVerification` is used to remove a [verification method](https://w3c.github.io/did-core/#verification-methods) and related [verification relationships](https://w3c.github.io/did-core/#verification-relationships) from a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `method_id` - a string containing the unique identifier of the verification method within the DID document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L96
### MsgAddService

The `MsgAddService` is used to add a [service](https://w3c.github.io/did-core/#services) to a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `service_data` - the [service](02_state.md#service) object to add to the DID document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L111
### MsgDeleteService

The `MsgDeleteService` is used to remove a [service](https://w3c.github.io/did-core/#services) from a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `service_id` - the unique id of the [service](02_state.md#service) in the DID document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/tx.proto#L122

### MsgAddController

The `MsgAddController` is used to add a [controller](https://w3c.github.io/did-core/#did-controller) to a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `controllerDid` - the controller DID to add to the DID document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:

-- 

### MsgDeleteController

The `MsgDeleteController` is used to remove a [controller](https://w3c.github.io/did-core/#did-controller) from a DID document. It has the following fields:

- `id` - the did string identifying the DID document
- `controllerDid` - the controller DID to add to the DID document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:

--


### QueryDidDocumentRequest

The `QueryDidDocumentRequest` is used to resolve a DID document. That is, to retrieve  a DID document from its id. It has the following fields:

- `id` - the did string identifying the DID document

#### Source: 
https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/did/query.proto#L45