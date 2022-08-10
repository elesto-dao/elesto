package cli_test

import (
	"crypto/rand"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/coinbase/kryptology/pkg/core/curves/native/bls12381"
	"github.com/gogo/protobuf/proto"
	"github.com/multiformats/go-multibase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elesto-dao/elesto/v2/x/did"
	"github.com/elesto-dao/elesto/v2/x/did/client/cli"

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

func name(others ...string) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return fmt.Sprintln(f.Name(), others)
}

func addNewDidDoc(s *IntegrationTestSuite, identifier string, val *network.Validator) {
	clientCtx := val.ClientCtx
	args := []string{
		identifier,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := cli.NewCreateDidDocumentCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}
	response := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
}

func (s *IntegrationTestSuite) TestGetCmdQueryDidDocument() {
	identifier := "123456789abcdefghijkb"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name() + "_1",
			codes.NotFound,
			&did.QueryDidDocumentResponse{},
			func() {},
		},
		{
			name() + "_2",
			codes.OK,
			&did.QueryDidDocumentResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIdentifer()
			identifiertoquery := "did:cosmos:" + clientCtx.ChainID + ":" + identifier
			args := []string{
				identifiertoquery,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				queryresponse := tc.respType.(*did.QueryDidDocumentResponse)
				diddoc := queryresponse.GetDidDocument()
				s.Require().Equal(identifiertoquery, diddoc.Id)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewCreateDidDocumentCmd() {
	identifier := "123456789abcdefghijkc"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
	}{
		{
			name(),
			[]string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				)},
			&sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {
			for i := 0; i < 3; i++ {
				cmd := cli.NewCreateDidDocumentCmd()
				tc.args[0] = identifier + fmt.Sprint(i)
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				//pull out the just created document
				cmd = cli.GetCmdQueryIdentifer()
				identifiertoquery := "did:cosmos:" + clientCtx.ChainID + ":" + tc.args[0]
				argstemp := []string{
					identifiertoquery,
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argstemp)
				s.Require().NoError(err)
				response1 := &did.QueryDidDocumentResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response1))
				s.Require().Equal(response1.GetDidDocument().Id, identifiertoquery)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddControllerCmd() {
	identifier1 := "123456789abcdefghijkd"
	identifier2 := "elesto1kslgpxklq75aj96cz3qwsczr95vdtrd3axw8ft"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			name(),
			[]string{
				identifier1,
				identifier2,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier1, val) },
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewAddControllerCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier1,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			controller := response.GetDidDocument().Controller
			s.Require().Equal(len(controller), 1)
			s.Require().Equal(controller[0], "did:cosmos:key:"+identifier2)
		})
	}
}

