package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/elesto-dao/elesto/v3/x/credential"
)

// SetCredentialDefinition persist a credential definition to the store. The credential definition ID is used as key.
func (k Keeper) SetCredentialDefinition(ctx sdk.Context, def *credential.CredentialDefinition) {
	k.Set(ctx, credential.CredentialDefinitionKey, []byte(def.Id), def, k.cdc.MustMarshal)
}

func (k Keeper) AllowPublicCredential(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(credential.PublicCredentialAllowKey, []byte(id)...), []byte(id))
}

func (k Keeper) RemovePublicCredentialFromAllowedList(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(credential.PublicCredentialAllowKey, []byte(id)...))
}

func (k Keeper) IsPublicCredentialDefinitionAllowed(ctx sdk.Context, id string) bool {
	return k.Exists(ctx, credential.PublicCredentialAllowKey, []byte(id))
}

// GetCredentialDefinition retrieve a credential definition by its key.
// The boolean return will be false if the credential definition is not found
func (k Keeper) GetCredentialDefinition(ctx sdk.Context, key string) (credential.CredentialDefinition, bool) {
	val, found := k.Get(ctx, credential.CredentialDefinitionKey, []byte(key), func(value []byte) (interface{}, bool) {
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
	k.Set(ctx, credential.PublicCredentialKey, []byte(pc.Id), pc, k.cdc.MustMarshal)
}

// GetPublicCredential retrieve a public verifiable credential by its key.
// The boolean return will be false if the credential is not found
func (k Keeper) GetPublicCredential(ctx sdk.Context, key string) (credential.PublicVerifiableCredential, bool) {
	val, found := k.Get(ctx, credential.PublicCredentialKey, []byte(key), func(value []byte) (interface{}, bool) {
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
func (k Keeper) GetCredentialDefinitionsWithFilter(ctx sdk.Context, paginate *query.PageRequest, filter func(credentialDefinition *credential.CredentialDefinition) bool) (cds []*credential.CredentialDefinition, pageRes *query.PageResponse, err error) {
	store := ctx.KVStore(k.storeKey)
	cdStore := prefix.NewStore(store, credential.CredentialDefinitionKey)

	pageRes, err = query.Paginate(cdStore, paginate, func(key []byte, value []byte) error {
		var cd credential.CredentialDefinition
		k.cdc.MustUnmarshal(value, &cd)
		if filter(&cd) {
			cds = append(cds, &cd)
		}
		return nil
	})
	return
}

// GetPublicCredentialWithFilter retrieve a list of verifiable credentials
// filtered by properties
func (k Keeper) GetPublicCredentialWithFilter(ctx sdk.Context, pagination *query.PageRequest, filter func(verifiableCredential *credential.PublicVerifiableCredential) bool) (pvcs []*credential.PublicVerifiableCredential, pageRes *query.PageResponse, err error) {
	store := ctx.KVStore(k.storeKey)
	pcStore := prefix.NewStore(store, credential.PublicCredentialKey)

	pageRes, err = query.Paginate(pcStore, pagination, func(key []byte, value []byte) error {
		var pvc credential.PublicVerifiableCredential
		k.cdc.MustUnmarshal(value, &pvc)
		if filter(&pvc) {
			pvcs = append(pvcs, &pvc)
		}
		return nil
	})
	return
}

func (k Keeper) GetAllowedCredentialDefinitions(ctx sdk.Context, req *query.PageRequest) (cds []*credential.CredentialDefinition, pageRes *query.PageResponse, err error) {
	store := ctx.KVStore(k.storeKey)
	cdStore := prefix.NewStore(store, credential.PublicCredentialAllowKey)
	pageRes, err = query.Paginate(cdStore, req, func(key []byte, value []byte) error {
		id := string(value)
		cd, found := k.GetCredentialDefinition(ctx, id)
		if !found {
			panic(fmt.Sprintf("credential definition with allowed id %s not found", id))
		}
		cds = append(cds, &cd)

		return nil
	})
	return
}
