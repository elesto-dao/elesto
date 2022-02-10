package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------

var _ sdk.Msg = &MsgCreateDidDocument{}

// NewMsgCreateDidDocument creates a new MsgCreateDidDocument instance
func NewMsgCreateDidDocument(
	id string,
	verifications []*Verification,
	services []*Service,
	signerAccount string,
) *MsgCreateDidDocument {
	return &MsgCreateDidDocument{
		Id:            id,
		Verifications: verifications,
		Services:      services,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgCreateDidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgCreateDidDocument) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgCreateDidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgCreateDidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

// msg types
const (
	TypeMsgUpdateDidDocument = "update-did"
)

func NewMsgUpdateDidDocument(
	didDoc *DidDocument,
	signerAccount string,
) *MsgUpdateDidDocument {
	return &MsgUpdateDidDocument{
		Doc:    didDoc,
		Signer: signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgUpdateDidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgUpdateDidDocument) Type() string {
	return TypeMsgUpdateDidDocument
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (MsgUpdateDidDocument) GetSignBytes() []byte {
	panic("TODO: needed in simulations for fuzz testing")
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateDidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD VERIFICATION
// --------------------------

var _ sdk.Msg = &MsgAddVerification{}

// NewMsgAddVerification creates a new MsgAddVerification instance
func NewMsgAddVerification(
	id string,
	verification *Verification,
	signerAccount string,
) *MsgAddVerification {
	return &MsgAddVerification{
		Id:           id,
		Verification: verification,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgAddVerification) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgAddVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgAddVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// REVOKE VERIFICATION
// --------------------------

var _ sdk.Msg = &MsgRevokeVerification{}

// NewMsgRevokeVerification creates a new MsgRevokeVerification instance
func NewMsgRevokeVerification(
	id string,
	methodID string,
	signerAccount string,
) *MsgRevokeVerification {
	return &MsgRevokeVerification{
		Id:       id,
		MethodId: methodID,
		Signer:   signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRevokeVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgRevokeVerification) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRevokeVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------
// msg types
const (
	TypeMsgSetVerificationRelationships = "set-verification-relationships"
)

func NewMsgSetVerificationRelationships(
	id string,
	methodID string,
	relationships []string,
	signerAccount string,
) *MsgSetVerificationRelationships {
	return &MsgSetVerificationRelationships{
		Id:            id,
		MethodId:      methodID,
		Relationships: relationships,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgSetVerificationRelationships) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSetVerificationRelationships) Type() string {
	return TypeMsgSetVerificationRelationships
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (MsgSetVerificationRelationships) GetSignBytes() []byte {
	panic("TODO: needed in simulations for fuzz testing")
}

// GetSigners implements sdk.Msg
func (msg MsgSetVerificationRelationships) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD SERVICE
// --------------------------

var _ sdk.Msg = &MsgAddService{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddService(
	id string,
	service *Service,
	signerAccount string,
) *MsgAddService {
	return &MsgAddService{
		Id:          id,
		ServiceData: service,
		Signer:      signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgAddService) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgAddService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgAddService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE SERVICE
// --------------------------

func NewMsgDeleteService(
	id string,
	serviceID string,
	signerAccount string,
) *MsgDeleteService {
	return &MsgDeleteService{
		Id:        id,
		ServiceId: serviceID,
		Signer:    signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgDeleteService) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgDeleteService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD CONTROLLER
// --------------------------

// msg types
const (
	TypeMsgAddController = "add-controller"
)

func NewMsgAddController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgAddController {
	return &MsgAddController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddController) Type() string {
	return TypeMsgAddController
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (MsgAddController) GetSignBytes() []byte {
	panic("TODO: needed in simulations for fuzz testing")
}

// GetSigners implements sdk.Msg
func (msg MsgAddController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE CONTROLLER
// --------------------------

// msg types
const (
	TypeMsgDeleteController = "delete-controller"
)

func NewMsgDeleteController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgDeleteController {
	return &MsgDeleteController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

//// Route implements sdk.Msg
func (MsgDeleteController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteController) Type() string {
	return TypeMsgDeleteController
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (MsgDeleteController) GetSignBytes() []byte {
	panic("TODO: needed in simulations for fuzz testing")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
