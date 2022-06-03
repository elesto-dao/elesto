package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/did"
)

// SetDidDocument store a did document in the keeper, existing DID document with the
// same key will be overwritten
func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document did.DidDocument) {
	k.Set(ctx, key, did.DidDocumentKey, &document, k.cdc.MustMarshal)
}

// GetDidDocument retrieve a DID document by its key.
// The boolean return will be false if the DID document is not found
func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (did.DidDocument, bool) {
	val, found := k.Get(ctx, key, did.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(did.DidDocument), found
}

// HasDidDocument checks if a DID document is in the store  by its key.
// The boolean return will be false if the DID document is not found
func (k Keeper) HasDidDocument(ctx sdk.Context, key []byte) bool {
	found := k.Has(ctx, key, did.DidDocumentKey)
	return found
}

// UnmarshalDidDocument unmarshall a did document and check if it is empty
// ad DID document is empty if contains no context
func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	data := did.DidDocument{}
	k.cdc.MustUnmarshal(value, &data)
	return data, did.IsValidDIDDocument(&data)
}

// ResolveDid returns the did document if its a did:key or did:cosmos
// this function is used to resolve ephemeral dids
func (k Keeper) ResolveDid(
	ctx sdk.Context,
	didDoc did.DID,
) (doc did.DidDocument, err error) {
	if strings.HasPrefix(didDoc.String(), did.DidKeyPrefix) {
		doc, err = did.ResolveAccountDID(didDoc.String(), ctx.ChainID())
		return
	}
	doc, found := k.GetDidDocument(ctx, []byte(didDoc.String()))
	if !found {
		err = did.ErrDidDocumentNotFound
		return
	}
	return
}

// GetAllDidDocumentsWithCondition retrieve a list of
// did document by some arbitrary criteria. The selector filter has access
// to both the did and its metadata
func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(did *did.DidDocument) bool,
) (didDocs []*did.DidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		didDoc, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped := didDoc.(did.DidDocument)

		if didSelector(&didTyped) {
			didDocs = append(didDocs, &didTyped)
		}
	}

	return didDocs
}

// GetAllDidDocuments returns all the DidDocuments
func (k Keeper) GetAllDidDocuments(ctx sdk.Context) []*did.DidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		did.DidDocumentKey,
		func(did *did.DidDocument) bool { return true },
	)
}
