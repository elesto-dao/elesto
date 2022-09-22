package mint_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/elesto-dao/elesto/v2/app"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

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
	ctx := app.BaseApp.NewContext(false, tmproto.Header{}).WithBlockTime(time.Now())

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.MintKeeper
}

func (s *ModuleTestSuite) TestInflationAmount() {

	params := s.app.MintKeeper.GetParams(s.ctx)

	ctx := s.ctx.WithBlockHeight(0).WithBlockTime(time.Date(2022, 06, 27, 21, 00, 00, 0, time.UTC))
	blockTime := 5 * time.Second
	feeCollector := s.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)

	// mint the initial
	err := s.keeper.MintCoins(s.ctx, sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, 200_000_000_000_000)))
	s.Assert().NoError(err)

	advanceHeight := func(h int64) int64 {
		ctx = ctx.WithBlockHeight(h).WithBlockTime(ctx.BlockTime().Add(blockTime))
		beforeBalance := s.app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
		mint.BeginBlocker(ctx, s.keeper)
		afterBalance := s.app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
		mintedAmt := afterBalance.Sub(beforeBalance)
		s.Assert().False(mintedAmt.IsNegative())
		return mintedAmt.Amount.Int64()
	}

	// start from block 0

	// advance for the first blocks
	s.Assert().EqualValues(advanceHeight(0), 0)
	s.Assert().EqualValues(advanceHeight(1), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(2), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(3), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(4), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(5), types.BlockInflationDistribution[0].BlockInflation)

	// check the supply
	s.Assert().EqualValues(s.keeper.GetSupply(ctx, params.MintDenom).Amount.Int64(), 200_000_158_548_960)

	// go to the end of the first epoch
	s.Assert().EqualValues(advanceHeight(6_307_198), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(6_307_199), types.BlockInflationDistribution[0].BlockInflation)
	s.Assert().EqualValues(advanceHeight(6_307_200), types.BlockInflationDistribution[1].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(6_307_201), types.BlockInflationDistribution[1].BlockInflation)
	s.Assert().EqualValues(advanceHeight(6_307_202), types.BlockInflationDistribution[1].BlockInflation)

	// 2nd epoch
	s.Assert().EqualValues(advanceHeight(12_614_398), types.BlockInflationDistribution[1].BlockInflation)
	s.Assert().EqualValues(advanceHeight(12_614_399), types.BlockInflationDistribution[1].BlockInflation)
	s.Assert().EqualValues(advanceHeight(12_614_400), types.BlockInflationDistribution[2].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(12_614_401), types.BlockInflationDistribution[2].BlockInflation)
	s.Assert().EqualValues(advanceHeight(12_614_402), types.BlockInflationDistribution[2].BlockInflation)

	// 3rd epoch
	s.Assert().EqualValues(advanceHeight(18_921_598), types.BlockInflationDistribution[2].BlockInflation)
	s.Assert().EqualValues(advanceHeight(18_921_599), types.BlockInflationDistribution[2].BlockInflation)
	s.Assert().EqualValues(advanceHeight(18_921_600), types.BlockInflationDistribution[3].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(18_921_601), types.BlockInflationDistribution[3].BlockInflation)
	s.Assert().EqualValues(advanceHeight(18_921_602), types.BlockInflationDistribution[3].BlockInflation)

	// 4th epoch
	s.Assert().EqualValues(advanceHeight(25_228_798), types.BlockInflationDistribution[3].BlockInflation)
	s.Assert().EqualValues(advanceHeight(25_228_799), types.BlockInflationDistribution[3].BlockInflation)
	s.Assert().EqualValues(advanceHeight(25_228_800), types.BlockInflationDistribution[4].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(25_228_801), types.BlockInflationDistribution[4].BlockInflation)
	s.Assert().EqualValues(advanceHeight(25_228_802), types.BlockInflationDistribution[4].BlockInflation)

	// 5th epoch
	s.Assert().EqualValues(advanceHeight(31_535_998), types.BlockInflationDistribution[4].BlockInflation)
	s.Assert().EqualValues(advanceHeight(31_535_999), types.BlockInflationDistribution[4].BlockInflation)
	s.Assert().EqualValues(advanceHeight(31_536_000), types.BlockInflationDistribution[5].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(31_536_001), types.BlockInflationDistribution[5].BlockInflation)
	s.Assert().EqualValues(advanceHeight(31_536_002), types.BlockInflationDistribution[5].BlockInflation)

	// 6th epoch
	s.Assert().EqualValues(advanceHeight(37_843_198), types.BlockInflationDistribution[5].BlockInflation)
	s.Assert().EqualValues(advanceHeight(37_843_199), types.BlockInflationDistribution[5].BlockInflation)
	s.Assert().EqualValues(advanceHeight(37_843_200), types.BlockInflationDistribution[6].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(37_843_201), types.BlockInflationDistribution[6].BlockInflation)
	s.Assert().EqualValues(advanceHeight(37_843_202), types.BlockInflationDistribution[6].BlockInflation)

	// 7th epoch
	s.Assert().EqualValues(advanceHeight(44_150_398), types.BlockInflationDistribution[6].BlockInflation)
	s.Assert().EqualValues(advanceHeight(44_150_399), types.BlockInflationDistribution[6].BlockInflation)
	s.Assert().EqualValues(advanceHeight(44_150_400), types.BlockInflationDistribution[7].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(44_150_401), types.BlockInflationDistribution[7].BlockInflation)
	s.Assert().EqualValues(advanceHeight(44_150_402), types.BlockInflationDistribution[7].BlockInflation)

	// 8th epoch
	s.Assert().EqualValues(advanceHeight(50_457_598), types.BlockInflationDistribution[7].BlockInflation)
	s.Assert().EqualValues(advanceHeight(50_457_599), types.BlockInflationDistribution[7].BlockInflation)
	s.Assert().EqualValues(advanceHeight(50_457_600), types.BlockInflationDistribution[8].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(50_457_601), types.BlockInflationDistribution[8].BlockInflation)
	s.Assert().EqualValues(advanceHeight(50_457_602), types.BlockInflationDistribution[8].BlockInflation)

	// 9th epoch
	s.Assert().EqualValues(advanceHeight(56_764_798), types.BlockInflationDistribution[8].BlockInflation)
	s.Assert().EqualValues(advanceHeight(56_764_799), types.BlockInflationDistribution[8].BlockInflation)
	s.Assert().EqualValues(advanceHeight(56_764_800), types.BlockInflationDistribution[9].BlockInflation) // new epoch
	s.Assert().EqualValues(advanceHeight(56_764_801), types.BlockInflationDistribution[9].BlockInflation)
	s.Assert().EqualValues(advanceHeight(56_764_802), types.BlockInflationDistribution[9].BlockInflation)

	// 10th epoch
	s.Assert().EqualValues(advanceHeight(63_071_998), types.BlockInflationDistribution[9].BlockInflation)
	s.Assert().EqualValues(advanceHeight(63_071_999), types.BlockInflationDistribution[9].BlockInflation)
	s.Assert().EqualValues(advanceHeight(63_072_000), 0) // new epoch
	s.Assert().EqualValues(advanceHeight(63_072_001), 0)
	s.Assert().EqualValues(advanceHeight(63_072_002), 0)

	for i := 0; i < 10; i++ {
		// get a random height after 63072000
		randomBlockOutsideEpoch := 63_072_000 + rand.Int63()
		s.Assert().EqualValues(advanceHeight(randomBlockOutsideEpoch), 0)
	}

}
