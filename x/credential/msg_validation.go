package credential

import (
	"fmt"

	"github.com/elesto-dao/elesto/x/did"
)

func (msg MsgPublishCredentialDefinitionRequest) ValidateBasic() error {
	if msg.CredentialDefinition == nil {
		return fmt.Errorf("credential definition must be set")
	}

	if !did.IsValidDID(msg.CredentialDefinition.Id) {
		return fmt.Errorf("credential definition id must be a valid DID")
	}

	if IsEmpty(msg.CredentialDefinition.Name) {
		return fmt.Errorf("credential definition name must not be empty")
	}

	if !did.IsValidDID(msg.CredentialDefinition.PublisherId) {
		return fmt.Errorf("publisher id must be a valid DID")
	}

	if IsEmpty(msg.CredentialDefinition.Schema) {
		return fmt.Errorf("schema cannot be empty")
	}

	return nil
}

func (msg MsgUpdateCredentialDefinitionRequest) ValidateBasic() error {
	return nil
}

func (MsgIssuePublicVerifiableCredentialRequest) ValidateBasic() error {
	return nil
}
