package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/elesto-dao/elesto/x/did"
	"github.com/elesto-dao/elesto/x/did/keeper"
)

var (
	TypeMsgCreateDidDocument = sdk.MsgTypeURL(&did.MsgCreateDidDocument{})
)

const (
	opWeightMsgCreateDidDocument = "op_weight_msg_create_did_document"

	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDidDocument int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func GenerateGenesisState(simState *module.SimulationState) {
}

// ProposalContents doesn't return any content functions for governance proposals
func ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func WeightedOperations(simState module.SimulationState, didKeeper keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDidDocument, &weightMsgCreateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDidDocument = defaultWeightMsgCreateDidDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDidDocument,
		SimulateMsgCreateDidDocument(didKeeper, bk, ak),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
