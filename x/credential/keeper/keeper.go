package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// MarshalFn is a generic function to marshal bytes
type MarshalFn func(o codec.ProtoMarshaler) []byte

type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey
	did      credential.DidKeeper
	account  credential.AccountKeeper
}

func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey, did credential.DidKeeper, account credential.AccountKeeper) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		did:      did,
		account:  account,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", credential.ModuleName))
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

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(
	ctx sdk.Context,
	prefix []byte,
) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k Keeper) Exists(
	ctx sdk.Context,
	prefix []byte,
	key []byte,
) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(append(prefix, key...))
	return bz != nil
}
