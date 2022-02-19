package credentials

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------

var _ sdk.Msg = &MsgRegisterCredentialIssuerRequest{}

// NewMsgRegisterCredentialIssuerRequest creates a new MsgRegisterCredentialIssuerRequest instance
func NewMsgRegisterCredentialIssuerRequest(
	issuer *CredentialIssuer,
	revocationServiceURL string,
	signerAccount string,
) *MsgRegisterCredentialIssuerRequest {
	return &MsgRegisterCredentialIssuerRequest{
		Issuer:               issuer,
		RevocationServiceURL: revocationServiceURL,
		Signer:               signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRegisterCredentialIssuerRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgRegisterCredentialIssuerRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRegisterCredentialIssuerRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterCredentialIssuerRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD VERIFICATION
// --------------------------

var _ sdk.Msg = &MsgPublishCredentialDefinitionRequest{}

// NewMsgPublishCredentialDefinition creates a new MsgPublishCredentialDefinition instance
func NewMsgPublishCredentialDefinition(
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

// --------------------------
// REVOKE VERIFICATION
// --------------------------

var _ sdk.Msg = &MsgUpdateRevocationListRequest{}

// NewMsgUpdateRevocationListRequest creates a new MsgUpdateRevocationListRequest instance
func NewMsgUpdateRevocationListRequest(
	issuerDID string,
	revocationList *RevocationList,
	signerAccount string,
) *MsgUpdateRevocationListRequest {
	return &MsgUpdateRevocationListRequest{
		IssuerDid:  issuerDID,
		Revocation: revocationList,
		Signer:     signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgUpdateRevocationListRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpdateRevocationListRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgUpdateRevocationListRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateRevocationListRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgAddCredentialIssuanceRequest{}

// NewMsgAddCredentialIssuanceRequest creates a new MsgAddCredentialIssuanceRequest instance
func NewMsgAddCredentialIssuanceRequest(
	issuerDID string,
	credentialIssuanceDef *CredentialIssuance,
	signerAccount string,
) *MsgAddCredentialIssuanceRequest {
	return &MsgAddCredentialIssuanceRequest{
		IssuerDid:                    issuerDID,
		CredentialIssuanceDefinition: credentialIssuanceDef,
		Signer:                       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddCredentialIssuanceRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgAddCredentialIssuanceRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgAddCredentialIssuanceRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgAddCredentialIssuanceRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgRemoveCredentialIssuanceRequest{}

// NewMsgRemoveCredentialIssuanceRequest creates a new MsgRemoveCredentialIssuanceRequest instance
func NewMsgRemoveCredentialIssuanceRequest(
	issuerDID string,
	credentialIssuanceID string,
	signerAccount string,
) *MsgRemoveCredentialIssuanceRequest {
	return &MsgRemoveCredentialIssuanceRequest{
		IssuerDid: issuerDID,
		Cid:       credentialIssuanceID,
		Signer:    signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRemoveCredentialIssuanceRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgRemoveCredentialIssuanceRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRemoveCredentialIssuanceRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgRemoveCredentialIssuanceRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgAddCredentialConstraintRequest{}

// NewMsgAddCredentialConstraintRequest creates a new MsgAddCredentialConstraintRequest instance
func NewMsgAddCredentialConstraintRequest(
	issuerDID string,
	constraint *CredentialConstraint,
	signerAccount string,
) *MsgAddCredentialConstraintRequest {
	return &MsgAddCredentialConstraintRequest{
		IssuerDid:  issuerDID,
		Constraint: constraint,
		Signer:     signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddCredentialConstraintRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgAddCredentialConstraintRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgAddCredentialConstraintRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgAddCredentialConstraintRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgRemoveCredentialConstraintRequest{}

// NewMsgRemoveCredentialConstraintRequest creates a new MsgRemoveCredentialConstraintRequest instance
func NewMsgRemoveCredentialConstraintRequest(
	issuerDID string,
	constraintID string,
	signerAccount string,
) *MsgRemoveCredentialConstraintRequest {
	return &MsgRemoveCredentialConstraintRequest{
		IssuerDid:    issuerDID,
		ConstraintId: constraintID,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRemoveCredentialConstraintRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgRemoveCredentialConstraintRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRemoveCredentialConstraintRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgRemoveCredentialConstraintRequest) GetSigners() []sdk.AccAddress {
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
	signerAccount string,
) *MsgIssuePublicVerifiableCredentialRequest {
	return &MsgIssuePublicVerifiableCredentialRequest{
		Credential: credential,
		Signer:     signerAccount,
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
