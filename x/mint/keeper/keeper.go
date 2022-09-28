package keeper

import (
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
		accountKeeper:    ak,
		distrKeeper:      dk,
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

// AddInflationToFeeCollector implements an alias call to send the given amount to the fee collector
// AddInflationToFeeCollector is to be used in BeginBlocker to distribute inflation as staking rewards.
func (k Keeper) AddInflationToFeeCollector(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}

// AddInflationToCommunityTax sends the given amount of tokens to fund the community pool
// AddInflationToCommunityTax is to be used in BeginBlocker to distribute inflation as community tax
func (k Keeper) AddInflationToCommunityTax(ctx sdk.Context, amount sdk.Coins) error {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	return k.distrKeeper.FundCommunityPool(ctx, amount, moduleAddr)
}

// SendTeamRewards sends the given amount of tokens to the developer team address which is set in the params
// SendTeamRewards is to be used in BeginBlocker to distribute inflation to team
func (k Keeper) SendTeamRewards(ctx sdk.Context, amount sdk.Coins) error {
	teamAddrParam := k.GetParams(ctx).TeamAddress
	teamAddr, err := sdk.AccAddressFromBech32(teamAddrParam)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, teamAddr, amount)
}

// GetSupply returns the current supply on the chain for a denom
func (k Keeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	return k.bankKeeper.GetSupply(ctx, denom)
}
