package simulation_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	elestoapp "github.com/elesto-dao/elesto/v3/app"
	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/simulation"
)

type SimTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *elestoapp.App
}

func (suite *SimTestSuite) SetupTest() {
	elestoApp := elestoapp.Setup(false)
	suite.app = elestoApp
	suite.ctx = elestoApp.BaseApp.NewContext(false, tmproto.Header{}).WithBlockTime(time.Now())

}

func (suite *SimTestSuite) TestWeightedOperations() {
	simState := module.SimulationState{}

	weightedOps := simulation.WeightedOperations(simState, suite.app.CredentialsKeeper, suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)

	// setup 3 accounts
	s := rand.NewSource(1)
	r := rand.New(s)
	accs := suite.getTestingAccounts(r, 3)

	expected := []struct {
		weight     int
		opMsgRoute string
		opMsgName  string
	}{
		{100, credential.ModuleName, simulation.TypeMsgPublishCredentialDefinition},
		{100, credential.ModuleName, simulation.TypeMsgPublishPublicVerifiableCredential},
		{100, credential.ModuleName, simulation.TypeMsgUpdateCredentialDefinition},
	}

	for i, w := range weightedOps {
		operationMsg, _, _ := w.Op()(r, suite.app.BaseApp, suite.ctx, accs, "")
		// the following checks are very much dependent from the ordering of the output given
		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
		// will fail
		suite.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
		suite.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
		suite.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
	}
}

func (suite *SimTestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
	accounts := simtypes.RandomAccounts(r, n)

	initAmt := suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 200000)
	initCoins := sdk.NewCoins(sdk.NewCoin("stake", initAmt))

	// add coins to the accounts
	for _, account := range accounts {
		acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, account.Address)
		suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
		suite.Require().NoError(simapp.FundAccount(suite.app.BankKeeper, suite.ctx, account.Address, initCoins))
	}

	return accounts
}

func (suite *SimTestSuite) TestMsgPublishCredentialDefinitionRequest() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 2)
	blockTime := time.Now().UTC()
	ctx := suite.ctx.WithBlockTime(blockTime)

	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// begin a new block
	suite.app.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmproto.Header{
				Height:  suite.app.LastBlockHeight() + 1,
				AppHash: suite.app.LastCommitID().Hash,
			},
		})

	signer := accounts[0]

	// execute operation
	op := simulation.SimulateMsgPublishCredentialDefinition(suite.app.CredentialsKeeper, suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
	suite.Require().NoError(err)

	var msg credential.MsgPublishCredentialDefinitionRequest
	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
	suite.Require().True(operationMsg.OK)
	suite.Require().Equal(signer.Address.String(), msg.Signer)
	suite.Require().Len(futureOperations, 0)
}

func (suite *SimTestSuite) TestSimulateMsgUpdateCredentialDefinition() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 2)
	blockTime := time.Now().UTC()
	ctx := suite.ctx.WithBlockTime(blockTime)

	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// begin a new block
	suite.app.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmproto.Header{
				Height:  suite.app.LastBlockHeight() + 1,
				AppHash: suite.app.LastCommitID().Hash,
			},
		})

	// execute operation
	op := simulation.SimulateMsgUpdateCredentialDefinition(suite.app.CredentialsKeeper, suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
	suite.Require().NoError(err)

	var msg credential.MsgUpdateCredentialDefinitionRequest
	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
	suite.Require().False(operationMsg.OK)
	suite.Require().Equal("", msg.Signer)
	suite.Require().Len(futureOperations, 0)
}

func (suite *SimTestSuite) TestSimulateMsgPublishPublicVerifiableCredential() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 2)
	blockTime := time.Now().UTC()
	ctx := suite.ctx.WithBlockTime(blockTime)

	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// begin a new block
	suite.app.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmproto.Header{
				Height:  suite.app.LastBlockHeight() + 1,
				AppHash: suite.app.LastCommitID().Hash,
			},
		})

	// execute operation
	op := simulation.SimulateMsgPublishPublicVerifiableCredential(suite.app.CredentialsKeeper, suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
	suite.Require().NoError(err)

	var msg credential.MsgIssuePublicVerifiableCredentialRequest
	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
	suite.Require().False(operationMsg.OK)
	suite.Require().Equal("", msg.Signer)
	suite.Require().Len(futureOperations, 0)
}

func TestSimTestSuite(t *testing.T) {
	suite.Run(t, new(SimTestSuite))
}
