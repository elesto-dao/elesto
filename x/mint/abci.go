package mint

import (
	"fmt"
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

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)

	if ctx.BlockHeight() == 1 || ctx.BlockHeight()%params.BlocksPerYear == 0 {
		// we change the inflation every year
		blocksPerYear := sdk.NewInt(params.BlocksPerYear)
		// calculate inflation
		inflationYear := int(math.Floor(float64(ctx.BlockHeight()) / float64(blocksPerYear.Int64())))
		// the inflation goes to 0 after we get over the last year
		if inflationYear > len(params.InflationRates) {
			// reset the mint amount
			k.SetBlockInflation(ctx, sdk.NewInt(0))
			return
		}
		// get the current year inflation rate
		inflationRate, err := sdk.NewDecFromStr(params.InflationRates[inflationYear])
		if err != nil {
			panic(fmt.Sprintf("begin block error: %v", err))
		}
		// get current supply
		circulatingSupply := k.GetSupply(ctx, params.MintDenom)
		// calculate total inflation amount per year
		yearInflationAmount := inflationRate.MulInt(circulatingSupply.Amount)

		// verify that does not overflow the max supply
		maxSupply := sdk.NewInt(params.MaxSupply)
		futureSupply := circulatingSupply.Amount.Add(yearInflationAmount.RoundInt())
		if futureSupply.GT(maxSupply) {
			// if it does overflow, adjust the total inflation amount so we converge to max supply
			yearInflationAmount = maxSupply.Sub(circulatingSupply.Amount).ToDec()
			// TODO check the following calculation
			// update the current inflation
			params.InflationRates[inflationYear] = yearInflationAmount.Quo(circulatingSupply.Amount.ToDec()).String()
			k.SetParams(ctx, params)
		}

		// mint team allocation
		teamAllocation, err := sdk.NewDecFromStr(params.TeamReward)
		if err != nil {
			ctx.Logger().Error("cannot retrieve team allocation reward", "error", err)
			panic(err)
		}
		teamAmount := yearInflationAmount.Mul(teamAllocation)
		// mint teamAmount and send to team address
		yearInflationAmount = yearInflationAmount.Sub(teamAmount)
		//
		mintedCoin := sdk.NewCoin(params.MintDenom, teamAmount.RoundInt())
		mintedCoins := sdk.NewCoins(mintedCoin)
		if err := k.MintCoins(ctx, mintedCoins); err != nil {
			panic(err)
		}
		// transfer the tokens to the team account
		if err := k.CollectAmount(ctx, params.TeamAddress, mintedCoins); err != nil {
			panic(err)
		}

		// calculate the amount to mint for each block
		// note: do not use floor or there is the risk of not reaching the max supply
		amountToMint := yearInflationAmount.Quo(blocksPerYear.ToDec()).RoundInt()
		// log the inflation change
		ctx.Logger().Info("updated inflation rate", "inflation", amountToMint.String())
		// save the block inflation amount to mint to the state
		k.SetBlockInflation(ctx, amountToMint)
	}
	blockInflationAmount := k.GetBlockInflation(ctx)
	mintedCoin := sdk.NewCoin(params.MintDenom, blockInflationAmount)
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
