# State Transitions


### PublishCredentialDefinition

Example: 

```
elestod tx credential publish-credential-definition \
 
 --from regulator --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet \
 
 example-definition-id example-credential schema.json vocab.json
```

```javascript
/* gRPC message */
IssuePublishCredentialDefinition(
    MsgPublishCredentialDefinition(
        
    )
)

/* Resolved Credential Definition */
{
    
}

```

##### Implementation Source

- server: [msg_server.go](../keeper/msg_server.go)
- client: [tx.go](../client/cli/tx.go)

### 

### IssuePublicVerifiableCredential

Example:

```
elestod tx credential issue-public-credential \
 
 --from regulator --node https://elesto.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet \
    example-definition-id credential.json

```

```javascript
/* gRPC message */
IssuePublicVerifiableCredential(
    MsgIssuePublicVerifiableCredential(
        
    )
)

/* Issued PVC */ 
{
}

```

##### Implementation Source

- server: [msg_server.go](../keeper/msg_server.go)
- client: [tx.go](../client/cli/tx.go)

### UpdateCredentialDefinition

```javascript
/* gRPC message */
UpdateCredentialDefinition(
    MsgUpdateCredentialDefinition(
        
    )
)

/* Resolved Credential Definition */ 
{
}

```
##### Implementation Source

- server: [msg_server.go](../keeper/msg_server.go)
- client: [tx.go](../client/cli/tx.go)


