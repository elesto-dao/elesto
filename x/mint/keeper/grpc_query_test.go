package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	chain "github.com/elesto-dao/elesto/v4/app"
	"github.com/elesto-dao/elesto/v4/x/mint/types"
)

type MintTestSuite struct {
	suite.Suite

	app         *chain.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *MintTestSuite) SetupTest() {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.MintKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *MintTestSuite) TestGRPCParams() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient

	params, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, app.MintKeeper.GetParams(ctx))
}

func (suite *MintTestSuite) TestGRPCInflation() {
	_, _, queryClient := suite.app, suite.ctx, suite.queryClient

	inflation, err := queryClient.Inflation(context.Background(), &types.QueryInflationRequest{
		Height: 1,
	})
	suite.Require().NoError(err)
	suite.Require().EqualValues(1, inflation.Epoch)
	suite.Require().EqualValues("1.00", fmt.Sprintf("%.2f", inflation.Inflation))

	inflation, err = queryClient.Inflation(context.Background(), &types.QueryInflationRequest{
		Height: 6307200,
	})
	suite.Require().NoError(err)
	suite.Require().EqualValues(2, inflation.Epoch)
	suite.Require().EqualValues("0.50", fmt.Sprintf("%.2f", inflation.Inflation))

	inflation, err = queryClient.Inflation(context.Background(), &types.QueryInflationRequest{
		Height: 75686400,
	})
	suite.Require().NoError(err)
	suite.Require().EqualValues(13, inflation.Epoch)
	suite.Require().EqualValues("0.00", fmt.Sprintf("%.2f", inflation.Inflation))

}

func TestMintTestSuite(t *testing.T) {
	suite.Run(t, new(MintTestSuite))
}
