package mint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

const (
	// BlocksPerEpoch number of blocks in an Epoch, an epoch is
	// roughly an year assuming a block production rate of 1 block every 5s
	BlocksPerEpoch int64 = 6_307_200
)

var (
	// BlockInflationDistribution are the amounts to be minted and distributed for each block
	// in an epoch.
	// The format of the BlockInflationDistribution is <Epoch Number>: <BlockInflation, TeamRewards, CommunityTax>
	BlockInflationDistribution = map[int64]types.InflationDistribution{
		0: {31_709_792, 1, 1},
		1: {31_709_792, 1, 1},
		2: {23_782_344, 1, 1},
		3: {14_863_965, 1, 1},
		4: {8_360_980, 1, 1},
		5: {4_441_771, 1, 1},
		6: {2_931_569, 1, 1},
		7: {2_990_200, 1, 1},
		8: {3_050_004, 1, 1},
		9: {2_998_751, 1, 1},
	}
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
	inflationDistribution, exists := BlockInflationDistribution[inflationEpoch]
	if !exists {
		return
	}

	// the coins to be minted are the block inflation - rewards
	coinsToMint := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, inflationDistribution.BlockInflation))
	if err := k.MintCoins(ctx, coinsToMint); err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	if err := k.AddInflationToFeeCollector(ctx, coinsToMint); err != nil {
		panic(fmt.Errorf("cannot distribute block inflation: %w", err))
	}

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
