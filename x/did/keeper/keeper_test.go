package keeper

import (
	"fmt"
	"testing"

	"github.com/elesto-dao/elesto/x/did"
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
	queryClient did.QueryClient
}

// SetupTest creates a test suite to test the did
func (suite *KeeperTestSuite) SetupTest() {
	keyDidDocument := sdk.NewKVStoreKey(did.StoreKey)
	memKeyDidDocument := sdk.NewKVStoreKey(did.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDidDocument, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyDidDocument, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, server.ZeroLogWrapper{log.Logger})

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	k := NewKeeper(
		marshaler,
		keyDidDocument,
		memKeyDidDocument,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	did.RegisterQueryServer(queryHelper, k)
	queryClient := did.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.queryClient = ctx, *k, queryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestGenericKeeperSetAndGet() {
	testCases := []struct {
		msg     string
		didFn   func() did.DidDocument
		expPass bool
	}{
		{
			"data stored successfully",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument(
					"did:cosmos:net:elesto:subject",
				)
				return dd
			},
			true,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()
		suite.keeper.Set(suite.ctx,
			[]byte(dd.Id),
			[]byte{0x01},
			&dd,
			suite.keeper.cdc.MustMarshal,
		)
		suite.keeper.Set(suite.ctx,
			[]byte(dd.Id+"1"),
			[]byte{0x01},
			&dd,
			suite.keeper.cdc.MustMarshal,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.Get(
					suite.ctx,
					[]byte(dd.Id),
					[]byte{0x01},
					suite.keeper.UnmarshalDidDocument,
				)
				suite.Require().True(found)

				iterator := suite.keeper.GetAll(
					suite.ctx,
					[]byte{0x01},
				)
				defer iterator.Close()

				var array []interface{}
				for ; iterator.Valid(); iterator.Next() {
					array = append(array, iterator.Value())
				}
				suite.Require().Equal(2, len(array))
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGenericKeeperDelete() {
	testCases := []struct {
		msg     string
		didFn   func() did.DidDocument
		expPass bool
	}{
		{
			"data stored successfully",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument(
					"did:cosmos:net:elesto:subject",
				)
				return dd
			},
			true,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()
		suite.keeper.Set(suite.ctx,
			[]byte(dd.Id),
			[]byte{0x01},
			&dd,
			suite.keeper.cdc.MustMarshal,
		)
		suite.keeper.Set(suite.ctx,
			[]byte(dd.Id+"1"),
			[]byte{0x01},
			&dd,
			suite.keeper.cdc.MustMarshal,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				suite.keeper.Delete(
					suite.ctx,
					[]byte(dd.Id),
					[]byte{0x01},
				)

				_, found := suite.keeper.Get(
					suite.ctx,
					[]byte(dd.Id),
					[]byte{0x01},
					suite.keeper.UnmarshalDidDocument,
				)
				suite.Require().False(found)

			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
