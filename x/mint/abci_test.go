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
	teamAddress := params.GetTeamAddress()

	// mint the initial
	err := s.keeper.MintCoins(s.ctx, sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, 200_000_000_000_000)))
	s.Assert().NoError(err)

	runDistributionTest := func(height int64, distribution types.InflationDistribution) {
		ctx = ctx.WithBlockHeight(height).WithBlockTime(ctx.BlockTime().Add(blockTime))
		oldDevTeamBalance := s.app.BankKeeper.GetBalance(ctx, sdk.AccAddress(params.GetTeamAddress()), params.GetMintDenom())
		oldCommunityPoolBalance := s.app.DistrKeeper.GetFeePoolCommunityCoins(ctx)
		oldFeeCollectorBalance := s.app.BankKeeper.GetBalance(ctx, feeCollector, params.GetMintDenom())

		mint.BeginBlocker(ctx, s.keeper)

		// assert dev team rewards
		newDevTeamBalance := s.app.BankKeeper.GetBalance(ctx, sdk.AccAddress(teamAddress), params.GetMintDenom())
		devTeamReward := newDevTeamBalance.Sub(oldDevTeamBalance)

		s.Assert().False(devTeamReward.IsNegative())
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
		s.Assert().EqualValues(distribution.BlockInflation, distribution.TeamRewards+communityReward.AmountOf(params.GetMintDenom()).TruncateInt64()+mintedAmt.Amount.Int64())
	}

	runEpochTest := func(height int64, epoch int64) {
		runDistributionTest(height-2, types.BlockInflationDistribution[epoch-1]) // 2 blocks to epoch
		runDistributionTest(height-1, types.BlockInflationDistribution[epoch-1]) // 1 block to epoch
		runDistributionTest(height, types.BlockInflationDistribution[epoch])     // new epoch
		runDistributionTest(height+1, types.BlockInflationDistribution[epoch])   // 1 block after epoch
		runDistributionTest(height+2, types.BlockInflationDistribution[epoch])   // 2 block after epoch
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

	// run tests for every epoch
	runEpochTest(6_307_200, 1)  // 1st epoch
	runEpochTest(12_614_400, 2) // 2nd epoch
	runEpochTest(18_921_600, 3) // 3rd epoch
	runEpochTest(25_228_800, 4) // 4th epoch
	runEpochTest(31_536_000, 5) // 5th epoch
	runEpochTest(37_843_200, 6) // 6th epoch
	runEpochTest(44_150_400, 7) // 7th epoch
	runEpochTest(50_457_600, 8) // 8th epoch
	runEpochTest(56_764_800, 9) // 9th epoch

	// 10th epoch and ensure no new mints after
	runDistributionTest(63_071_998, types.BlockInflationDistribution[9])
	runDistributionTest(63_071_999, types.BlockInflationDistribution[9])
	runDistributionTest(63_072_000, types.InflationDistribution{}) // new epoch
	runDistributionTest(63_072_001, types.InflationDistribution{})
	runDistributionTest(63_072_002, types.InflationDistribution{})

	// ensure no new mints after block 63_072_002
	for i := 0; i < 10; i++ {
		// get a random height after 63072000
		randomBlockOutsideEpoch := 63_072_000 + rand.Int63()
		runDistributionTest(randomBlockOutsideEpoch, types.InflationDistribution{})
	}
}
