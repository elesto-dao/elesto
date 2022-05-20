package keeper

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	server "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tm-db"
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

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyCreden, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyCreden, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, server.ZeroLogWrapper{log.Logger})

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	k := NewKeeper(
		marshaler,
		keyCreden,
		memKeyCreden,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	credential.RegisterQueryServer(queryHelper, k)
	queryClient := credential.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.queryClient = ctx, *k, queryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
