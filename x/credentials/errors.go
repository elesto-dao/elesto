package credentials

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
// TODO update error messages
var (
	ErrCredentialDefinitionFound    = sdkerrors.Register(ModuleName, 2100, "credential definition found")
	ErrCredentialDefinitionNotFound = sdkerrors.Register(ModuleName, 2101, "credential definition not found")
)
