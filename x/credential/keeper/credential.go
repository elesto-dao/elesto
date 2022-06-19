package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

// SetCredentialDefinition persist a credential definition to the store. The credential definition ID is used as key.
func (k Keeper) SetCredentialDefinition(ctx sdk.Context, def *credential.CredentialDefinition) {
	k.Set(ctx, []byte(def.Id), credential.CredentialDefinitionKey, def, k.cdc.MustMarshal)
}

// GetCredentialDefinition retrieve a credential definition by its key.
// The boolean return will be false if the credential definition is not found
func (k Keeper) GetCredentialDefinition(ctx sdk.Context, key string) (credential.CredentialDefinition, bool) {
	val, found := k.Get(ctx, []byte(key), credential.CredentialDefinitionKey, func(value []byte) (interface{}, bool) {
		var data credential.CredentialDefinition
		ok := k.Unmarshal(value, &data)
		return data, ok
	})
	return val.(credential.CredentialDefinition), found
}

// GetCredentialDefinitions list credential definitions
func (k Keeper) GetCredentialDefinitions(ctx sdk.Context, req *credential.QueryCredentialDefinitionsRequest) (cds []*credential.CredentialDefinition, pageRes *query.PageResponse, err error) {
	store := ctx.KVStore(k.storeKey)
	cdStore := prefix.NewStore(store, credential.CredentialDefinitionKey)
	pageRes, err = query.Paginate(cdStore, req.Pagination, func(key []byte, value []byte) error {
		var cd credential.CredentialDefinition
		k.cdc.MustUnmarshal(value, &cd)
		cds = append(cds, &cd)
		return nil
	})
	return
}

// SetPublicCredential persist a public verifiable credential to the store. The credential ID is used as key
func (k Keeper) SetPublicCredential(ctx sdk.Context, pc *credential.PublicVerifiableCredential) {
	k.Set(ctx, []byte(pc.Id), credential.PublicCredentialKey, pc, k.cdc.MustMarshal)
}

// GetPublicCredential retrieve a public verifiable credential by its key.
// The boolean return will be false if the credential is not found
func (k Keeper) GetPublicCredential(ctx sdk.Context, key string) (credential.PublicVerifiableCredential, bool) {
	val, found := k.Get(ctx, []byte(key), credential.PublicCredentialKey, func(value []byte) (interface{}, bool) {
		var data credential.PublicVerifiableCredential
		ok := k.Unmarshal(value, &data)
		return data, ok
	})
	return val.(credential.PublicVerifiableCredential), found
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

// GetCredentialDefinitionsWithFilter retrieve a list of credential definitions with filter
func (k Keeper) GetCredentialDefinitionsWithFilter(ctx sdk.Context, filter func(credentialDefinition *credential.CredentialDefinition) bool) []*credential.CredentialDefinition {
	var cds []*credential.CredentialDefinition

	iterator := k.GetAll(ctx, credential.CredentialDefinitionKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var cd credential.CredentialDefinition
		k.cdc.MustUnmarshal(iterator.Value(), &cd)
		if filter(&cd) {
			cds = append(cds, &cd)
		}
	}
	return cds
}

// GetPublicCredentialWithFilter retrieve a list of verifiable credentials
// filtered by properties
func (k Keeper) GetPublicCredentialWithFilter(ctx sdk.Context, filter func(verifiableCredential *credential.PublicVerifiableCredential) bool) []*credential.PublicVerifiableCredential {
	var pvcs []*credential.PublicVerifiableCredential

	iterator := k.GetAll(ctx, credential.PublicCredentialKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var pvc credential.PublicVerifiableCredential
		k.cdc.MustUnmarshal(iterator.Value(), &pvc)
		if filter(&pvc) {
			pvcs = append(pvcs, &pvc)
		}
	}
	return pvcs
}
