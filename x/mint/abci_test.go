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

	runDistributionTest := func(h int64, distribution types.InflationDistribution) {
		ctx = ctx.WithBlockHeight(h).WithBlockTime(ctx.BlockTime().Add(blockTime))
		//oldDevTeamBalance := s.app.BankKeeper.GetBalance(ctx, sdk.AccAddress(params.GetTeamAddress()), params.GetMintDenom())
		oldCommunityPoolBalance := s.app.DistrKeeper.GetFeePoolCommunityCoins(ctx)
		oldFeeCollectorBalance := s.app.BankKeeper.GetBalance(ctx, feeCollector, params.GetMintDenom())

		mint.BeginBlocker(ctx, s.keeper)

		// assert dev team rewards
		//addr := params.GetTeamAddress()
		//newDevTeamBalance := s.app.BankKeeper.GetBalance(ctx, sdk.AccAddress(addr), params.GetMintDenom())
		//devTeamReward := newDevTeamBalance.Sub(oldDevTeamBalance)

		//s.Assert().False(devTeamReward.IsNegative())
		//s.Assert().EqualValues(distribution.TeamRewards, devTeamReward.Amount.Int64(), "Dev Team rewards not matching expected values")

		// assert community pool
		newCommunityPoolBalance := s.app.DistrKeeper.GetFeePoolCommunityCoins(ctx)
		communityReward := newCommunityPoolBalance.Sub(oldCommunityPoolBalance)

		s.Assert().False(communityReward.IsAnyNegative())
		s.Assert().EqualValues(distribution.CommunityTax, communityReward.AmountOf(params.GetMintDenom()).TruncateInt64(), "CommunityTax not matching expected values")

		// assert staking rewards
		newFeeCollectorBalance := s.app.BankKeeper.GetBalance(ctx, feeCollector, params.GetMintDenom())
		mintedAmt := newFeeCollectorBalance.Sub(oldFeeCollectorBalance)
		s.Assert().False(mintedAmt.IsNegative())
		s.Assert().EqualValues(distribution.StakingRewards, mintedAmt.Amount.Int64(), "Staking rewards not matching expected values")

		// assert they add upto block inflation
	}

	// start from block 0

	// advance for the first few blocks
	runDistributionTest(0, types.InflationDistribution{})
	runDistributionTest(1, types.BlockInflationDistribution[0])
	runDistributionTest(2, types.BlockInflationDistribution[0])
	runDistributionTest(3, types.BlockInflationDistribution[0])
	runDistributionTest(4, types.BlockInflationDistribution[0])
	runDistributionTest(5, types.BlockInflationDistribution[0])

	// check the supply
	s.Assert().EqualValues(s.keeper.GetSupply(ctx, params.MintDenom).Amount.Int64(), 200_000_158_548_960)

	// go to the end of the first epoch
	runDistributionTest(6_307_198, types.BlockInflationDistribution[0])
	runDistributionTest(6_307_199, types.BlockInflationDistribution[0])
	runDistributionTest(6_307_200, types.BlockInflationDistribution[1]) // new epoch
	runDistributionTest(6_307_201, types.BlockInflationDistribution[1])
	runDistributionTest(6_307_202, types.BlockInflationDistribution[1])

	// 2nd epoch
	runDistributionTest(12_614_398, types.BlockInflationDistribution[1])
	runDistributionTest(12_614_399, types.BlockInflationDistribution[1])
	runDistributionTest(12_614_400, types.BlockInflationDistribution[2]) // new epoch
	runDistributionTest(12_614_401, types.BlockInflationDistribution[2])
	runDistributionTest(12_614_402, types.BlockInflationDistribution[2])

	// 3rd epoch
	runDistributionTest(18_921_598, types.BlockInflationDistribution[2])
	runDistributionTest(18_921_599, types.BlockInflationDistribution[2])
	runDistributionTest(18_921_600, types.BlockInflationDistribution[3]) // new epoch
	runDistributionTest(18_921_601, types.BlockInflationDistribution[3])
	runDistributionTest(18_921_602, types.BlockInflationDistribution[3])

	// 4th epoch
	runDistributionTest(25_228_798, types.BlockInflationDistribution[3])
	runDistributionTest(25_228_799, types.BlockInflationDistribution[3])
	runDistributionTest(25_228_800, types.BlockInflationDistribution[4]) // new epoch
	runDistributionTest(25_228_801, types.BlockInflationDistribution[4])
	runDistributionTest(25_228_802, types.BlockInflationDistribution[4])

	// 5th epoch
	runDistributionTest(31_535_998, types.BlockInflationDistribution[4])
	runDistributionTest(31_535_999, types.BlockInflationDistribution[4])
	runDistributionTest(31_536_000, types.BlockInflationDistribution[5]) // new epoch
	runDistributionTest(31_536_001, types.BlockInflationDistribution[5])
	runDistributionTest(31_536_002, types.BlockInflationDistribution[5])

	// 6th epoch
	runDistributionTest(37_843_198, types.BlockInflationDistribution[5])
	runDistributionTest(37_843_199, types.BlockInflationDistribution[5])
	runDistributionTest(37_843_200, types.BlockInflationDistribution[6]) // new epoch
	runDistributionTest(37_843_201, types.BlockInflationDistribution[6])
	runDistributionTest(37_843_202, types.BlockInflationDistribution[6])

	// 7th epoch
	runDistributionTest(44_150_398, types.BlockInflationDistribution[6])
	runDistributionTest(44_150_399, types.BlockInflationDistribution[6])
	runDistributionTest(44_150_400, types.BlockInflationDistribution[7]) // new epoch
	runDistributionTest(44_150_401, types.BlockInflationDistribution[7])
	runDistributionTest(44_150_402, types.BlockInflationDistribution[7])

	// 8th epoch
	runDistributionTest(50_457_598, types.BlockInflationDistribution[7])
	runDistributionTest(50_457_599, types.BlockInflationDistribution[7])
	runDistributionTest(50_457_600, types.BlockInflationDistribution[8]) // new epoch
	runDistributionTest(50_457_601, types.BlockInflationDistribution[8])
	runDistributionTest(50_457_602, types.BlockInflationDistribution[8])

	// 9th epoch
	runDistributionTest(56_764_798, types.BlockInflationDistribution[8])
	runDistributionTest(56_764_799, types.BlockInflationDistribution[8])
	runDistributionTest(56_764_800, types.BlockInflationDistribution[9]) // new epoch
	runDistributionTest(56_764_801, types.BlockInflationDistribution[9])
	runDistributionTest(56_764_802, types.BlockInflationDistribution[9])

	// 10th epoch
	runDistributionTest(63_071_998, types.BlockInflationDistribution[9])
	runDistributionTest(63_071_999, types.BlockInflationDistribution[9])
	runDistributionTest(63_072_000, types.InflationDistribution{}) // new epoch
	runDistributionTest(63_072_001, types.InflationDistribution{})
	runDistributionTest(63_072_002, types.InflationDistribution{})

	for i := 0; i < 10; i++ {
		// get a random height after 63072000
		randomBlockOutsideEpoch := 63_072_000 + rand.Int63()
		runDistributionTest(randomBlockOutsideEpoch, types.InflationDistribution{})
	}
}
