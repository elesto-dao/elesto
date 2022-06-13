# Messages

TODO

In this section we describe the processing of the staking messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.



### MsgPublishCredentialDefinition

A `MsgPublishCredentialDefinition` is used to publish a credential definition on chain, it has the following fields

- `id` - the did string identifying the credential definition
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source

https://github.com/elesto-dao/elesto/blob/v1.0.0/proto/credential/tx.proto#L45
