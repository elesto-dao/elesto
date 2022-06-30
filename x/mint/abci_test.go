package mint_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	chain "github.com/elesto-dao/elesto/v2/app"

	"github.com/elesto-dao/elesto/v2/x/mint"
	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

type ModuleTestSuite struct {
	suite.Suite

	app    *chain.App
	ctx    sdk.Context
	keeper keeper.Keeper
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.MintKeeper
}

func (s *ModuleTestSuite) TestInflationRate() {
	// This test simulates data taken from here:
	// https://docs.google.com/spreadsheets/d/1sXwR-cYHS98in1aMzBabF7FNjqikbA3btnPMlFqgQ50/edit#gid=0

	type estimatedSupply struct {
		amount    int64
		tolerance int64
	}

	const defaultTolerance = int64(4) // https://xkcd.com/221/

	// Here we define the expected supply amount for each year.
	// Amounts here are defined in *tokens*, not microtokens.
	// We assume `defaultTolerance` amount of tokens as tolerance value between what's computed by the
	// chain and what's expected from the simulation Google Sheet linked at the beginning of this test.
	// Rationale on how the default tolerance has been determined is explained in the link above.
	// Year 10 is a special case, because we want to produce a maximum amount of 1 billion tokens by then.
	expectedEstimatedSupply := map[int]estimatedSupply{
		0: {amount: 400000000, tolerance: defaultTolerance},
		1: {amount: 600000000, tolerance: defaultTolerance},
		2: {amount: 750000000, tolerance: defaultTolerance},
		3: {amount: 843750000, tolerance: defaultTolerance},
		4: {amount: 896484375, tolerance: defaultTolerance},
		5: {amount: 924499512, tolerance: defaultTolerance},
		6: {amount: 942989502, tolerance: defaultTolerance},
		7: {amount: 961849292, tolerance: defaultTolerance},
		8: {amount: 981086278, tolerance: defaultTolerance},
		9: {amount: 1000000000, tolerance: 0},
	}

	blocksPerYear := 6_307_200

	// We run the simulation for at least two times the amount of supply years to
	// make sure past year 10, no more tokens are minted.
	simulationYears := len(expectedEstimatedSupply) * 2

	initialSupply := 200_000_000_000_000

	ctx := s.ctx.WithBlockHeight(int64(0))

	err := s.keeper.MintCoins(ctx, sdk.NewCoins(sdk.NewInt64Coin("stake", int64(initialSupply))))
	s.Require().NoError(err)
	s.Require().EqualValues(s.keeper.GetSupply(
		s.ctx.WithBlockHeight(1),
		"stake",
	).Amount.Int64(), int64(initialSupply))

	s.T().Log("circulating supply at block 0:", s.keeper.GetSupply(ctx, "stake").String())

	for year := 0; year <= simulationYears; year++ {
		// Adding 1 here because we're running the simulation on the first day of the following year.
		blockHeight := (year * blocksPerYear) + 1

		s.T().Log("simulating year", year, "block height", blockHeight)

		ctx := s.ctx.WithBlockHeight(int64(blockHeight))

		mint.BeginBlocker(ctx, s.keeper)

		// Since running this simulation for each block would make this test take too much time,
		// we mint the total amount of tokens minted in one year, minus 1 block since the `mint.BeginBlocker()`
		// call already mints once for us.
		blockInflationAmount := mint.BlockInflationAmount[year]
		mintAmount := sdk.NewInt(int64(blockInflationAmount) * int64(blocksPerYear-1))
		mintedCoin := sdk.NewCoin("stake", mintAmount)
		mintedCoins := sdk.NewCoins(mintedCoin)
		s.Require().NoError(s.keeper.MintCoins(ctx, mintedCoins))
		supply := s.keeper.GetSupply(ctx, "stake")
		s.T().Log("inflation for year", year, ":", "supply", supply)

		supplyInTokens := supply.Amount.Quo(sdk.NewInt(1000000)).ToDec().RoundInt64()

		estimatedYear := year
		if year > 9 {
			estimatedYear = 9 // past year 10, we expect always the same supply
		}

		yearExpectedSupply, found := expectedEstimatedSupply[estimatedYear]
		s.Require().True(found, "did not found expected supply for year %v", year)

		// Calculate the absolute difference between supply actually generated and supply expected
		// in the table above.
		difference := math.Abs(float64(supplyInTokens - yearExpectedSupply.amount))
		s.T().Log("difference:", difference, "tolerance:", yearExpectedSupply.tolerance, "expected:", yearExpectedSupply.amount, "got:", supplyInTokens)

		// Since we're dealing with absolute value, we can bypass checking negative amounts.
		if difference > float64(yearExpectedSupply.tolerance) {
			s.Require().Fail(
				"too big difference between expected and obtained supply",
				"difference between expected supply %d and obtained supply %d is not within acceptable range: %f, original supply %v",
				yearExpectedSupply.amount,
				supplyInTokens,
				difference,
				supply.String(),
			)
		}
	}

}

func (s *ModuleTestSuite) TestDefaultGenesis() {
	genState := *types.DefaultGenesisState()

	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState)
	got := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(genState, *got)
}
