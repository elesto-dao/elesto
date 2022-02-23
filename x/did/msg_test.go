package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------

func TestMsgCreateDidDocument_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgCreateDidDocument{}.Route(), "Route()")
}

func TestMsgCreateDidDocument_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgCreateDidDocument{}), MsgCreateDidDocument{}.Type(), "Type()")
}

func TestMsgCreateDidDocument_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgCreateDidDocument{})), MsgCreateDidDocument{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgCreateDidDocument_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgCreateDidDocument{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgCreateDidDocument{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

func TestMsgUpdateDidDocument_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgUpdateDidDocument{}.Route(), "Route()")
}

func TestMsgUpdateDidDocument_Type(t *testing.T) {
	assert.Equalf(t, TypeMsgUpdateDidDocument, MsgUpdateDidDocument{}.Type(), "Type()")
}

func TestMsgUpdateDidDocument_GetSignBytes(t *testing.T) {
	assert.Panicsf(t, func() { MsgUpdateDidDocument{}.GetSignBytes() }, "GetSignBytes()")
}

func TestMsgUpdateDidDocument_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgUpdateDidDocument{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgUpdateDidDocument{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// ADD VERIFICATION
// --------------------------

func TestMsgAddVerification_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgAddVerification{}.Route(), "Route()")
}

func TestMsgAddVerification_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgAddVerification{}), MsgAddVerification{}.Type(), "Type()")
}

func TestMsgAddVerification_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgAddVerification{})), MsgAddVerification{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgAddVerification_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgAddVerification{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgAddVerification{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// REVOKE VERIFICATION
// --------------------------

func TestMsgRevokeVerification_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgRevokeVerification{}.Route(), "Route()")
}

func TestMsgRevokeVerification_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgRevokeVerification{}), MsgRevokeVerification{}.Type(), "Type()")
}

func TestMsgRevokeVerification_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgRevokeVerification{})), MsgRevokeVerification{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgRevokeVerification_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgRevokeVerification{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgRevokeVerification{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------

func TestMsgSetVerificationRelationships_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgSetVerificationRelationships{}.Route(), "Route()")
}

func TestMsgSetVerificationRelationships_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgSetVerificationRelationships{}), MsgSetVerificationRelationships{}.Type(), "Type()")
}

func TestMsgSetVerificationRelationships_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgSetVerificationRelationships{})), MsgSetVerificationRelationships{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgSetVerificationRelationships_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgSetVerificationRelationships{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgSetVerificationRelationships{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// DELETE SERVICE
// --------------------------

func TestMsgDeleteService_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgDeleteService{}.Route(), "Route()")
}

func TestMsgDeleteService_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgDeleteService{}), MsgDeleteService{}.Type(), "Type()")
}

func TestMsgDeleteService_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgDeleteService{})), MsgDeleteService{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgDeleteService_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgDeleteService{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgDeleteService{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// ADD SERVICE
// --------------------------

func TestMsgAddService_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgAddService{}.Route(), "Route()")
}

func TestMsgAddService_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgAddService{}), MsgAddService{}.Type(), "Type()")
}

func TestMsgAddService_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgAddService{})), MsgAddService{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgAddService_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgAddService{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgAddService{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// ADD CONTROLLER
// --------------------------

func TestMsgAddController_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgAddController{}.Route(), "Route()")
}

func TestMsgAddController_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgAddController{}), MsgAddController{}.Type(), "Type()")
}

func TestMsgAddController_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgAddController{})), MsgAddController{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgAddController_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgAddController{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgAddController{Signer: "invalid"}.GetSigners() })
}

// --------------------------
// DELETE CONTROLLER
// --------------------------

func TestMsgDeleteController_Route(t *testing.T) {
	assert.Equalf(t, ModuleName, MsgDeleteController{}.Route(), "Route()")
}

func TestMsgDeleteController_Type(t *testing.T) {
	assert.Equalf(t, sdk.MsgTypeURL(&MsgDeleteController{}), MsgDeleteController{}.Type(), "Type()")
}

func TestMsgDeleteController_GetSignBytes(t *testing.T) {
	assert.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&MsgDeleteController{})), MsgDeleteController{}.GetSignBytes(), "GetSignBytes()")
}

func TestMsgDeleteController_GetSigners(t *testing.T) {
	a, err := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
	assert.NoError(t, err)
	assert.Equal(t,
		MsgDeleteController{Signer: "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"}.GetSigners(),
		[]sdk.AccAddress{a},
	)
	assert.Panics(t, func() { MsgDeleteController{Signer: "invalid"}.GetSigners() })
}
