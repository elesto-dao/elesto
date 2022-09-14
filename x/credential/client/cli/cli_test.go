package cli_test

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/noandrea/rl2020"

	"github.com/cosmos/cosmos-sdk/client"

	"os"
	"runtime"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
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

	"github.com/elesto-dao/elesto/v3/app"
	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/client/cli"
	"github.com/elesto-dao/elesto/v3/x/did"
	didcli "github.com/elesto-dao/elesto/v3/x/did/client/cli"
)

var (
	//go:embed testdata/schema.json
	testSchema string
	//go:embed testdata/vocab.json
	testVocab string
)

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(encodingCfg app.EncodingConfig) network.AppConstructor {
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
	cfg.AppConstructor = NewAppConstructor(app.MakeEncodingConfig(app.ModuleBasics))
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

func publishCredentialDefinition(s *IntegrationTestSuite, identifier, name, schemaFile, vocabFile string, isPublic bool, val *network.Validator) {

	clientCtx := val.ClientCtx
	args := []string{
		identifier, name, schemaFile, vocabFile,
		fmt.Sprintf("--%s=%s", "public", strconv.FormatBool(isPublic)),
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

func createDidDocument(s *IntegrationTestSuite, identifier string, val *network.Validator) {
	clientCtx := val.ClientCtx
	args := []string{
		identifier,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := didcli.NewCreateDidDocumentCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	response := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
}

func createRevocationList(s *IntegrationTestSuite, rlID, issuerDidID, credDefID string, val *network.Validator) {
	clientCtx := val.ClientCtx
	args := []string{
		rlID,
		fmt.Sprintf("--issuer=%s", did.NewChainDID(s.cfg.ChainID, issuerDidID)),
		fmt.Sprintf("--definition-id=%s", credDefID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := cli.NewCreateRevocationListCmd()
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
		fixture   func()
		cdID      string
	}{
		{
			"FAIL: empty credential definition id",
			codes.InvalidArgument,
			&credential.QueryCredentialDefinitionResponse{},
			func() {},
			"",
		},
		{
			"PASS: found credential",
			codes.OK,
			&credential.QueryCredentialDefinitionResponse{},
			func() {
				var (
					identifier = "test-11234"
					label      = "test-11234"
					schemaFile = "testdata/schema.json"
					vocabFile  = "testdata/vocab.json"
				)
				publishCredentialDefinition(s, identifier, label, schemaFile, vocabFile, false, val)
			},
			"test-11234",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.fixture()
			cmd := cli.NewQueryCredentialDefinitionCmd()
			args := []string{
				tc.cdID,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				qr := tc.respType.(*credential.QueryCredentialDefinitionResponse)
				cd := qr.Definition
				s.Require().Equal(tc.cdID, cd.Id)
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

func (s *IntegrationTestSuite) TestNewPublishCredentialDefinitionCmd() {
	val1 := s.network.Validators[0]
	val2 := s.network.Validators[1]

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		args      []string
		ctx       client.Context
	}{
		{
			"PASS: valid data provided",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				did.NewChainDID(s.cfg.ChainID, "validdid").String(),
				"ValidDefName",
				"testdata/schema.json",
				"testdata/vocab.json",
				"--expiration=123",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val1.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val1.ClientCtx,
		},
		{
			"PASS: update creddef name, visibility and active status",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				did.NewChainDID(s.cfg.ChainID, "validdid").String(),
				"ValidDefName2",
				"testdata/schema.json",
				"testdata/vocab.json",
				"--public=true",
				"--inactive=true",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val1.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val1.ClientCtx,
		},
		{
			"FAIL: credDef with given ID is already published",
			codes.Unknown,
			&sdk.TxResponse{},
			[]string{
				did.NewChainDID(s.cfg.ChainID, "validdid").String(),
				"ValidDefName",
				"testdata/schema.json",
				"testdata/vocab.json",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val2.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val2.ClientCtx,
		},
		{
			"FAIL: bad file path",
			codes.Unknown,
			&sdk.TxResponse{},
			[]string{
				did.NewChainDID(s.cfg.ChainID, "validdid").String(),
				"ValidDefName",
				"testdata/bad/schema.json",
				"testdata/bad/vocab.json",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val1.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val1.ClientCtx,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewPublishCredentialDefinitionCmd()
			out, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				// check whether the credential def was published
				credDefID := tc.args[0] // first arg is the cred def id
				cmd = cli.NewQueryCredentialDefinitionCmd()
				queryArgs := []string{
					credDefID,
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(tc.ctx, cmd, queryArgs)

				s.Require().NoError(err)
				res := &credential.QueryCredentialDefinitionResponse{}
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), res))
				s.Require().JSONEq(testVocab, string(res.Definition.Vocab))
				s.Require().JSONEq(testSchema, string(res.Definition.Schema))
			}
		})
	}

}

func (s *IntegrationTestSuite) TestNewIssuePublicCredentialCmd() {
	val := s.network.Validators[0]
	credDefID := "http://example.id/credential/x"
	testCredFile := fmt.Sprintf("testdata/credential-%s.json", s.cfg.ChainID)

	// it is shared by all test cases
	publishCredentialDefinition(s,
		credDefID,
		"testDefName",
		"testdata/schema.json",
		"testdata/vocab.json",
		true,
		val,
	)
	issuerDidID := uuid.New().String()
	createDidDocument(s, issuerDidID, val)

	// copy the credential to a new file, with updated issuer DID with current chainID
	credentialBytes, err := os.ReadFile("testdata/credential.json")
	s.Require().NoError(err)
	var credentialJSON map[string]interface{}
	err = json.Unmarshal(credentialBytes, &credentialJSON)
	s.Require().NoError(err)
	credentialJSON["issuer"] = did.NewChainDID(s.cfg.ChainID, issuerDidID)

	updatedCredentialBytes, err := json.Marshal(credentialJSON)
	s.Require().NoError(err)

	err = os.WriteFile(testCredFile, updatedCredentialBytes, 0600)
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		args      []string
		credID    string
		ctx       client.Context
	}{
		{
			"PASS: valid credential is provided",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				credDefID,
				testCredFile,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			"https://example.com/credentials/status/3",
			val.ClientCtx,
		},
		{
			"FAIL: non-existent file",
			codes.Unknown,
			&sdk.TxResponse{},
			[]string{
				credDefID,
				"testdata/bad/credential.json",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			"https://example.com/credentials/status/3",
			val.ClientCtx,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewIssuePublicCredentialCmd()
			out, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				// check whether the credential was published
				cmd = cli.NewQueryPublicCredentialCmd()
				queryArgs := []string{
					tc.credID,
					fmt.Sprintf("--native=true"),
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(tc.ctx, cmd, queryArgs)

				s.Require().NoError(err)
				res := &credential.PublicVerifiableCredential{}
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), res))
			}
		})
	}

}

func (s *IntegrationTestSuite) TestNewCreateRevocationListCmd() {
	val := s.network.Validators[0]

	// shared by all test cases
	didID := uuid.New().String()
	createDidDocument(s, didID, val)

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		args      []string
		fixture   func()
		ctx       client.Context
	}{
		{
			"PASS: using default flags",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				"https://elesto.id/rl2020-1",
				fmt.Sprintf("--issuer=%s", did.NewChainDID(s.cfg.ChainID, didID)),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			func() {
				// default definition
				publishCredentialDefinition(s,
					"https://w3id.org/vc-revocation-list-2020/v1",
					"RevocationList2020",
					"testdata/schema.json",
					"testdata/vocab.json",
					true,
					val,
				)
			},
			val.ClientCtx,
		},
		{
			"PASS: custom flags provided",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				"https://elesto.id/rl2020-2",
				fmt.Sprintf("--issuer=%s", did.NewChainDID(s.cfg.ChainID, didID)),
				fmt.Sprintf("--size=%d", 32),
				fmt.Sprintf("--definition-id=%s", "https://w3id.org/vc-revocation-list-2020/bespoke"),
				fmt.Sprintf("--revoke=%s", "1"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			func() {
				// custom definition
				publishCredentialDefinition(s,
					"https://w3id.org/vc-revocation-list-2020/bespoke",
					"RevocationList2020CustomDef",
					"testdata/schema.json",
					"testdata/vocab.json",
					true,
					val,
				)
			},
			val.ClientCtx,
		},
		{
			"FAIL: duplicate credential",
			codes.Unknown,
			&sdk.TxResponse{},
			[]string{
				"https://elesto.id/rl2020-1",
				fmt.Sprintf("--issuer=%s", did.NewChainDID(s.cfg.ChainID, didID)),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			func() {
			},
			val.ClientCtx,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.fixture()
			cmd := cli.NewCreateRevocationListCmd()
			out, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				// check whether the credential was published
				credID := tc.args[0] // the first arg should be the cred ID
				cmd = cli.NewQueryPublicCredentialCmd()
				queryArgs := []string{
					credID,
					fmt.Sprintf("--native=true"),
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(tc.ctx, cmd, queryArgs)

				s.Require().NoError(err)
				res := &credential.PublicVerifiableCredential{}
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), res))
				s.Require().Len(res.Type, 2)
				s.Require().Equal(rl2020.TypeRevocationList2020Credential, res.Type[1])
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewUpdateRevocationListCmd() {
	val := s.network.Validators[0]

	// shared by all test cases
	didID := uuid.New().String()
	createDidDocument(s, didID, val)
	credDefID := "https://w3id.org/vc-revocation-list-2020/v1"
	publishCredentialDefinition(s,
		credDefID,
		"RevocationList2020Def",
		"testdata/schema.json",
		"testdata/vocab.json",
		true,
		val,
	)
	rlID := "https://elesto.id/rl2020-test"
	createRevocationList(s, rlID, didID, credDefID, val)

	testCases := []struct {
		name             string
		expectErr        codes.Code
		respType         proto.Message
		args             []string
		expectRevoked    []int
		expectNonRevoked []int
		ctx              client.Context
	}{
		{
			"PASS: revoke credential at index 1",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				rlID,
				fmt.Sprintf("--revoke=%d", 1),
				fmt.Sprintf("--definition-id=%s", credDefID),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			[]int{1},
			[]int{},
			val.ClientCtx,
		},
		{
			"PASS: reset credential at index 1",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				rlID,
				fmt.Sprintf("--reset=%d", 1),
				fmt.Sprintf("--definition-id=%s", credDefID),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			[]int{},
			[]int{1},
			val.ClientCtx,
		},
		{
			"FAIL: nonexistent credential",
			codes.Unknown,
			&sdk.TxResponse{},
			[]string{
				rlID + "non-existing",
				fmt.Sprintf("--reset=%d", 1),
				fmt.Sprintf("--definition-id=%s", credDefID),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			[]int{},
			[]int{1},
			val.ClientCtx,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewUpdateRevocationListCmd()
			out, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				// check whether the credential was published
				rlID := tc.args[0] // the first arg should be the cred ID
				cmd = cli.NewQueryPublicCredentialCmd()
				queryArgs := []string{
					rlID,
					fmt.Sprintf("--native=true"),
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(tc.ctx, cmd, queryArgs)
				s.Require().NoError(err)

				pubCred := &credential.PublicVerifiableCredential{}
				s.Require().NoError(tc.ctx.Codec.UnmarshalJSON(out.Bytes(), pubCred))

				revList, err := rl2020.NewRevocationListFromJSON(pubCred.CredentialSubject)
				s.Require().NoError(err)

				for _, idx := range tc.expectRevoked {
					isRevoked, err := revList.IsRevoked(rl2020.NewCredentialStatus(rlID, idx))
					s.Require().NoError(err)
					s.Require().True(isRevoked)
				}

				for _, idx := range tc.expectNonRevoked {
					isRevoked, err := revList.IsRevoked(rl2020.NewCredentialStatus(rlID, idx))
					s.Require().NoError(err)
					s.Require().False(isRevoked)
				}
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
