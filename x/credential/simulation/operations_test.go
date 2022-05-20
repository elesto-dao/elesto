package simulation_test
//
//import (
//	"encoding/json"
//	"math/rand"
//	"testing"
//	"time"
//
//	"github.com/cosmos/cosmos-sdk/simapp"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
//	"github.com/stretchr/testify/suite"
//	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
//	abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/libs/log"
//	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
//	dbm "github.com/tendermint/tm-db"
//
//	elestoapp "github.com/elesto-dao/elesto/app"
//	"github.com/elesto-dao/elesto/x/did"
//	"github.com/elesto-dao/elesto/x/did/simulation"
//)
//
//type SimTestSuite struct {
//	suite.Suite
//
//	ctx sdk.Context
//	app *elestoapp.App
//}
//
//func (suite *SimTestSuite) SetupTest() {
//	checkTx := false
//
//	db := dbm.NewMemDB()
//	app := elestoapp.New(
//		log.NewNopLogger(),
//		db,
//		nil,
//		true,
//		make(map[int64]bool),
//		elestoapp.DefaultNodeHome,
//		0,
//		cosmoscmd.MakeEncodingConfig(elestoapp.ModuleBasics),
//		simapp.EmptyAppOptions{},
//	)
//	suite.app = app.(*elestoapp.App)
//	cdc := suite.app.AppCodec()
//	genesisState := elestoapp.ModuleBasics.DefaultGenesis(cdc)
//
//	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
//	if err != nil {
//		panic(err)
//	}
//
//	// Initialize the chain
//	app.InitChain(
//		abci.RequestInitChain{
//			Validators:      []abci.ValidatorUpdate{},
//			ConsensusParams: simapp.DefaultConsensusParams,
//			AppStateBytes:   stateBytes,
//		},
//	)
//
//	suite.ctx = suite.app.GetBaseApp().NewContext(checkTx, tmproto.Header{})
//}
//
//func (suite *SimTestSuite) TestWeightedOperations() {
//	simState := module.SimulationState{}
//
//	weightedOps := simulation.WeightedOperations(simState, suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//
//	// setup 3 accounts
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accs := suite.getTestingAccounts(r, 3)
//
//	expected := []struct {
//		weight     int
//		opMsgRoute string
//		opMsgName  string
//	}{
//		{100, did.ModuleName, simulation.TypeMsgCreateDidDocument},
//		{100, did.ModuleName, simulation.TypeMsgAddVerification},
//		{100, did.ModuleName, simulation.TypeMsgRevokeVerification},
//		{200, did.ModuleName, simulation.TypeMsgSetVerificationRelationships},
//		{100, did.ModuleName, simulation.TypeMsgAddService},
//		{100, did.ModuleName, simulation.TypeMsgDeleteService},
//		{100, did.ModuleName, simulation.TypeMsgAddController},
//		{100, did.ModuleName, simulation.TypeMsgDeleteController},
//	}
//
//	for i, w := range weightedOps {
//		operationMsg, _, _ := w.Op()(r, suite.app.BaseApp, suite.ctx, accs, "")
//		// the following checks are very much dependent from the ordering of the output given
//		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
//		// will fail
//		suite.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
//		suite.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
//		suite.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
//	}
//}
//
//func (suite *SimTestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
//	accounts := simtypes.RandomAccounts(r, n)
//
//	initAmt := suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 200000)
//	initCoins := sdk.NewCoins(sdk.NewCoin("stake", initAmt))
//
//	// add coins to the accounts
//	for _, account := range accounts {
//		acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, account.Address)
//		suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
//		suite.Require().NoError(simapp.FundAccount(suite.app.BankKeeper, suite.ctx, account.Address, initCoins))
//	}
//
//	return accounts
//}
//
//func (suite *SimTestSuite) TestSimulateCreateDidDocument() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgCreateDidDocument(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgCreateDidDocument
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//	suite.Require().True(operationMsg.OK)
//	suite.Require().Equal(signer.Address.String(), msg.Signer)
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateAddVerification() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgAddVerification(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgAddVerification
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateRevokeVerification() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgRevokeVerification(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgRevokeVerification
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateSetVerificationRelationships() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgSetVerificationRelationships(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgSetVerificationRelationships
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateAddService() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgAddService(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgAddService
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateDeleteService() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgDeleteService(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgDeleteService
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateAddController() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgAddController(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgAddController
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func (suite *SimTestSuite) TestSimulateDeleteController() {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//	accounts := suite.getTestingAccounts(r, 2)
//	blockTime := time.Now().UTC()
//	ctx := suite.ctx.WithBlockTime(blockTime)
//
//	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})
//
//	// begin a new block
//	suite.app.BeginBlock(
//		abci.RequestBeginBlock{
//			Header: tmproto.Header{
//				Height:  suite.app.LastBlockHeight() + 1,
//				AppHash: suite.app.LastCommitID().Hash,
//			},
//		})
//
//	// TODO: create a DID for this account and add it to the store
//	// signer := accounts[0]
//
//	// execute operation
//	op := simulation.SimulateMsgDeleteController(suite.app.DidKeeper, suite.app.BankKeeper, suite.app.AccountKeeper)
//	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, ctx, accounts, "")
//	suite.Require().NoError(err)
//
//	var msg did.MsgAddController
//	suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
//
//	// TODO: check for success, needs a did in the store
//	// check the message was unsuccessful
//	suite.Require().False(operationMsg.OK)
//	suite.Require().Equal("", msg.Signer)
//
//	suite.Require().Len(futureOperations, 0)
//}
//
//func TestSimTestSuite(t *testing.T) {
//	suite.Run(t, new(SimTestSuite))
//}
