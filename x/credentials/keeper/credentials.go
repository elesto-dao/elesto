package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/credentials"
)

func (k Keeper) SetCredentialIssuer(ctx sdk.Context, issuer credentials.CredentialIssuer) {
	k.Set(ctx, []byte(issuer.Did), credentials.CredentialIssuerKey, issuer, k.Marshal)
}

// GetCredentialIssuer retrieve a DID document by its key.
// The boolean return will be false if the DID document is not found
func (k Keeper) GetCredentialIssuer(ctx sdk.Context, key []byte) (credentials.CredentialIssuer, bool) {
	val, found := k.Get(ctx, key, credentials.CredentialIssuerKey, k.UnmarshalCredentialIssuer)
	return val.(credentials.CredentialIssuer), found
}

// UnmarshalCredentialIssuer unmarshall a did document= and check if it is empty
// ad DID document is empty if contains no context
func (k Keeper) UnmarshalCredentialIssuer(value []byte) (interface{}, bool) {
	data := credentials.CredentialIssuer{}
	ok := k.Unmarshal(value, &data)
	return data, ok
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case credentials.CredentialIssuer:
		bytes = k.cdc.MustMarshal(&value)
	case credentials.PublicVerifiableCredential:
		bytes = k.cdc.MustMarshal(&value)
	default:
		panic("serialization not supported")
	}
	return
}

// Unmarshal unmarshal a byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}
