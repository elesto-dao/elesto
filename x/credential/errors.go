package credential

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/credential module sentinel errors
// TODO update error messages
var (
	ErrCredentialDefinitionFound     = sdkerrors.Register(ModuleName, 2100, "credential definition found")
	ErrCredentialDefinitionNotFound  = sdkerrors.Register(ModuleName, 2101, "credential definition not found")
	ErrCredentialDefinitionCorrupted = sdkerrors.Register(ModuleName, 2109, "credential definition corrupted")
	ErrCredentialDefinitionNotPublic = sdkerrors.Register(ModuleName, 2104, "credential definition is not public")

	ErrVerifiableCredentialNotFound = sdkerrors.Register(ModuleName, 2102, "vc not found")
	ErrVerifiableCredentialFound    = sdkerrors.Register(ModuleName, 2103, "vc found")
	ErrVerifiableCredentialIssuer   = sdkerrors.Register(ModuleName, 2105, "provided verifiable credential and did public key do not match")
	ErrInvalidProof                 = sdkerrors.Register(ModuleName, 2106, "credential proof validation error")
	ErrCredentialNotIssuable        = sdkerrors.Register(ModuleName, 2107, "credential cannot be issued on-chain")
	ErrInvalidCredential            = sdkerrors.Register(ModuleName, 2110, "credential is invalid")
	ErrCredentialSchema             = sdkerrors.Register(ModuleName, 2130, "the credential doesn't match the definition schema")
)
