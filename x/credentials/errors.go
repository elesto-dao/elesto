package credentials

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
// TODO update error messages
var (
	ErrCredentialIssuerNotFound     = sdkerrors.Register(ModuleName, 2100, "credential issuer not found")
	ErrCredentialIssuerFound        = sdkerrors.Register(ModuleName, 2101, "credential issuer found")
	ErrCredentialDefinitionFound    = sdkerrors.Register(ModuleName, 2102, "credential definition found")
	ErrCredentialDefinitionNotFound = sdkerrors.Register(ModuleName, 2103, "credential definition not found")
)
