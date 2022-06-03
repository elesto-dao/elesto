package mint

import (
	"math"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/mint/keeper"
	"github.com/elesto-dao/elesto/x/mint/types"
)

/**
initialize the network with a new account
that has 200_000_000 tokens as strategic reserve
TODO: testnet upgrade
*/

var (
	BlockInflationAmount = map[int]int64{
		1:  31709792,
		2:  31709792,
		3:  23782344,
		4:  14863965,
		5:  8360980,
		6:  4441771,
		7:  2931569,
		8:  2990200,
		9:  3050004,
		10: 2998751,
	}

	blocksPerYear = int64(6_307_200)
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	if ctx.BlockHeight() == 0 {
		return
	}

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)

	// calculate inflation
	inflationYear := int(math.Floor(float64(ctx.BlockHeight()) / float64(blocksPerYear)))

	inflationAmount, ok := BlockInflationAmount[inflationYear]
	if !ok {
		return
	}

	mintedCoin := sdk.NewCoin(params.MintDenom, sdk.NewInt(inflationAmount))
	mintedCoins := sdk.NewCoins(mintedCoin)
	if err := k.MintCoins(ctx, mintedCoins); err != nil {
		panic(err)
	}

	// TODO: 80/10/10 splitted between
	// - 80: default fee collector account
	// - 10: team account
	// - 10: community pool
	// send the minted coins to the fee collector account
	if err := k.AddInflationToFeeCollector(ctx, mintedCoins); err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}
	if err := ctx.EventManager().EmitTypedEvent(&types.MintEvent{
		Amount: mintedCoin.Amount.String(),
	}); err != nil {
		panic(err)
	}
}
