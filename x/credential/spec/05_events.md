# Events Emitted

```go
type CredentialDefinitionPublishedEvent struct {
	CredentialDefinitionID string 
	PublisherID            string 
}

type CredentialDefinitionUpdatedEvent struct {
        CredentialDefinitionID string 
}

type PublicCredentialIssuedEvent struct {
        CredentialDefinitionID string 
        CredentialID           string 
        IssuerID               string 
}
```
###source
- [event.proto](../../../proto/credential/v1/event.proto)

