package elesto

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/elesto-dao/elesto/v2/x/did"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"

	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/elesto-dao/elesto/v2/app"
	dbm "github.com/tendermint/tm-db"
)

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(encodingCfg cosmoscmd.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return app.New(
			val.Ctx.Logger,
			dbm.NewMemDB(), nil, true, make(map[int64]bool),
			val.Ctx.Config.RootDir,
			0,
			encodingCfg,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}

type IntegrationTestSuite struct {
	suite.Suite
	cfg     network.Config
	network *network.Network
}

// SetupSuite executes bootstrapping logic before all the tests, i.e. once before
// the entire suite, start executing.
func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	cfg := network.DefaultConfig()
	app.Setup(false)
	cfg.GenesisState = app.NewDefaultGenesisState(cfg.Codec)
	did.RegisterInterfaces(cfg.InterfaceRegistry)
	cfg.AppConstructor = NewAppConstructor(cosmoscmd.MakeEncodingConfig(app.ModuleBasics))
	cfg.NumValidators = 2
	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

// TearDownSuite performs cleanup logic after all the tests, i.e. once after the
// entire suite, has finished executing.
func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

// CreateGRPCConnection createa a grpc connection to a given url
func (s *IntegrationTestSuite) createGRPCConnection(addr string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(10)*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	s.NoError(err)

	return conn
}

func (s *IntegrationTestSuite) createDidDocViaValidator(identifier string) {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	pubKey, _ := clientCtx.Keyring.KeyByAddress(val.Address)
	address := pubKey.GetAddress().String()
	conn := s.createGRPCConnection(val.AppConfig.GRPC.Address)
	txClient := txtypes.NewServiceClient(conn)
	authQueryClient := authtypes.NewQueryClient(conn)
	accountRes, err := authQueryClient.Account(context.Background(), &authtypes.QueryAccountRequest{Address: address})
	s.NoError(err)

	baseAccount := &authtypes.BaseAccount{}
	err = baseAccount.Unmarshal(accountRes.Account.Value)
	s.NoError(err)

	vm := did.NewVerificationMethod(
		fmt.Sprintf("%s#%s", identifier, address),
		did.DID(identifier),
		did.NewPublicKeyMultibase(pubKey.GetPubKey().Bytes()),
		did.EcdsaSecp256k1VerificationKey2019,
	)
	ver := did.NewVerification(
		vm,
		[]string{
			did.Authentication,
			did.AssertionMethod,
			did.KeyAgreement,
			did.CapabilityInvocation,
			did.CapabilityDelegation,
		},
		nil,
	)
	service := did.NewService("service:agent1", "DIDComm", "https://elesto")
	msg := did.NewMsgCreateDidDocument(
		identifier,
		[]*did.Verification{ver},
		[]*did.Service{service},
		address,
	)
	txBuilder := s.cfg.TxConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(msg)
	s.NoError(err)

	txBuilder.SetGasLimit(200000000)
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin(s.network.Config.BondDenom, 1200)})
	// set the signatories of the tx
	sigV2 := signing.SignatureV2{
		PubKey: &secp256k1.PubKey{Key: val.PubKey.Bytes()}, // <------ passing an ed25519 pubkey, a secp256k1 pubkey is expected
		Data: &signing.SingleSignatureData{
			SignMode:  s.cfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: baseAccount.Sequence,
	}
	err = txBuilder.SetSignatures(sigV2)
	s.NoError(err)

	signerData := xauthsigning.SignerData{
		ChainID:       s.network.Config.ChainID,
		AccountNumber: baseAccount.AccountNumber,
		Sequence:      baseAccount.Sequence,
	}
	signBytes, err := s.cfg.TxConfig.SignModeHandler().GetSignBytes(
		s.cfg.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder.GetTx(),
	)
	s.NoError(err)
	signature, _, err := clientCtx.Keyring.SignByAddress(val.Address, signBytes)
	s.NoError(err)

	// Construct the SignatureV2 struct
	sigData := signing.SingleSignatureData{
		SignMode:  s.cfg.TxConfig.SignModeHandler().DefaultMode(),
		Signature: signature,
	}

	sigV2 = signing.SignatureV2{
		PubKey:   &secp256k1.PubKey{Key: val.PubKey.Bytes()}, // <------ passing an ed25519 pubkey, a secp256k1 pubkey is expected
		Data:     &sigData,
		Sequence: baseAccount.Sequence,
	}
	err = txBuilder.SetSignatures(sigV2)
	s.NoError(err)

	txBytes, err := s.cfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	s.NoError(err)
	resp, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	s.Require().NoError(err)
	s.Require().Equal(111222, int(resp.TxResponse.Code))
}

func (s *IntegrationTestSuite) TestCode111222() {
	identifier := "did:cosmos:" + s.network.Config.ChainID + ":bigbadwolf"
	s.createDidDocViaValidator(identifier)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