func (s *IntegrationTestSuite) TestNewDeleteControllerCmd() {
	identifier1 := "123456789abcdefghijkd"
	identifier2 := "elesto1kslgpxklq75aj96cz3qwsczr95vdtrd3axw8ft"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			name(),
			[]string{
				identifier1,
				identifier2,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			&sdk.TxResponse{},
			func() {
				// create a new did document
				addNewDidDoc(s, identifier1, val)
				// add a controller parameters
				args := []string{
					identifier1,
					identifier2,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				// add a controller paramter
				cmd := cli.NewAddControllerCmd()
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewDeleteControllerCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier1,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			controller := response.GetDidDocument().Controller
			s.Require().Equal(0, len(controller))
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddVerificationCmd() {
	identifier := "123456789abcdefghijke"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	generateG1 := func(compressed bool) string {
		var err error
		g1 := &bls12381.G1{}
		g1, err = g1.Random(rand.Reader)
		if err != nil {
			panic(err)
		}

		var ret []byte
		if compressed {
			rawRet := g1.ToCompressed()
			ret = rawRet[:]
		} else {
			rawRet := g1.ToUncompressed()
			ret = rawRet[:]
		}

		retStr, err := multibase.Encode(multibase.Base58BTC, ret)
		if err != nil {
			panic(err)
		}

		return retStr
	}

	generateG2 := func(compressed bool) string {
		var err error
		g2 := &bls12381.G2{}
		g2, err = g2.Random(rand.Reader)
		if err != nil {
			panic(err)
		}

		var ret []byte
		if compressed {
			rawRet := g2.ToCompressed()
			ret = rawRet[:]
		} else {
			rawRet := g2.ToUncompressed()
			ret = rawRet[:]
		}

		retStr, err := multibase.Encode(multibase.Base58BTC, ret)
		if err != nil {
			panic(err)
		}

		return retStr
	}

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		cliErr    require.ErrorAssertionFunc
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			[]string{
				identifier,
				`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AhJhB4NzRr2+pRpW4jDfajpML2h9yuBONsSqz6aXKZ6s"}`,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				"--verification-method-type=EcdsaSecp256k1VerificationKey2019",
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			require.NoError,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
		{
			name(),
			[]string{
				identifier,
				generateG1(true),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				"--verification-method-type=Bls12381G1Key2020",
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			require.NoError,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
		{
			name(),
			[]string{
				identifier,
				generateG1(false),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				"--verification-method-type=Bls12381G1Key2020",
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			require.NoError,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
		{
			name(),
			[]string{
				identifier,
				generateG2(true),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				"--verification-method-type=Bls12381G2Key2020",
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			require.NoError,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
		{
			name(),
			[]string{
				identifier,
				generateG2(false),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				"--verification-method-type=Bls12381G2Key2020",
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			require.NoError,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},

	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewAddVerificationCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			tc.cliErr(s.T(), err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
		})
	}

	//check for update
	cmd := cli.GetCmdQueryIdentifer()
	argsTemp := []string{
		"did:cosmos:" + clientCtx.ChainID + ":" + identifier,
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
	s.Require().NoError(err)
	response := &did.QueryDidDocumentResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
	authentications := response.GetDidDocument().Authentication
	verificationmethods := response.GetDidDocument().VerificationMethod
	expectedMethodsAmount := len(testCases) + 1 // there's a method in the did already
	s.Require().Equal(expectedMethodsAmount, len(authentications))
	s.Require().Equal(expectedMethodsAmount, len(verificationmethods))
	for i := 0; i < expectedMethodsAmount; i++ {
		s.Require().Equal(authentications[i], verificationmethods[i].Id)
	}

	// verify BLS keys
	for i := 2; i < expectedMethodsAmount; i++ {
		rawKey := testCases[i-1].args[1] // get the second argument from the i-nth test case

		_, rawBytes, err := multibase.Decode(rawKey)
		if err != nil {
			panic(err)
		}

		hexMb := did.NewPublicKeyMultibase(rawBytes).PublicKeyMultibase

		s.Equal(hexMb, verificationmethods[i].GetPublicKeyMultibase())
	}
}

func (s *IntegrationTestSuite) TestNewSetVerificationRelationshipsCmd() {
	identifier := "123456789abcdefghijkf"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			[]string{
				identifier,
				"",
				fmt.Sprintf("--relationship=%s", did.CapabilityDelegation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			verificationmethods := response.GetDidDocument().VerificationMethod
			s.Require().Greater(len(verificationmethods), 0)
			temp := strings.Split(verificationmethods[0].Id, "#")
			tc.args[1] = temp[len(temp)-1]
			cmd = cli.NewSetVerificationRelationshipCmd()

			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response = &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			capabilitydelegation := response.GetDidDocument().CapabilityDelegation
			s.Require().Equal(1, len(capabilitydelegation))
			s.Require().Equal(verificationmethods[0].Id, capabilitydelegation[0])
		})
	}
}

func (s *IntegrationTestSuite) TestNewRevokeVerificationCmd() {
	identifier := "123456789abcdefghijkg"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			[]string{
				identifier,
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))

			verificationmethods := response.GetDidDocument().VerificationMethod
			s.Require().Greater(len(verificationmethods), 0)
			temp := strings.Split(verificationmethods[0].Id, "#")
			tc.args[1] = temp[len(temp)-1]
			cmd = cli.NewRevokeVerificationCmd()

			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response = &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			s.Require().Equal(0, len(response.GetDidDocument().VerificationMethod))
			s.Require().Equal(0, len(response.GetDidDocument().Authentication))
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddServiceCmd() {
	identifier := "123456789abcdefghijkh"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			[]string{
				identifier,
				"service:seuro",
				"DIDComm",
				"service:euro/SIGNATURE",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			codes.OK,
			&sdk.TxResponse{},
			func() { addNewDidDoc(s, identifier, val) },
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewAddServiceCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			s.Require().Equal(1, len(response.GetDidDocument().Service))
			s.Require().Equal(tc.args[1], response.GetDidDocument().Service[0].Id)
			s.Require().Equal(tc.args[2], response.GetDidDocument().Service[0].Type)
			s.Require().Equal(tc.args[3], response.GetDidDocument().Service[0].ServiceEndpoint)
		})
	}
}

func (s *IntegrationTestSuite) TestNewDeleteServiceCmd() {
	identifier := "123456789abcdefghijki"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	args := []string{
		identifier,
		"service:seuro",
		"DIDComm",
		"service:euro/SIGNATURE",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	testCases := []struct {
		name     string
		respType proto.Message
		malleate func()
	}{
		{
			name(),
			&sdk.TxResponse{},
			func() {
				addNewDidDoc(s, identifier, val)
				cmd := cli.NewAddServiceCmd()
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {

			tc.malleate()
			cmd := cli.NewDeleteServiceCmd()

			args = append(args[:2], args[4:]...)

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			argsTemp := []string{
				"did:cosmos:" + clientCtx.ChainID + ":" + identifier,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, argsTemp)
			s.Require().NoError(err)
			response := &did.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			s.Require().Equal(0, len(response.GetDidDocument().Service))
		})
	}
}

func TestGetTxCmd(t *testing.T) {
	expectedCommands := map[string]struct{}{
		"create-did":                    {},
		"add-controller":                {},
		"delete-controller":             {},
		"add-service":                   {},
		"delete-service":                {},
		"add-verification-method":       {},
		"set-verification-relationship": {},
		"revoke-verification-method":    {},
		"link-aries-agent":              {},
	}

	t.Run("PASS: Verify command are there ", func(t *testing.T) {
		for _, x := range cli.GetTxCmd().Commands() {
			if _, ok := expectedCommands[x.Name()]; !ok {
				t.Errorf("GetTxCmd(): expected command not found %s", x.Name())
			}
		}
	})
}

func TestGetQueryCmd(t *testing.T) {
	expectedCommands := map[string]struct{}{
		"did": {},
	}

	t.Run("PASS: Verify command are there ", func(t *testing.T) {
		for _, x := range cli.GetQueryCmd("").Commands() {
			if _, ok := expectedCommands[x.Name()]; !ok {
				t.Errorf("GetQueryCmd(): expected command not found %s", x.Name())
			}
		}
	})
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
