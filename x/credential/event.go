package credential

func NewCredentialIssuerRegisteredEvent(issuerDID string) *CredentialIssuerRegisteredEvent {
	return &CredentialIssuerRegisteredEvent{
		IssuerId: issuerDID,
	}
}

// NewCredentialDefinitionPublishedEvent constructs a new did_created sdk.Event
func NewCredentialDefinitionPublishedEvent(credentialDefinitionDID, publisherDID string) *CredentialDefinitionPublishedEvent {
	return &CredentialDefinitionPublishedEvent{
		CredentialDefinitionId: credentialDefinitionDID,
		PublisherId:            publisherDID,
	}
}

// NewPublicCredentialIssuedEvent constructs a new did_created sdk.Event
func NewPublicCredentialIssuedEvent(credentialDefinitionDID, credentialID, issuerDID, holderDID string) *PublicCredentialIssuedEvent {
	return &PublicCredentialIssuedEvent{
		CredentialDefinitionId: credentialDefinitionDID,
		CredentialId:           credentialID,
		IssuerId:               issuerDID,
		HolderId:               holderDID,
	}
}
