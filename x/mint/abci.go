package mint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

// BlockMintAmounts is used to define the amounts that are involved in the mint schedule
type BlockMintAmounts struct {
	// This is the block inflation, how much tokens should be minted for a block
	Inflation int64
	// CommunityTax is the rewards amount to be sent to the community pool (per block)
	CommunityTax int64
	// TeamReward is the rewards amount to be sent to the team acconut (per block)
	TeamReward int64
}

const (
	// BlocksPerEpoch number of blocks in an Epoch, an epoch is
	// roughly an year assuming a block production rate of 1 block every 5s
	BlocksPerEpoch int64 = 6_307_200
)

var (
	// BlockInflationAmount are the amounts to be minted/distributed for each block
	// in an epoch.
	// The format of the BlockInflationAmount is <Epoch Number>: <BlockMintAmonuts>
	BlockInflationAmount = map[int64]BlockMintAmounts{
		0: {Inflation: 31_709_792, CommunityTax: 3_170_979, TeamReward: 3_170_979},
		1: {Inflation: 31_709_792, CommunityTax: 3_170_979, TeamReward: 3_170_979},
		2: {Inflation: 23_782_344, CommunityTax: 2_378_234, TeamReward: 2_378_234},
		3: {Inflation: 14_863_965, CommunityTax: 1_486_396, TeamReward: 1_486_396},
		4: {Inflation: 8_360_980, CommunityTax: 836_098, TeamReward: 836_098},
		5: {Inflation: 4_441_771, CommunityTax: 444_177, TeamReward: 444_177},
		6: {Inflation: 2_931_569, CommunityTax: 293_156, TeamReward: 293_156},
		7: {Inflation: 2_990_200, CommunityTax: 299_020, TeamReward: 299_020},
		8: {Inflation: 3_050_004, CommunityTax: 305_000, TeamReward: 305_000},
		9: {Inflation: 2_998_751, CommunityTax: 299_875, TeamReward: 299_875},
	}
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	if ctx.BlockHeight() <= 1 {
		// at block height 1 we have the initial supply, minting starts at block 2
		// TODO: start minting at block 2 probably breaks the expected supply by the end of the epoch
		return
	}
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// read the module parameters
	params := k.GetParams(ctx)

	// get the current inflation epoch
	inflationEpoch := ctx.BlockHeight() / BlocksPerEpoch

	// fetch the set of amounts from the inflationEpoch table
	// if there is no such epoch, then no mint is taking places
	amounts, exists := BlockInflationAmount[inflationEpoch]
	if !exists {
		return
	}

	// the coins to be minted are the block inflation - rewards
	coinsToMint := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, amounts.Inflation))
	if err := k.MintCoins(ctx, coinsToMint); err != nil {
		panic(err)
	}

	// send them from mint moduleAccount to the dev fund address
	teamRewardCoins := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, amounts.TeamReward))
	if err := k.CollectAmount(ctx, params.TeamAddress, teamRewardCoins); err != nil {
		panic(fmt.Errorf("cannot fund team account, %w", err))
	}

	// fund the community pool
	CommunityTaxCoins := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, amounts.CommunityTax))
	if err := k.FundCommunityPool(ctx, CommunityTaxCoins); err != nil {
		panic(fmt.Errorf("cannot fund community pool, %w", err))
	}

	// send the remaining minted coins to the fee collector account
	validatorFees := amounts.Inflation - (amounts.CommunityTax + amounts.TeamReward)
	validatorFeesCoins := sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, validatorFees))
	if err := k.AddInflationToFeeCollector(ctx, validatorFeesCoins); err != nil {
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
