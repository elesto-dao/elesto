package credential

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/did"
)

// --------------------------
// CREDENTIAL DEFINITIONS
// --------------------------

var _ sdk.Msg = &MsgPublishCredentialDefinitionRequest{}

// NewMsgPublishCredentialDefinitionRequest creates a new MsgPublishCredentialDefinition instance
func NewMsgPublishCredentialDefinitionRequest(
	credentialDefinition *CredentialDefinition,
	signerAccount string,
) *MsgPublishCredentialDefinitionRequest {
	return &MsgPublishCredentialDefinitionRequest{
		CredentialDefinition: credentialDefinition,
		Signer:               signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgPublishCredentialDefinitionRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgPublishCredentialDefinitionRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgPublishCredentialDefinitionRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgPublishCredentialDefinitionRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgUpdateCredentialDefinitionRequest{}

// NewMsgUpdateCredentialDefinitionRequest creates a new MsgUpdateCredentialDefinitionRequest instance
func NewMsgUpdateCredentialDefinitionRequest(
	isActive bool, supersededBy did.DID,
	signerAccount string,
) *MsgUpdateCredentialDefinitionRequest {
	return &MsgUpdateCredentialDefinitionRequest{
		Active:       isActive,
		SupersededBy: supersededBy.String(),
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgUpdateCredentialDefinitionRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpdateCredentialDefinitionRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgUpdateCredentialDefinitionRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateCredentialDefinitionRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgIssuePublicVerifiableCredentialRequest{}

// NewMsgIssuePublicVerifiableCredentialRequest creates a new MsgIssuePublicVerifiableCredentialRequest instance
func NewMsgIssuePublicVerifiableCredentialRequest(
	credential *PublicVerifiableCredential,
	definitionID string,
	signerAccount sdk.AccAddress,
) *MsgIssuePublicVerifiableCredentialRequest {
	return &MsgIssuePublicVerifiableCredentialRequest{
		Credential:             credential,
		CredentialDefinitionID: definitionID,
		Signer:                 signerAccount.String(),
	}
}

// Route implements sdk.Msg
func (MsgIssuePublicVerifiableCredentialRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgIssuePublicVerifiableCredentialRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgIssuePublicVerifiableCredentialRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgIssuePublicVerifiableCredentialRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
