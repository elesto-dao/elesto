package mint_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/elesto-dao/elesto/v2/app"

	"github.com/elesto-dao/elesto/v2/x/mint"
	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
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

func (s *ModuleTestSuite) TestInflationRate() {

	// Here we define the expected supply amount for each epoch.
	// Amounts here are defined in uTokens

	expectedSupplies := map[int]int64{
		1:          200_000_000_000_000,
		6_307_200:  400_000_000_102_400,
		12_614_400: 600_000_000_102_400,
		18_921_600: 750_000_000_076_800,
		25_228_800: 843_750_000_048_000,
		31_536_000: 896_484_373_056_000,
		37_843_200: 924_499_513_051_200,
		44_150_400: 942_989_503_715_550,
		50_457_600: 961_849_291_393_125,
		56_764_800: 981_086_277_220_987,
		63_072_000: 1_000_000_000_139_230,
		69_379_200: 1_000_000_000_139_230,
		75_686_400: 1_000_000_000_139_230,
		81_993_600: 1_000_000_000_139_230,
	}

	// mint the initial supply
	initialSupply := 200_000_000_000_000

	params := s.app.MintKeeper.GetParams(s.ctx)

	// let's set the context at block height 1 to init the chain
	ctx := s.ctx.WithBlockHeight(int64(1))
	err := s.keeper.MintCoins(ctx, sdk.NewCoins(sdk.NewInt64Coin(params.MintDenom, int64(initialSupply))))
	s.Require().NoError(err)

	for block := 0; block <= 6_307_200; block++ {

		ctx := s.ctx.WithBlockHeight(int64(block))
		mint.BeginBlocker(ctx, s.keeper)

		if expectedSupply, ok := expectedSupplies[block]; ok {
			s.T().Log("reached block", block)
			s.Require().EqualValues(s.keeper.GetSupply(
				s.ctx.WithBlockHeight(int64(block)),
				params.MintDenom,
			).Amount.Int64(), expectedSupply)
		}

	}

}
