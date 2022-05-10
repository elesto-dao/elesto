package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/credentials"
)

// SetCredentialIssuer persist a credential issuer to the store. The credential issuer DID is used as key.
func (k Keeper) SetCredentialIssuer(ctx sdk.Context, issuer credentials.CredentialIssuer) {
	k.Set(ctx, []byte(issuer.Did), credentials.CredentialIssuerKey, &issuer, k.cdc.MustMarshal)
}

// GetCredentialIssuer retrieve a credential issuer by its key.
// The boolean return will be false if the credential issuer is not found
func (k Keeper) GetCredentialIssuer(ctx sdk.Context, key []byte) (credentials.CredentialIssuer, bool) {
	val, found := k.Get(ctx, key, credentials.CredentialIssuerKey, func(value []byte) (interface{}, bool) {
		var data credentials.CredentialIssuer
		ok := k.Unmarshal(value, &data)
		return data, ok
	})
	return val.(credentials.CredentialIssuer), found
}

// SetCredentialDefinition persist a credential definition to the store. The credential definition ID is used as key.
func (k Keeper) SetCredentialDefinition(ctx sdk.Context, def *credentials.CredentialDefinition) {
	k.Set(ctx, []byte(def.Id), credentials.CredentialDefinitionKey, def, k.cdc.MustMarshal)
}

// GetCredentialDefinition retrieve a credential definition by its key.
// The boolean return will be false if the credential definition is not found
func (k Keeper) GetCredentialDefinition(ctx sdk.Context, key string) (credentials.CredentialDefinition, bool) {
	val, found := k.Get(ctx, []byte(key), credentials.CredentialDefinitionKey, func(value []byte) (interface{}, bool) {
		var data credentials.CredentialDefinition
		ok := k.Unmarshal(value, &data)
		return data, ok
	})
	return val.(credentials.CredentialDefinition), found
}

// Unmarshal from byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}
