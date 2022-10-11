package cli_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/noandrea/rl2020"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	var govGenState govtypes.GenesisState
	err := cfg.Codec.UnmarshalJSON(cfg.GenesisState[govtypes.ModuleName], &govGenState)
	s.Require().NoError(err)

	govGenState.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, sdk.NewInt(10)))
	govGenState.VotingParams.VotingPeriod = 2 * time.Second

	allowedIds := []string{"https://w3id.org/vc-revocation-list-2020/v1", "https://w3id.org/vc-revocation-list-2020/proposal", "https://w3id.org/vc-revocation-list-2020/bespoke", "https://example.id/credential/x"}
	// for governance tests allowed ids does not include v1/test
	sampleIds := append(allowedIds, "https://w3id.org/vc-revocation-list-2020/v1/test")

	cfg.GenesisState[credential.ModuleName] = cfg.Codec.MustMarshalJSON(sampleCredDefGenState(sampleIds, allowedIds))
	cfg.GenesisState[govtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&govGenState)

	cfg.AppConstructor = NewAppConstructor(app.MakeEncodingConfig(app.ModuleBasics))
	cfg.NumValidators = 2
	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func sampleCredDefGenState(ids []string, allowedIds []string) *credential.GenesisState {
	var genState credential.GenesisState
	for _, id := range ids {
		cred := credential.CredentialDefinition{
			Id:           id,
			PublisherId:  fmt.Sprintf("test-publisher-%v", id),
			Schema:       []byte(testSchema),
			Vocab:        []byte(testVocab),
			Name:         fmt.Sprintf("test-name-%v", id),
			Description:  fmt.Sprintf("test-desc-%v", id),
			SupersededBy: "",
			IsActive:     true,
		}
		genState.CredentialDefinitions = append(genState.CredentialDefinitions, cred)
	}
	genState.AllowedCredentialIds = allowedIds

	return &genState
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

func publishCredentialDefinition(s *IntegrationTestSuite, identifier, name, schemaFile, vocabFile string, val *network.Validator, code uint32) {

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
	s.Require().Equal(code, response.Code)
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

func createRevocationList(s *IntegrationTestSuite, rlID, issuerDidID, credDefID string, val *network.Validator, expectedCode uint32) {
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
	s.Require().Equal(response.Code, expectedCode)
}

func queryAllowedCredentials(s *IntegrationTestSuite, val *network.Validator) credential.QueryAllowedPublicCredentialsResponse {
	clientCtx := val.ClientCtx
	args := []string{
		fmt.Sprintf("--%s=1", flags.FlagPage),
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	cmd := cli.NewQueryAllowedCredentialDefinitionsCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	response := &credential.QueryAllowedPublicCredentialsResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
	return *response
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
				publishCredentialDefinition(s, identifier, label, schemaFile, vocabFile, val, 0)
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
		"credential-definition":          {},
		"credential-definitions":         {},
		"public-credential":              {},
		"public-credentials":             {},
		"credential-status":              {},
		"prepare-credential":             {},
		"public-credential-status":       {},
		"public-credentials-by-issuer":   {},
		"allowed-credential-definitions": {},
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
			"PASS: update creddef name, and active status",
			codes.OK,
			&sdk.TxResponse{},
			[]string{
				did.NewChainDID(s.cfg.ChainID, "validdid").String(),
				"ValidDefName2",
				"testdata/schema.json",
				"testdata/vocab.json",
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
	credDefID := "https://example.id/credential/x"
	testCredFile := fmt.Sprintf("%s/credential-%s.json", s.T().TempDir(), s.cfg.ChainID)

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
			val.ClientCtx,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
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

	rlID := "https://elesto.id/rl2020-test"
	createRevocationList(s, rlID, didID, credDefID, val, 0)

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

func (s *IntegrationTestSuite) TestNewSubmitProposePublicCredentialID() {
	val := s.network.Validators[0]
	didID := uuid.New().String()
	createDidDocument(s, didID, val)

	credDefID := "https://w3id.org/vc-revocation-list-2020/v1/test"
	rlID := "https://elesto.id/rl2020-ProposalPublicTest"

	// credential def not in allow list
	createRevocationList(s, rlID, didID, credDefID, val, 2104)

	testCases := []struct {
		name       string
		voteOption string
		respType   proto.Message
		args       []string
		ctx        client.Context
		expectErr  bool
		errorCode  uint32
	}{
		{
			"PASS: Propose cred def id to be a public credential",
			"yes",
			&sdk.TxResponse{},
			[]string{
				"submit-proposal",
				"propose-public-id",
				credDefID,
				fmt.Sprintf("--title=%s", "test proposal"),
				fmt.Sprintf("--description=%s", "test description"),
				fmt.Sprintf("--deposit=%s", sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val.ClientCtx,
			false,
			0,
		},
		{
			"INVALID: cred def id does not exist",
			"",
			&sdk.TxResponse{},
			[]string{
				"submit-proposal",
				"propose-public-id",
				"invalid-test-id",
				fmt.Sprintf("--title=%s", "test proposal"),
				fmt.Sprintf("--description=%s", "test description"),
				fmt.Sprintf("--deposit=%s", sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val.ClientCtx,
			true,
			5,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := govcli.NewTxCmd([]*cobra.Command{cli.NewSubmitProposePublicCredentialID()})

			proposalOut, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			s.Require().NoError(err)

			if tc.expectErr {
				s.Require().Contains(proposalOut.String(), "failed to execute")
				res := &sdk.TxResponse{}
				err := tc.ctx.Codec.UnmarshalJSON(proposalOut.Bytes(), res)
				s.Require().NoError(err)
				s.Require().Equal(res.Code, tc.errorCode)
			} else {
				s.Require().NotContains(proposalOut.String(), "failed to execute")

				// vote for the proposal
				out, err := clitestutil.ExecTestCLICmd(tc.ctx, govcli.NewCmdVote(), []string{
					"1",
					tc.voteOption,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				})
				s.Require().NotContains(out.String(), "failed to execute")
				s.Require().NoError(err)

				if tc.voteOption == "yes" {
					// query id from allowed list
					res := queryAllowedCredentials(s, val)
					found := false
					for _, def := range res.Credentials {
						if def.Id == credDefID {
							found = true
							break
						}
					}
					s.Require().True(found)

					// try to publish list according to allow list
					createRevocationList(s, rlID, didID, credDefID, val, 0)
				} else {
					createRevocationList(s, rlID, didID, credDefID, val, 2104)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewSubmitProposeRemovePublicCredentialID() {
	val := s.network.Validators[0]
	didID := uuid.New().String()
	createDidDocument(s, didID, val)
	credDefID := "https://w3id.org/vc-revocation-list-2020/proposal"
	rlID := "https://elesto.id/rl2020-ProposalRemoveFromPublicTest"

	// credential def allowed to be published
	createRevocationList(s, rlID, didID, credDefID, val, 0)

	testCases := []struct {
		name       string
		voteOption string
		respType   proto.Message
		args       []string
		ctx        client.Context
		expectErr  bool
		errorCode  uint32
	}{
		{
			"PASS: Propose cred def id to be removed from allowed list",
			"yes",
			&sdk.TxResponse{},
			[]string{
				"submit-proposal",
				"propose-remove-public-id",
				credDefID,
				fmt.Sprintf("--title=%s", "test proposal"),
				fmt.Sprintf("--description=%s", "test description"),
				fmt.Sprintf("--deposit=%s", sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val.ClientCtx,
			false,
			0,
		},
		{
			"INVALID: cred def id does not exist",
			"",
			&sdk.TxResponse{},
			[]string{
				"submit-proposal",
				"propose-remove-public-id",
				"invalid-test-id",
				fmt.Sprintf("--title=%s", "test proposal"),
				fmt.Sprintf("--description=%s", "test description"),
				fmt.Sprintf("--deposit=%s", sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			val.ClientCtx,
			true,
			5,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := govcli.NewTxCmd([]*cobra.Command{cli.NewSubmitRemoveProposePublicCredentialID()})

			proposalOut, err := clitestutil.ExecTestCLICmd(tc.ctx, cmd, tc.args)
			s.Require().NoError(err)

			if tc.expectErr {
				s.Require().Contains(proposalOut.String(), "failed to execute")
				res := &sdk.TxResponse{}
				err := tc.ctx.Codec.UnmarshalJSON(proposalOut.Bytes(), res)
				s.Require().NoError(err)
				s.Require().Equal(res.Code, tc.errorCode)
			} else {
				s.T().Log(proposalOut.String())
				s.Require().NotContains(proposalOut.String(), "failed to execute")

				// vote for the proposal
				out, err := clitestutil.ExecTestCLICmd(tc.ctx, govcli.NewCmdVote(), []string{
					"2",
					tc.voteOption,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				})
				s.Require().NotContains(out.String(), "failed to execute")
				s.Require().NoError(err)

				if tc.voteOption == "yes" {
					res := queryAllowedCredentials(s, val)
					found := false
					for _, def := range res.Credentials {
						if def.Id == credDefID {
							found = true
							break
						}
					}
					s.Require().False(found)

					// cannot publish the credential as removed from allowed
					createRevocationList(s, rlID+"-1", didID, credDefID, val, 2104)
				} else {
					// other id exists
					createRevocationList(s, rlID+"-1", didID, credDefID, val, 0)
				}
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
