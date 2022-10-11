package mint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/mint/keeper"
	"github.com/elesto-dao/elesto/v4/x/mint/types"
)

const (
	// BlocksPerEpoch number of blocks in an Epoch, an epoch is
	// roughly an year assuming a block production rate of 1 block every 5s
	BlocksPerEpoch int64 = 6_307_200
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	if ctx.BlockHeight() < 1 {
		// the inflation is already calculated for block 1. The inflation of the block 1 esist in the time
		// between block 0 (chain start) and block 1.
		return
	}
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// read the module parameters
	params := k.GetParams(ctx)

	// get the current inflation epoch
	inflationEpoch := ctx.BlockHeight() / BlocksPerEpoch

	// fetch inflation from the inflationEpoch table
	// if there is no such epoch, then no mint is taking places
	inflationDistribution, exists := types.BlockInflationDistribution[inflationEpoch]
	if !exists {
		return
	}

	// the coins to be minted are the block inflation - rewards
	coinsToMint := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, inflationDistribution.BlockInflation))
	if err := k.MintCoins(ctx, coinsToMint); err != nil {
		panic(err)
	}

	distributeInflation(ctx, k, inflationDistribution, params)

	// telemetry
	if coinsToMint.AmountOf(params.MintDenom).IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(coinsToMint.AmountOf(params.MintDenom).Int64()), "minted_tokens")
	}
	// fire the event
	if err := ctx.EventManager().EmitTypedEvent(&types.MintEvent{
		Amount: coinsToMint.AmountOf(params.MintDenom).String(),
	}); err != nil {
		panic(err)
	}
}

func distributeInflation(ctx sdk.Context, k keeper.Keeper, inflationDistribution types.InflationDistribution, params types.Params) {
	// distribute team rewards
	teamRewards := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, inflationDistribution.TeamRewards))
	if err := k.SendTeamRewards(ctx, teamRewards); err != nil {
		panic(fmt.Errorf("cannot distribute block inflation to dev team address: %w", err))
	}

	// distribute community tax
	communityTax := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, inflationDistribution.CommunityTax))
	if err := k.AddInflationToCommunityTax(ctx, communityTax); err != nil {
		panic(fmt.Errorf("cannot distribute block inflation to community pool: %w", err))
	}

	// distribute staking rewards - send the minted coins to the fee collector account
	stakingRewards := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, inflationDistribution.StakingRewards))
	if err := k.AddInflationToFeeCollector(ctx, stakingRewards); err != nil {
		panic(fmt.Errorf("cannot distribute block inflation to fee collector account: %w", err))
	}
}
