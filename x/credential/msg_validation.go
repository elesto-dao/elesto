package credential

import (
	"errors"

	"github.com/elesto-dao/elesto/v2/x/did"
)

func (m MsgPublishCredentialDefinitionRequest) ValidateBasic() error {

	if m.CredentialDefinition == nil {
		return errors.New("credential definition must be set")
	}

	if !did.IsValidDID(m.CredentialDefinition.Id) {
		return errors.New("credential definition id must be a valid DID")
	}

	if IsEmpty(m.CredentialDefinition.Name) {
		return errors.New("credential definition name must not be empty")
	}

	if !did.IsValidDID(m.CredentialDefinition.PublisherId) {
		return errors.New("publisher id must be a valid DID")
	}

	if len(m.CredentialDefinition.Schema) == 0 {
		return errors.New("schema cannot be empty")
	}

	return nil
}

func (m MsgUpdateCredentialDefinitionRequest) ValidateBasic() error {
	if !IsEmpty(m.SupersededBy) && !did.IsValidDID(m.SupersededBy) {
		return errors.New("SupersededBy must be a valid DID")
	}
	return nil
}

func (m MsgIssuePublicVerifiableCredentialRequest) ValidateBasic() error {
	if m.Credential == nil {
		return errors.New("credential must be set")
	}
	if IsEmpty(m.CredentialDefinitionID) {
		return errors.New("credential definition DID must be set")
	}
	if !did.IsValidDID(m.CredentialDefinitionID) {
		return errors.New("credential definition id must be a valid DID")
	}
	return nil
}
