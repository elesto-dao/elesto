package keeper

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocdc "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	server "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/elesto-dao/elesto/v2/x/credential"
	"github.com/elesto-dao/elesto/v2/x/did"
	didkeeper "github.com/elesto-dao/elesto/v2/x/did/keeper"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      Keeper
	queryClient credential.QueryClient
	keyring     keyring.Keyring
}

// SetupTest creates a test suite to test the did
func (suite *KeeperTestSuite) SetupTest() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("elesto", "elestopub")
	//config.Seal()

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
	ms.MountStoreWithDB(keyDidDocument, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyDidDocument, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true,
		server.ZeroLogWrapper{Logger: log.Logger},
	)

	interfaceRegistry := ct.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	cryptocdc.RegisterInterfaces(interfaceRegistry)

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

	// create the account for testing the credential signatures
	suite.keyring = keyring.NewInMemory()

	rc := func(index int, mnemonic string) {
		// register the test account
		i, err := suite.keyring.NewAccount(
			fmt.Sprint("test", index),
			mnemonic,
			keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1,
		)
		if err != nil {
			suite.Require().FailNow("cannot register test account")
		}
		a := accountKeeper.NewAccountWithAddress(ctx, i.GetAddress())
		a.SetPubKey(i.GetPubKey())
		accountKeeper.SetAccount(ctx, accountKeeper.NewAccount(ctx, a))
	}
	// register test accounts
	rc(0, "coil animal waste sound canvas weekend struggle skirt donor boil around bounce grant right silent year subway boost banana unlock powder riot spawn nerve")
	rc(1, "regret virtual damp hybrid armed powder motor open slim fall defy river goddess perfect invite orange assault reject involve quit salmon sunny abuse team")

	suite.ctx, suite.keeper, suite.queryClient = ctx, *credentialKeeper, queryClient
}

func (suite KeeperTestSuite) GetKeyAddress(uid string) sdk.Address {
	i, err := suite.keyring.Key(uid)
	if err != nil {
		suite.Require().FailNow("address for uid " + uid + " not found")
	}
	return i.GetAddress()
}

func (suite KeeperTestSuite) GetTestAccount() sdk.Address {
	return suite.GetTestAccountByIndex(0)
}

func (suite KeeperTestSuite) GetTestAccountByIndex(index int) sdk.Address {
	return suite.GetKeyAddress(fmt.Sprint("test", index))
}

func (suite KeeperTestSuite) GetRandomAccount() sdk.Address {
	uid := fmt.Sprint("a", rand.Int())
	i, _, err := suite.keyring.NewMnemonic(uid, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	if err != nil {
		suite.Require().FailNow("address for uid " + uid + " not found")
	}
	return i.GetAddress()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
