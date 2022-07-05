package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         sdk.StoreKey
	paramSpace       paramtypes.Subspace
	accountKeeper    types.AccountKeeper
	bankKeeper       types.BankKeeper
	distrKeeper      types.DistributionKeeper
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace, ak types.AccountKeeper, bk types.BankKeeper, dk types.DistributionKeeper, feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	//TODO: what is paramSpace this used for?
	//set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		bankKeeper:       bk,
		distrKeeper:      dk,
		accountKeeper:    ak,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetParams returns the total set of mint parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of mint parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// SetBootstrapDateCanary sets the bootstrap date canary boolean to true.
// If overwrite is true, regardless of whether or not the canary was already set, it will be set again.
func (k Keeper) SetBootstrapDateCanary(ctx sdk.Context, value bool, overwrite bool) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.BootstrapDateCanaryKey) && !overwrite {
		return fmt.Errorf("bootstrap date canary already set")
	}

	var bValue []byte

	switch value {
	case true:
		bValue = []byte{1}
	default:
		bValue = []byte{0}
	}

	store.Set(types.BootstrapDateCanaryKey, bValue)

	return nil
}

// BootstrapDateCanarySet returns a bool representing the bootstrap date canary.
func (k Keeper) BootstrapDateCanarySet(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)

	res := store.Get(types.BootstrapDateCanaryKey)
	if len(res) == 0 || res[0] == 0 {
		return false
	}

	return true
}

// SetBootstrapDate reads the block timestamp from ctx and stores it in the
// module's store.
// Returns error if the bootstrap date was already set when overwriteDate is true.
func (k Keeper) SetBootstrapDate(ctx sdk.Context, overwriteDate bool) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.BootstrapDateKey) && !overwriteDate {
		return fmt.Errorf("bootstrap date already set")
	}

	tsBytes, err := ctx.BlockTime().MarshalBinary()
	if err != nil {
		return fmt.Errorf("cannot format block timestamp as bytes, %w", err)
	}

	store.Set(types.BootstrapDateKey, tsBytes)

	return nil
}

// BootstrapDate returns the stored bootstrap date as a time.Time object.
// Returns error if the stored byte array is malformed or not present.
func (k Keeper) BootstrapDate(ctx sdk.Context) (time.Time, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.BootstrapDateKey) {
		return time.Time{}, fmt.Errorf("bootstrap date not set")
	}

	bdBytes := store.Get(types.BootstrapDateKey)
	if bdBytes == nil {
		return time.Time{}, fmt.Errorf("bootstrap date present, but content is nil")
	}

	t := time.Time{}
	if err := t.UnmarshalBinary(bdBytes); err != nil {
		return time.Time{}, fmt.Errorf("malformed bootstrap date bytes, %w", err)
	}

	return t, nil
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddInflationToFeeCollector implements an alias call to the underlying supply keeper's
// AddInflationToFeeCollector to be used in BeginBlocker.
func (k Keeper) AddInflationToFeeCollector(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}

// CollectAmount implement an alias to call SendCoinsFromModuleToAccount
func (k Keeper) CollectAmount(ctx sdk.Context, address string, amount sdk.Coins) error {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amount)
}

// GetSupply returns the current supply on the chain for a denom
func (k Keeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	return k.bankKeeper.GetSupply(ctx, denom)
}

// FundCommunityPool funds the x/distribution community pool with amount coins, directly from
// this module's ModuleAccount.
func (k Keeper) FundCommunityPool(ctx sdk.Context, amount sdk.Coins) error {
	if amount.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	addr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if addr == nil {
		panic("the mint module account has not been set")
	}

	return k.distrKeeper.FundCommunityPool(ctx, amount, addr)
}
