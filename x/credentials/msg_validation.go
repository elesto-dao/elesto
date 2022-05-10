package credentials

import (
	"fmt"
	"strings"

	"github.com/elesto-dao/elesto/x/did"
)

func isEmpty(v string) bool {
	return strings.TrimSpace(v) == ""
}

func (msg MsgPublishCredentialDefinitionRequest) ValidateBasic() error {
	if !did.IsValidDID(msg.CredentialDefinition.Id) {
		return fmt.Errorf("credential definition id must be a valid DID")
	}

	if isEmpty(msg.CredentialDefinition.Name) {
		return fmt.Errorf("credential definition name must not be empty")
	}

	if !did.IsValidDID(msg.CredentialDefinition.PublisherId) {
		return fmt.Errorf("publisher id must be a valid DID")
	}

	if isEmpty(msg.CredentialDefinition.Schema) {
		return fmt.Errorf("schema cannot be empty")
	}

	return nil
}

func (MsgRegisterCredentialIssuerRequest) ValidateBasic() error {
	return nil
}

func (MsgUpdateCredentialDefinitionRequest) ValidateBasic() error {
	return nil
}

func (MsgAddCredentialIssuanceRequest) ValidateBasic() error {
	return nil
}

func (MsgRemoveCredentialIssuanceRequest) ValidateBasic() error {
	return nil
}

func (MsgAddCredentialConstraintRequest) ValidateBasic() error {
	return nil
}

func (MsgRemoveCredentialConstraintRequest) ValidateBasic() error {
	return nil
}

func (MsgIssuePublicVerifiableCredentialRequest) ValidateBasic() error {
	return nil
}
