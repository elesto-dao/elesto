package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/elesto-dao/elesto/x/did"
	"github.com/elesto-dao/elesto/x/did/keeper"
)

var (
	TypeMsgCreateDidDocument = sdk.MsgTypeURL(&did.MsgCreateDidDocument{})
	TypeMsgAddVerification   = sdk.MsgTypeURL(&did.MsgAddVerification{})
)

const (
	opWeightMsgCreateDidDocument = "op_weight_msg_create_did_document"
	opWeightMsgAddVerification   = "op_weight_msg_create_add_verification"

	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDidDocument int = 100
	defaultWeightMsgAddVerification   int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// WeightedOperations returns the all the gov module operations with their respective weights.
func WeightedOperations(simState module.SimulationState, didKeeper keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDidDocument, &weightMsgCreateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDidDocument = defaultWeightMsgCreateDidDocument
		},
	)

	var weightMsgAddVerification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddVerification, &weightMsgCreateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgAddVerification = defaultWeightMsgAddVerification
		},
	)

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDidDocument,
		SimulateMsgCreateDidDocument(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddVerification,
		SimulateMsgAddVerification(didKeeper, bk, ak),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// SimulateMsgCreateDidDocument simulates a MsgCreateDidDocument message
func SimulateMsgCreateDidDocument(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		vmID := didID.NewVerificationMethodID(ownerAddress)
		vmType := did.EcdsaSecp256k1VerificationKey2019
		auth := did.NewVerification(
			did.NewVerificationMethod(
				vmID,
				didID,
				did.NewPublicKeyMultibase(didOwner.PubKey.Bytes()),
				vmType,
			),
			[]string{did.Authentication},
			nil,
		)

		_, found := k.GetDidDocument(ctx, []byte(didID))
		_, _ = did.NewDidDocument(didID.String(), did.WithVerifications(auth))
		msg := did.NewMsgCreateDidDocument(
			didID.String(),
			did.Verifications{auth},
			did.Services{},
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgCreateDidDocument,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		if found {
			// return an error that will not stop simulation as the did was found
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgCreateDidDocument, "did found, could not create did"), nil, nil
		}
		return opMsg, fOp, err
	}
}

// SimulateMsgAddVerification simulates a MsgCreateDidDocument message
func SimulateMsgAddVerification(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()

		vmKey, _ := simtypes.RandomAcc(r, accs)
		keyAddress := vmKey.Address.String()

		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)

		didvmkeyID := did.NewChainDID(ctx.ChainID(), keyAddress)
		vmkeyID := didvmkeyID.NewVerificationMethodID(keyAddress)
		vmkeyType := did.EcdsaSecp256k1VerificationKey2019

		auth := did.NewVerification(
			did.NewVerificationMethod(
				vmkeyID,
				didID,
				did.NewPublicKeyMultibase(vmKey.PubKey.Bytes()),
				vmkeyType,
			),
			[]string{did.Authentication},
			nil,
		)

		msg := did.NewMsgAddVerification(
			didID.String(),
			auth,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgAddVerification,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		_, found := k.GetDidDocument(ctx, []byte(didID))
		if !found {
			// return an error that will not stop simulation as the did was not found or the verification method already exists
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgAddVerification, "did not found, could not add verification method"), nil, nil
		}
		return opMsg, fOp, err

	}
}
