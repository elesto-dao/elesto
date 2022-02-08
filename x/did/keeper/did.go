package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/did"
)

// SetDidDocument store a did document in the keeper, existing DID document with the
// same key will be overwritten
func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document did.DidDocument) {
	k.Set(ctx, key, did.DidDocumentKey, document, k.Marshal)
}

// GetDidDocument retrieve a DID document by its key.
// The boolean return will be false if the DID document is not found
func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (did.DidDocument, bool) {
	val, found := k.Get(ctx, key, did.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(did.DidDocument), found
}

// UnmarshalDidDocument unmarshall a did document= and check if it is empty
// ad DID document is empty if contains no context
func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	data := did.DidDocument{}
	k.Unmarshal(value, &data)
	return data, did.IsValidDIDDocument(&data)
}

func (k Keeper) SetDidMetadata(ctx sdk.Context, key []byte, meta did.DidMetadata) {
	k.Set(ctx, key, did.DidMetadataKey, meta, k.Marshal)
}

func (k Keeper) GetDidMetadata(ctx sdk.Context, key []byte) (did.DidMetadata, bool) {
	val, found := k.Get(ctx, key, did.DidMetadataKey, k.UnmarshalDidMetadata)
	return val.(did.DidMetadata), found
}

func (k Keeper) UnmarshalDidMetadata(value []byte) (interface{}, bool) {
	data := did.DidMetadata{}
	k.Unmarshal(value, &data)
	return data, did.IsValidDIDMetadata(&data)
}

// ResolveDid returning the did document and associated metadata
func (k Keeper) ResolveDid(ctx sdk.Context, didDoc did.DID) (doc did.DidDocument, meta did.DidMetadata, err error) {
	if strings.HasPrefix(didDoc.String(), did.DidKeyPrefix) {
		doc, meta, err = did.ResolveAccountDID(didDoc.String(), ctx.ChainID())
		return
	}
	doc, found := k.GetDidDocument(ctx, []byte(didDoc.String()))
	if !found {
		err = did.ErrDidDocumentNotFound
		return
	}
	meta, _ = k.GetDidMetadata(ctx, []byte(didDoc.String()))
	return
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case did.DidDocument:
		bytes = k.cdc.MustMarshal(&value)
	case did.DidMetadata:
		bytes = k.cdc.MustMarshal(&value)
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

// GetAllDidDocumentsWithCondition retrieve a list of
// did document by some arbitrary criteria. The selector filter has access
// to both the did and its metadata
func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(did did.DidDocument) bool,
) (didDocs []did.DidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		didDoc, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped := didDoc.(did.DidDocument)

		if didSelector(didTyped) {
			didDocs = append(didDocs, didTyped)
		}
	}

	return didDocs
}

// GetAllDidDocuments returns all the DidDocuments
func (k Keeper) GetAllDidDocuments(ctx sdk.Context) []did.DidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		did.DidDocumentKey,
		func(did did.DidDocument) bool { return true },
	)
}
