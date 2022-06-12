package credential

// NewCredentialDefinitionPublishedEvent creates the definition is published
func NewCredentialDefinitionPublishedEvent(credentialDefinitionDID, publisherDID string) *CredentialDefinitionPublishedEvent {
	return &CredentialDefinitionPublishedEvent{
		CredentialDefinitionID: credentialDefinitionDID,
		PublisherID:            publisherDID,
	}
}

// NewCredentialDefinitionUpdatedEvent creates a new event for when a credential definition is updated
func NewCredentialDefinitionUpdatedEvent(credentialDefinitionDID string) *CredentialDefinitionUpdatedEvent {
	return &CredentialDefinitionUpdatedEvent{
		CredentialDefinitionID: credentialDefinitionDID,
	}
}

// NewPublicCredentialIssuedEvent creates a new event for when a credential is issued on-chain
func NewPublicCredentialIssuedEvent(credentialDefinitionDID, credentialID, issuerDID string) *PublicCredentialIssuedEvent {
	return &PublicCredentialIssuedEvent{
		CredentialDefinitionID: credentialDefinitionDID,
		CredentialID:           credentialID,
		IssuerID:               issuerDID,
	}
}
