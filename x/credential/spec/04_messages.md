# Messages


In this section we describe the processing of the messsages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./03_state_transitions.md) section.



### MsgPublishCredentialDefinition

A `MsgPublishCredentialDefinition` is used to publish a credential definition on chain, it has the following fields

- `credentialDefinition` - the credential definition to publish
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source

 - [credentialDefinition](./02_state.md#credential-definition)
 - [tx.proto](../../../proto/credential/v1/tx.proto)


### MsgUpdateCredentialDefinition

A `MsgUpdateCredentialDefinition` is used to update a credential definition

- `credentialDefinitionID` - id of credential definition that has to be updated
- `active` - update the active field
- `supersededBy` - update the SupersededBy field
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source
- [tx.proto](../../../proto/credential/v1/tx.proto)

### MsgIssuePublicVerifiableCredential
A `MsgIssuePublicVerifiableCredential` is used to publish public verifiable credential for a credential definition
- `credential` - public verifiable credential to publish
- `credentialDefinitionID` - credential definition id for which public verifiable credential is published
- `signer` - a string containing the cosmos address of the private key signing the transaction


#### Source
- [tx.proto](../../../proto/credential/v1/tx.proto)
- [credential.proto](../../../proto/credential/v1/credential.proto)
