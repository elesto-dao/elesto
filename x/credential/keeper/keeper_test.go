package keeper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	server "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
	didkeeper "github.com/elesto-dao/elesto/x/did/keeper"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      Keeper
	queryClient credential.QueryClient
}

// SetupTest creates a test suite to test the did
func (suite *KeeperTestSuite) SetupTest() {
	keyCreden := sdk.NewKVStoreKey(credential.StoreKey)
	memKeyCreden := sdk.NewKVStoreKey(credential.MemStoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramtypes.StoreKey)
	memKeyParams := sdk.NewKVStoreKey(paramtypes.TStoreKey)
	keyDidDocument := sdk.NewKVStoreKey(did.StoreKey)
	memKeyDidDocument := sdk.NewKVStoreKey(did.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyCreden, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyCreden, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, server.ZeroLogWrapper{log.Logger})

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
	}

	// init params keeper
	paramsKeeper := paramskeeper.NewKeeper(marshaler, nil, keyParams, memKeyParams)
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(did.ModuleName)
	paramsKeeper.Subspace(credential.ModuleName)

	authSubspace, ok := paramsKeeper.GetSubspace(authtypes.ModuleName)
	if !ok {
		suite.T().Log("subspace for auth module not found")
	}

	// init keepers
	accountKeeper := authkeeper.NewAccountKeeper(
		marshaler, keyAcc, authSubspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	didKeeper := didkeeper.NewKeeper(
		marshaler,
		keyDidDocument,
		memKeyDidDocument,
	)

	credentialKeeper := NewKeeper(
		marshaler,
		keyCreden,
		memKeyCreden,
		didKeeper,
		accountKeeper,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	credential.RegisterQueryServer(queryHelper, credentialKeeper)
	queryClient := credential.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.queryClient = ctx, *credentialKeeper, queryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
