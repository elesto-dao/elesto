package cli_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/credential/client/cli"
	"github.com/elesto-dao/elesto/x/did"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/elesto-dao/elesto/app"
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
	credential.RegisterInterfaces(cfg.InterfaceRegistry)
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

func publishCredentialDefinition(s *IntegrationTestSuite, identifier, name, schemaFile, vocabFile string, val *network.Validator) {

	clientCtx := val.ClientCtx
	args := []string{
		identifier, name, schemaFile, vocabFile,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := cli.NewPublishCredentialDefinitionCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	response := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
}

func (s *IntegrationTestSuite) TestGetCmdQueryCredentialDefinition() {

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
		cdDID     string
	}{
		{
			name() + "_1",
			codes.InvalidArgument,
			&credential.QueryCredentialDefinitionResponse{},
			func() {},
			"",
		},
		{
			name() + "_2",
			codes.OK,
			&credential.QueryCredentialDefinitionResponse{},
			func() {
				var (
					identifier = "test-11234"
					label      = "test-11234"
					schemaFile = "testdata/schema.json"
					vocabFile  = "testdata/vocab.json"
				)
				publishCredentialDefinition(s, identifier, label, schemaFile, vocabFile, val)
			},
			did.NewChainDID(clientCtx.ChainID, "test-11234").String(),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewQueryCredentialDefinitionCmd()
			args := []string{
				tc.cdDID,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				queryresponse := tc.respType.(*credential.QueryCredentialDefinitionResponse)
				cd := queryresponse.Definition
				s.Require().Equal(tc.cdDID, cd.Id)
			}
		})
	}
}

func TestGetTxCmd(t *testing.T) {
	//TODO: implement

	expectedCommands := map[string]struct{}{
		"issue-public-credential":       {},
		"publish-credential-definition": {},
		"create-revocation-list":        {},
		"update-revocation-list":        {},
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
		"credential-definition":        {},
		"credential-definitions":       {},
		"public-credential":            {},
		"public-credentials":           {},
		"credential-status":            {},
		"prepare-credential":           {},
		"public-credential-status":     {},
		"public-credentials-by-issuer": {},
	}

	t.Run("PASS: Verify command are there ", func(t *testing.T) {
		for _, x := range cli.GetQueryCmd("").Commands() {
			if _, ok := expectedCommands[x.Name()]; !ok {
				t.Errorf("GetTxCmd(): expected command not found %s", x.Name())
			}
		}
	})
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
