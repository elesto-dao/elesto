package module

import ( // this line is used by starport scaffolding # 1
	"context"
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/client/cli"
	"github.com/elesto-dao/elesto/v3/x/credential/keeper"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.Codec
}

func NewAppModuleBasic(cdc codec.Codec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return credential.ModuleName
}

//nolint
func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	credential.RegisterCodec(cdc)
}

//nolint
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	credential.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	credential.RegisterInterfaces(reg)
}

// DefaultGenesis returns the capability module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(credential.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	config client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(
	clientCtx client.Context,
	mux *runtime.ServeMux,
) {
	_ = credential.RegisterQueryHandlerClient(context.Background(), mux, credential.NewQueryClient(clientCtx))
	// this line is used by starport scaffolding # 2
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(credential.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	ak        credential.AccountKeeper
	bk        credential.BankKeeper
	didKeeper credential.DidKeeper
	keeper    keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak credential.AccountKeeper, bk credential.BankKeeper, didKeeper credential.DidKeeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		ak:             ak,
		bk:             bk,
		didKeeper:      didKeeper,
	}
}

// Name returns the capability module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the capability module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the capability module's query routing key.
func (AppModule) QuerierRoute() string { return credential.QuerierRoute }

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	credential.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	credential.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the capability module's invariants.
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the capability module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	gs json.RawMessage,
) []abci.ValidatorUpdate {
	var genState credential.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)
	keeper.InitGenesis(ctx, am.keeper, &genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(keeper.ExportGenesis(ctx, am.keeper))
}

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// ____________________________________________________________________________

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the credentials module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
}

// ProposalContents returns all the credentials content functions used to
// simulate governance proposals.
func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized credentials param changes for the simulator.
func (AppModule) RandomizedParams(*rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder registers a decoder for credentials module's types
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	//sdr[credentials.StoreKey] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	//return simulation.WeightedOperations(
	//	simState,
	//	am.keeper,
	//	am.bk,
	//	am.ak,
	//)
	return []simtypes.WeightedOperation{}
}

// ----------------------------------------------------------------------------
// Deprecation notice
// ----------------------------------------------------------------------------

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {

}

//nolint
// LegacyQuerierHandler returns the capability module's Querier.
func (AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}
