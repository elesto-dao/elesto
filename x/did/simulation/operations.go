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
	TypeMsgCreateDidDocument            = sdk.MsgTypeURL(&did.MsgCreateDidDocument{})
	TypeMsgUpdateDidDocument            = sdk.MsgTypeURL(&did.MsgUpdateDidDocument{})
	TypeMsgAddVerification              = sdk.MsgTypeURL(&did.MsgAddVerification{})
	TypeMsgRevokeVerification           = sdk.MsgTypeURL(&did.MsgRevokeVerification{})
	TypeMsgSetVerificationRelationships = sdk.MsgTypeURL(&did.MsgSetVerificationRelationships{})
	TypeMsgAddService                   = sdk.MsgTypeURL(&did.MsgAddService{})
	TypeMsgDeleteService                = sdk.MsgTypeURL(&did.MsgDeleteService{})
	TypeMsgAddController                = sdk.MsgTypeURL(&did.MsgAddController{})
	TypeMsgDeleteController             = sdk.MsgTypeURL(&did.MsgDeleteController{})
)

const (
	opWeightMsgCreateDidDocument            = "op_weight_msg_create_did_document"
	opWeightMsgUpdateDidDocument            = "op_weight_msg_update_did_document"
	opWeightMsgAddVerification              = "op_weight_msg_create_add_verification"
	opWeightMsgRevokeVerification           = "op_weight_msg_create_revoke_verification"
	opWeightMsgSetVerificationRelationships = "op_weight_msg_create_set_verification_relationships"
	opWeightMsgAddService                   = "op_weight_msg_create_add_service"
	opWeightMsgDeleteService                = "op_weight_msg_create_delete_service"
	opWeightMsgAddController                = "op_weight_msg_create_add_controller"
	opWeightMsgDeleteController             = "op_weight_msg_create_delete_controller"

	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDidDocument            int = 100
	defaultWeightMsgUpdateDidDocument            int = 100
	defaultWeightMsgAddVerification              int = 100
	defaultWeightMsgRevokeVerification           int = 100
	defaultWeightMsgSetVerificationRelationships int = 200
	defaultWeightMsgAddService                   int = 100
	defaultWeightMsgDeleteService                int = 100
	defaultWeightMsgAddController                int = 100
	defaultWeightMsgDeleteController             int = 100

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

	var weightMsgUpdateDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDidDocument, &weightMsgUpdateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDidDocument = defaultWeightMsgUpdateDidDocument
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

	var weightMsgSetVerificationRelationships int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetVerificationRelationships, &weightMsgSetVerificationRelationships, nil,
		func(_ *rand.Rand) {
			weightMsgSetVerificationRelationships = defaultWeightMsgSetVerificationRelationships
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

	var weightMsgAddController int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddController, &weightMsgAddController, nil,
		func(_ *rand.Rand) {
			weightMsgAddController = defaultWeightMsgAddController
		},
	)

	var weightMsgDeleteController int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteController, &weightMsgDeleteController, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteController = defaultWeightMsgDeleteController
		},
	)

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDidDocument,
		SimulateMsgCreateDidDocument(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDidDocument,
		SimulateMsgUpdateDidDocument(didKeeper, bk, ak),
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
		weightMsgSetVerificationRelationships,
		SimulateMsgSetVerificationRelationships(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddService,
		SimulateMsgAddService(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteService,
		SimulateMsgDeleteService(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddController,
		SimulateMsgAddController(didKeeper, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteController,
		SimulateMsgDeleteController(didKeeper, bk, ak),
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
		vm := did.NewVerification(
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
		_, _ = did.NewDidDocument(didID.String(), did.WithVerifications(vm))
		msg := did.NewMsgCreateDidDocument(
			didID.String(),
			did.Verifications{vm},
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
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgCreateDidDocument,
				"did found, could not create did",
			), nil, nil
		}
		return opMsg, fOp, err
	}
}

// SimulateMsgUpdateDidDocument simulates a MsgUpdateDidDocument message
func SimulateMsgUpdateDidDocument(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)

		// build vm 1
		vmKey, _ := simtypes.RandomAcc(r, accs)
		keyAddress := vmKey.Address.String()
		didvmkeyID := did.NewChainDID(ctx.ChainID(), keyAddress)
		vmkeyID := didvmkeyID.NewVerificationMethodID(keyAddress)
		vmkeyType := did.EcdsaSecp256k1VerificationKey2019
		vm := did.NewVerification(
			did.NewVerificationMethod(
				vmkeyID,
				didID,
				did.NewPublicKeyMultibase(vmKey.PubKey.Bytes()),
				vmkeyType,
			),
			[]string{did.Authentication},
			nil,
		)

		// build vm 2
		vmKey2, _ := simtypes.RandomAcc(r, accs)
		accAddress2 := sdk.AccAddress(vmKey2.PubKey.Address())
		vmID2 := didID.NewVerificationMethodID(accAddress2.String())
		vmType2 := did.CosmosAccountAddress
		vm2 := did.NewVerification(
			did.NewVerificationMethod(
				vmID2,
				didID,
				did.NewBlockchainAccountID("elesto", accAddress2.String()),
				vmType2,
			),
			[]string{did.KeyAgreement, did.AssertionMethod, did.Authentication},
			nil,
		)

		// create services to add to did doc
		service1 := did.NewService(fmt.Sprint("service:emti-agent1", r.Int()), "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
		service2 := did.NewService(fmt.Sprint("service:emti-agent", r.Int()), "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
		service3 := did.NewService(fmt.Sprint("service:emti-agent", r.Int()), "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")

		// create controllers to add to did doc
		controller, _ := simtypes.RandomAcc(r, accs)
		controllerAddress := sdk.AccAddress(controller.PubKey.Address())
		controllerDidID := did.NewKeyDID(controllerAddress.String())

		// update the did document
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))
		err := didDoc.AddControllers(controllerDidID.String())
		if err != nil {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgUpdateDidDocument,
				"did not found, could not add controller",
			), nil, err
		}

		err = didDoc.AddVerifications(vm, vm2)
		if err != nil {
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgUpdateDidDocument,
				"did not found, could not add vm",
			), nil, err
		}
		err = didDoc.AddServices(service1, service2, service3)
		if err != nil {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgUpdateDidDocument,
				"did not found, could not add service",
			), nil, err
		}

		msg := did.NewMsgUpdateDidDocument(
			&didDoc,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgUpdateDidDocument,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		if !found {
			// return an error that will not stop simulation as the did was not found
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgUpdateDidDocument,
				"did not found, could not add verification method",
			), nil, nil
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

		didDoc, found := k.GetDidDocument(ctx, []byte(didID))
		for _, vm := range didDoc.VerificationMethod {
			if vmkeyID == vm.Id {
				return simtypes.NoOpMsg(
					did.ModuleName,
					TypeMsgAddVerification,
					"vm already exists, could not add verification method",
				), nil, nil
			}
		}

		vm := did.NewVerification(
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
			vm,
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

		if !found {
			// return an error that will not stop simulation as the did was not found or the verification method already exists
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgAddVerification,
				"did not found, could not add verification method",
			), nil, nil
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
		vm := didDoc.VerificationMethod

		// return an error that will not stop simulation as the did was not found or the verification method does not exists
		if !found {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgRevokeVerification,
				"did not found, could not remove verification method",
			), nil, nil
		}

		// return an error that will not stop simulation if the length of the VM array is 1
		if len(vm) == 1 {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgRevokeVerification,
				"could not remove verification method as it would break the did document",
			), nil, nil
		}

		msg := did.NewMsgRevokeVerification(
			didID.String(),
			vm[1].Id,
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

// SimulateMsgSetVerificationRelationships simulates a MsgSetVerificationRelationships message
func SimulateMsgSetVerificationRelationships(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))
		vm := didDoc.VerificationMethod

		// return an error that will not stop simulation as the did was not found or the verification method does not exists
		if !found {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgSetVerificationRelationships,
				"did not found, could not remove verification method",
			), nil, nil
		}

		// return an error that will not stop simulation if the length of the VM array is 1
		if len(vm) == 1 {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgSetVerificationRelationships,
				"could not remove verification method as it would break the did document",
			), nil, nil
		}

		msg := did.NewMsgSetVerificationRelationships(
			didID.String(),
			vm[1].Id,
			[]string{did.KeyAgreement, did.CapabilityDelegation},
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgSetVerificationRelationships,
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

// SimulateMsgAddService simulates a MsgAddService message
func SimulateMsgAddService(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		didDoc, found := k.GetDidDocument(ctx, []byte(didID))

		serviceID := "service:emtiagent" + fmt.Sprint(len(didDoc.Service))
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
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgAddService,
				"did not found, could not add service",
			), nil, nil
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
		service := didDoc.Service

		serviceID := "service:emtiagent" + fmt.Sprint(len(didDoc.Service))

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
		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		// return an error that will not stop simulation as the did was not found
		if !found {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgDeleteService,
				"did not found, could not remove service",
			), nil, nil
		}

		// return an error that will not stop simulation if the length of the Service array is 0
		if len(service) == 0 {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgDeleteService,
				"could not remove verification method as it would break the did document",
			), nil, nil
		}

		return opMsg, fOp, err
	}
}

