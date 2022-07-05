package mint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

/**
initialize the network with a new account
that has 200_000_000 tokens as strategic reserve
TODO: testnet upgrade
*/

var (
	BlockInflationAmount = map[int]int64{
		0: 31709792,
		1: 31709792,
		2: 23782344,
		3: 14863965,
		4: 8360980,
		5: 4441771,
		6: 2931569,
		7: 2990200,
		8: 3050004,
		9: 2998751,
	}
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	var bootstrapDate time.Time

	switch ctx.BlockHeight() {
	case 0: // first block, write down the block timestamp
		if k.BootstrapDateCanarySet(ctx) {
			return // don't set the bootstrap date if canary was already set
		}

		if err := k.SetBootstrapDate(ctx, false); err != nil {
			panic(fmt.Errorf("cannot set bootstrap date at beginblock time, %w", err))
		}

		if err := k.SetBootstrapDateCanary(ctx, true, false); err != nil {
			panic(fmt.Errorf("cannot set bootstrap date canary at beginblock time, %w", err))
		}

		return
	default:
		bd, err := k.BootstrapDate(ctx)
		if err != nil {
			panic(fmt.Errorf("cannot fetch bootstrap date, %w", err))
		}

		bootstrapDate = bd
	}

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)

	// calculate inflation
	inflationYear := ctx.BlockTime().Year() - bootstrapDate.Year()

	inflationAmount, ok := BlockInflationAmount[inflationYear]
	if !ok {
		return
	}

	mintedCoin := sdk.NewCoin(params.MintDenom, sdk.NewInt(inflationAmount))
	mintedCoins := sdk.NewCoins(mintedCoin)
	if err := k.MintCoins(ctx, mintedCoins); err != nil {
		panic(err)
	}

	// calculate 10% off mintedCoins
	// this is handled in sdk.Ints already, some rounding error might persist.
	tenPercentRawAmt := mintedCoin.Amount.MulRaw(10).QuoRaw(100)
	tenPercentAmt := sdk.NewCoins(sdk.NewCoin(params.MintDenom, tenPercentRawAmt))

	// mintedCoins now has devFundAmt less coins, two times
	// one for the dev fund, one for the community pool
	mintedCoins = mintedCoins.Sub(tenPercentAmt).Sub(tenPercentAmt)

	// send them from mint moduleAccount to the dev fund address
	if err := k.CollectAmount(ctx, params.TeamAddress, tenPercentAmt); err != nil {
		panic(fmt.Errorf("cannot send coins to team account, %w", err))
	}

	// fund the community pool
	if err := k.FundCommunityPool(ctx, tenPercentAmt); err != nil {
		panic(fmt.Errorf("cannot fund community pool, %w", err))
	}

	// send the remaining minted coins to the fee collector account
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
