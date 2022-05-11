package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/elesto-dao/elesto/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         sdk.StoreKey
	paramSpace       paramtypes.Subspace
	bankKeeper       types.BankKeeper
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace, ak types.AccountKeeper, bk types.BankKeeper, feeCollectorName string,
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

// GetBlockInflation returns the last block time.
func (k Keeper) GetBlockInflation(ctx sdk.Context) (amount sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.CurrentInflationKey)
	if bz == nil {
		return
	}
	amount, ok := sdk.NewIntFromString(string(bz))
	if !ok {
		panic("cannot read the block inflation amount")
	}
	return
}

// SetBlockInflation stores the last block time.
func (k Keeper) SetBlockInflation(ctx sdk.Context, amount sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(amount.String())
	store.Set(types.CurrentInflationKey, bz)
}