// SimulateMsgAddController simulates a MsgAddController message
func SimulateMsgAddController(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didController, _ := simtypes.RandomAcc(r, accs)
		controllerAddress := didController.Address.String()

		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		controllerDidID := did.NewKeyDID(controllerAddress)
		_, found := k.GetDidDocument(ctx, []byte(didID))

		msg := did.NewMsgAddController(
			didID.String(),
			controllerDidID.String(),
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgAddController,
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
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgAddController,
				"did not found, could not add controller",
			), nil, nil
		}

		return opMsg, fOp, err
	}
}

// SimulateMsgDeleteController simulates a MsgDeleteController message
func SimulateMsgDeleteController(k keeper.Keeper, bk did.BankKeeper, ak did.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		didController, _ := simtypes.RandomAcc(r, accs)
		controllerAddress := didController.Address.String()

		didID := did.NewChainDID(ctx.ChainID(), ownerAddress)
		controllerDidID := did.NewKeyDID(controllerAddress)

		msg := did.NewMsgDeleteController(
			didID.String(),
			controllerDidID.String(),
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgDeleteController,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      did.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		didDoc, found := k.GetDidDocument(ctx, []byte(didID))

		// return an error that will not stop simulation as the did was not found
		if !found {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgDeleteController,
				"did not found, could not remove service",
			), nil, nil
		}

		controller := didDoc.Controller

		// return an error that will not stop simulation if the length of the Controller array is 0
		if len(controller) == 0 {
			return simtypes.NoOpMsg(
				did.ModuleName,
				TypeMsgDeleteController,
				"could not remove verification method as it would break the did document",
			), nil, nil
		}

		return opMsg, fOp, err
	}
}
