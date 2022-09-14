package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/elesto-dao/elesto/v3/x/did"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// MarshalFn is a generic function to marshal bytes
type MarshalFn func(o codec.ProtoMarshaler) []byte

// Keeper defines the Keeper for the did module
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey
}

// NewKeeper instanstiates the Keeper for the did module
func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

// Logger defines the logger for the did module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", did.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context,
	key []byte,
	prefix []byte,
	i codec.ProtoMarshaler,
	marshal MarshalFn,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Delete deletes a value form the store
func (k Keeper) Delete(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(prefix, key...))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
	unmarshal UnmarshalFn,
) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// Has returns a bool if item is in the store
func (k Keeper) Has(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) (found bool) {
	store := ctx.KVStore(k.storeKey)
	return store.Has(append(prefix, key...))
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(
	ctx sdk.Context,
	prefix []byte,
) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
