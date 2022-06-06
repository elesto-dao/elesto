package mint_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	chain "github.com/elesto-dao/elesto/app"

	"github.com/elesto-dao/elesto/x/mint"
	"github.com/elesto-dao/elesto/x/mint/keeper"
	"github.com/elesto-dao/elesto/x/mint/types"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000),
	)
)

type ModuleTestSuite struct {
	suite.Suite

	app    *chain.App
	ctx    sdk.Context
	keeper keeper.Keeper

	// TODO(gsora): figure out if we need addresses funded in some way.
	//addrs  []sdk.AccAddress
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

func (s *ModuleTestSuite) TestInitGenesis() {
	// default gent state case
	genState := types.DefaultGenesisState()
	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, genState)
	got := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(*genState, *got)
}

func (s *ModuleTestSuite) TestImportExportGenesis() {
	k, ctx := s.keeper, s.ctx
	genState := mint.ExportGenesis(ctx, k)
	bz := s.app.AppCodec().MustMarshalJSON(genState)

	var genState2 types.GenesisState
	s.app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	mint.InitGenesis(ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState2)

	genState3 := mint.ExportGenesis(ctx, k)
	s.Require().Equal(*genState, genState2)
	s.Require().Equal(genState2, *genState3)
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
		1:  {amount: 400000000, tolerance: defaultTolerance},
		2:  {amount: 600000000, tolerance: defaultTolerance},
		3:  {amount: 750000000, tolerance: defaultTolerance},
		4:  {amount: 843750000, tolerance: defaultTolerance},
		5:  {amount: 896484375, tolerance: defaultTolerance},
		6:  {amount: 924499512, tolerance: defaultTolerance},
		7:  {amount: 942989502, tolerance: defaultTolerance},
		8:  {amount: 961849292, tolerance: defaultTolerance},
		9:  {amount: 981086278, tolerance: defaultTolerance},
		10: {amount: 1000000000, tolerance: 0},
	}

	blocksPerYear := 6_307_200

	// We run the simulation for at least two times the amount of supply years to 
	// make sure past year 10, no more tokens are minted.
	simulationYears := len(expectedEstimatedSupply)*2
	
	initialSupply := 200_000_000_000_000

	ctx := s.ctx.WithBlockHeight(int64(0))

	err := s.keeper.MintCoins(ctx, sdk.NewCoins(sdk.NewInt64Coin("stake", int64(initialSupply))))
	s.Require().NoError(err)
	s.Require().EqualValues(s.keeper.GetSupply(
		s.ctx.WithBlockHeight(1),
		"stake",
	).Amount.Int64(), int64(initialSupply))

	s.T().Log("circulating supply at block 0:", s.keeper.GetSupply(ctx, "stake").String())

	for year := 1; year <= simulationYears; year++ {
		blockHeight := year * blocksPerYear

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
		if year > 10 {
			estimatedYear = 10 // past year 10, we expect always the same supply
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

func TestConstantInflation(t *testing.T) {
	//app := chain.Setup(false)
	//ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	//
	//app.InitChain(
	//	abcitypes.RequestInitChain{
	//		AppStateBytes: []byte("{}"),
	//		ChainId:       "test-chain-id",
	//	},
	//)
	//
	//blockTime := 5 * time.Second
	//
	//feeCollector := app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	//advanceHeight := func() sdk.Int {
	//	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(ctx.BlockTime().Add(blockTime))
	//	beforeBalance := app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
	//	mint.BeginBlocker(ctx, app.MintKeeper)
	//	afterBalance := app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
	//	mintedAmt := afterBalance.Sub(beforeBalance)
	//	require.False(t, mintedAmt.IsNegative())
	//	return mintedAmt.Amount
	//}
	//
	//ctx = ctx.WithBlockHeight(0).WithBlockTime(utils.ParseTime("2022-01-01T00:00:00Z"))
	//
	//// skip first block inflation, not set LastBlockTime
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(0))
	//
	//// after 2022-01-01 00:00:00
	//// 47564687 / 5 * (365 * 24 * 60 * 60) / 300000000000000 ~= 1
	//// 47564687 ~= 300000000000000 / (365 * 24 * 60 * 60) * 5
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	//
	//ctx = ctx.WithBlockHeight(100).WithBlockTime(utils.ParseTime("2023-01-01T00:00:00Z"))
	//
	//// applied 10sec(params.BlockTimeThreshold) block time due to block time diff is over params.BlockTimeThreshold
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	//
	//// 317097919 / 5 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	//// 317097919 ~= 200000000000000 / (365 * 24 * 60 * 60) * 5
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	//
	//blockTime = 10 * time.Second
	//// 634195839 / 10 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	//// 634195839 ~= 200000000000000 / (365 * 24 * 60 * 60) * 10
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	//
	//// over BlockTimeThreshold 10sec
	//blockTime = 20 * time.Second
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	//require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	//
	//// no inflation
	//ctx = ctx.WithBlockHeight(300).WithBlockTime(utils.ParseTime("2030-01-01T01:00:00Z"))
	//require.True(t, advanceHeight().IsZero())
	//require.True(t, advanceHeight().IsZero())
	//require.True(t, advanceHeight().IsZero())
	//require.True(t, advanceHeight().IsZero())
}

func (s *ModuleTestSuite) TestDefaultGenesis() {
	genState := *types.DefaultGenesisState()

	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState)
	got := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(genState, *got)
}

func (s *ModuleTestSuite) TestImportExportGenesisEmpty() {
	//emptyParams := types.DefaultParams()
	//emptyParams.InflationSchedules = []types.InflationSchedule{}
	//s.app.MintKeeper.SetParams(s.ctx, emptyParams)
	//genState := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	//
	//var genState2 types.GenesisState
	//bz := s.app.AppCodec().MustMarshalJSON(genState)
	//s.app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	//mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState2)
	//
	//genState3 := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	//s.Require().Equal(*genState, genState2)
	//s.Require().EqualValues(genState2, *genState3)
}
