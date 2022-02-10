package simulation

import (
	"fmt"
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
	TypeMsgCreateDidDocument  = sdk.MsgTypeURL(&did.MsgCreateDidDocument{})
	TypeMsgAddVerification    = sdk.MsgTypeURL(&did.MsgAddVerification{})
	TypeMsgRevokeVerification = sdk.MsgTypeURL(&did.MsgRevokeVerification{})
	TypeMsgAddService         = sdk.MsgTypeURL(&did.MsgAddService{})
	TypeMsgDeleteService      = sdk.MsgTypeURL(&did.MsgDeleteService{})
)

const (
	opWeightMsgCreateDidDocument  = "op_weight_msg_create_did_document"
	opWeightMsgAddVerification    = "op_weight_msg_create_add_verification"
	opWeightMsgRevokeVerification = "op_weight_msg_create_revoke_verification"
	opWeightMsgAddService         = "op_weight_msg_create_add_service"
	opWeightMsgDeleteService      = "op_weight_msg_create_delete_service"

	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDidDocument  int = 100
	defaultWeightMsgAddVerification    int = 100
	defaultWeightMsgRevokeVerification int = 100
	defaultWeightMsgAddService         int = 100
	defaultWeightMsgDeleteService      int = 100

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
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddVerification, &weightMsgAddVerification, nil,
		func(_ *rand.Rand) {
			weightMsgAddVerification = defaultWeightMsgAddVerification
		},
	)

	var weightMsgRevokeVerification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeVerification, &weightMsgRevokeVerification, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeVerification = defaultWeightMsgRevokeVerification
		},
	)

	var weightMsgAddService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddService, &weightMsgAddService, nil,
		func(_ *rand.Rand) {
			weightMsgAddService = defaultWeightMsgAddService
		},
	)

	var weightMsgDeleteService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteService, &weightMsgDeleteService, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteService = defaultWeightMsgDeleteService
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

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeVerification,
		SimulateMsgRevokeVerification(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddService,
		SimulateMsgAddService(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteService,
		SimulateMsgDeleteService(didKeeper, bk, ak),
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

// SimulateMsgAddVerification simulates a MsgAddVerification message
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

// SimulateMsgRevokeVerification simulates a MsgRevokeVerification message
func SimulateMsgRevokeVerification(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))

		// return an error that will not stop simulation as the did was not found or the verification method does not exists
		if !found {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgRevokeVerification, "did not found, could not remove verification method"), nil, nil
		}

		auth := didDoc.VerificationMethod

		// return an error that will not stop simulation if the length of the VM array is 1
		if len(auth) == 1 {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgRevokeVerification, "could not remove verification method as it would break the did document"), nil, nil
		}

		msg := did.NewMsgRevokeVerification(
			didID.String(),
			auth[1].Id,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgRevokeVerification,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		return opMsg, fOp, err
	}
}

// SimulateMsgAddService simulates a MsgAddService message
func SimulateMsgAddService(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))

		serviceID := "emti-agent" + fmt.Sprint(len(didDoc.Service))
		serviceType := "DIDComm"
		serviceURL := "https://agents.elesto.app.beta.starport.cloud/emti"

		newService := did.NewService(serviceID, serviceType, serviceURL)
		msg := did.NewMsgAddService(
			didID.String(),
			newService,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgAddService,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		// return an error that will not stop simulation as the did was not found
		if !found {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgAddService, "did not found, could not add service"), nil, nil
		}

		return opMsg, fOp, err
	}
}

// SimulateMsgDeleteService simulates a MsgDeleteService message
func SimulateMsgDeleteService(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))

		// return an error that will not stop simulation as the did was not found
		if !found {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgDeleteService, "did not found, could not remove service"), nil, nil
		}

		service := didDoc.Service

		// return an error that will not stop simulation if the length of the Service array is 0
		if len(service) == 0 {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgDeleteService, "could not remove verification method as it would break the did document"), nil, nil
		}

		serviceID := "emti-agent" + fmt.Sprint(len(didDoc.Service))

		msg := did.NewMsgDeleteService(
			didID.String(),
			serviceID,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgDeleteService,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}