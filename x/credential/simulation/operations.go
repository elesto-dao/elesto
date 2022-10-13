package simulation

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/keeper"
	"github.com/elesto-dao/elesto/v3/x/did"
)

var (
	//go:embed testdata/dummy.schema.json
	dummySchemaOk string
	//go:embed testdata/dummy.vocab.json
	dummyVocabOk string
)

var (
	TypeMsgPublishCredentialDefinition       = sdk.MsgTypeURL(&credential.MsgPublishCredentialDefinitionRequest{})
	TypeMsgPublishPublicVerifiableCredential = sdk.MsgTypeURL(&credential.MsgIssuePublicVerifiableCredentialRequest{})
	TypeMsgUpdateCredentialDefinition        = sdk.MsgTypeURL(&credential.MsgUpdateCredentialDefinitionRequest{})
)

const (
	opWeightMsgPublishCredentialDefinition       = "op_weight_msg_publish_credential_definition"
	opWeightMsgPublishPublicVerifiableCredential = "op_weight_msg_publish_public_verifiable_credential"
	opWeightMsgUpdateCredentialDefinition        = "op_weight_msg_update_credential_definition"

	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishCredentialDefinition       int = 100
	defaultWeightMsgPublishPublicVerifiableCredential int = 100
	defaultWeightMsgUpdateCredentialDefinition        int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// WeightedOperations returns the all the gov module operations with their respective weights.
func WeightedOperations(simState module.SimulationState, keeper keeper.Keeper, dk credential.DidKeeper, bk credential.BankKeeper, ak credential.AccountKeeper) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgPublishCredentialDefinition int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishCredentialDefinition, &weightMsgPublishCredentialDefinition, nil,
		func(_ *rand.Rand) {
			weightMsgPublishCredentialDefinition = defaultWeightMsgPublishCredentialDefinition
		},
	)

	var weightMsgPublishPublicVerifiableCredential int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishPublicVerifiableCredential, &weightMsgPublishPublicVerifiableCredential, nil,
		func(_ *rand.Rand) {
			weightMsgPublishPublicVerifiableCredential = defaultWeightMsgPublishPublicVerifiableCredential
		},
	)

	var weightMsgUpdateCredentialDefinition int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCredentialDefinition, &weightMsgUpdateCredentialDefinition, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCredentialDefinition = defaultWeightMsgUpdateCredentialDefinition
		},
	)

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishCredentialDefinition,
		SimulateMsgPublishCredentialDefinition(keeper, dk, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishPublicVerifiableCredential,
		SimulateMsgPublishPublicVerifiableCredential(keeper, dk, bk, ak),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCredentialDefinition,
		SimulateMsgUpdateCredentialDefinition(keeper, dk, bk, ak),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// SimulateMsgPublishCredentialDefinition simulates a MsgPublishCredentialDefinition message
func SimulateMsgPublishCredentialDefinition(k keeper.Keeper, dk credential.DidKeeper, bk credential.BankKeeper, ak credential.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()

		publisherDID := did.NewKeyDID(ownerAddress)
		didDocument, _ := did.NewDidDocument(publisherDID.String())
		dk.SetDidDocument(ctx, []byte(publisherDID.String()), didDocument)

		credDef := &credential.CredentialDefinition{
			Id:           fmt.Sprintf("id-%s", ownerAddress),
			PublisherId:  publisherDID.String(),
			Schema:       []byte(dummySchemaOk),
			Vocab:        []byte(dummyVocabOk),
			Name:         "Credential Definition 1",
			Description:  "This is a sample credential",
			SupersededBy: "",
			IsActive:     true,
		}

		_, found := k.GetCredentialDefinition(ctx, credDef.Id)
		if found {
			return simtypes.NoOpMsg(credential.ModuleName, TypeMsgPublishCredentialDefinition, "credential definition already exists"), nil, nil
		}

		msg := credential.NewMsgPublishCredentialDefinitionRequest(
			credDef,
			ownerAddress,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgPublishCredentialDefinition,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      credential.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		_, found = k.GetCredentialDefinition(ctx, credDef.Id)
		if !found {
			return simtypes.NoOpMsg(credential.ModuleName, TypeMsgPublishCredentialDefinition, "credential not found, could not create credential definition"), nil, nil
		}

		if r.Intn(1000)%2 == 0 {
			k.AllowPublicCredential(ctx, credDef.Id)
		}

		return opMsg, fOp, err
	}
}

// SimulateMsgUpdateCredentialDefinition simulates a MsgUpdateCredentialDefinition message
func SimulateMsgUpdateCredentialDefinition(k keeper.Keeper, dk credential.DidKeeper, bk credential.BankKeeper, ak credential.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)

		ownerAddress := didOwner.Address.String()
		credDefID := fmt.Sprintf("id-%s", ownerAddress)

		_, found := k.GetCredentialDefinition(ctx, credDefID)
		if !found {
			return simtypes.NoOpMsg(credential.ModuleName, TypeMsgUpdateCredentialDefinition, "credential definition does not exist"), nil, nil
		}

		msg := credential.NewMsgUpdateCredentialDefinitionRequest(
			true,
			"",
			ownerAddress,
		)
		msg.CredentialDefinitionID = credDefID

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgUpdateCredentialDefinition,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      credential.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		return opMsg, fOp, err
	}
}

// SimulateMsgPublishPublicVerifiableCredential simulates a MsgPublishPublicVerifiableCredential message
func SimulateMsgPublishPublicVerifiableCredential(k keeper.Keeper, dk credential.DidKeeper, bk credential.BankKeeper, ak credential.AccountKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		didOwner, _ := simtypes.RandomAcc(r, accs)
		ownerAddress := didOwner.Address.String()
		issuerDID := did.NewChainDID(chainID, ownerAddress)

		didDocument, err := did.NewDidDocument(issuerDID.String(), did.WithVerifications(
			did.NewVerification(
				did.NewVerificationMethod(
					issuerDID.NewVerificationMethodID(ownerAddress),
					"did:cosmos:elesto:subject",
					did.NewBlockchainAccountID(chainID, ownerAddress),
					did.CosmosAccountAddress,
				),
				[]string{
					did.AssertionMethod,
				},
				nil,
			),
		),
		)
		if err != nil {
			panic(err)
		}

		dk.SetDidDocument(ctx, []byte(issuerDID.String()), didDocument)

		_, err = ak.GetPubKey(ctx, didOwner.Address)
		if err != nil {
			panic(err)
		}

		credDefID := fmt.Sprintf("id-%s", ownerAddress)

		_, found := k.GetCredentialDefinition(ctx, credDefID)
		if !found {
			return simtypes.NoOpMsg(credential.ModuleName, TypeMsgPublishPublicVerifiableCredential, "credential definition does not exist"), nil, nil
		}

		allowed := k.IsPublicCredentialDefinitionAllowed(ctx, credDefID)
		if !allowed {
			return simtypes.NoOpMsg(credential.ModuleName, TypeMsgPublishPublicVerifiableCredential, "credential definition not allowed"), nil, nil
		}

		pwc, err := credential.NewWrappedCredential(
			credential.NewPublicVerifiableCredential("https://example.credential/01",
				credential.WithType("SpecialCredential"),
				credential.WithIssuerDID(issuerDID),
				credential.WithIssuanceDate(time.Now())),
		)

		if err != nil {
			panic(err)
		}

		err = pwc.SetSubject(map[string]any{"id": "https://something.something"})
		if err != nil {
			panic(err)
		}
		vmID := pwc.GetIssuerDID().NewVerificationMethodID(ownerAddress)
		data, err := pwc.GetBytes()
		if err != nil {
			panic(err)
		}
		signature, err := didOwner.PrivKey.Sign(data)
		if err != nil {
			panic(err)
		}

		pwc.Proof = credential.NewProof(
			didOwner.PubKey.Type(),
			time.Now().Format(time.RFC3339),
			did.AssertionMethod,
			vmID,
			base64.StdEncoding.EncodeToString(signature),
		)

		msg := credential.NewMsgIssuePublicVerifiableCredentialRequest(
			pwc.PublicVerifiableCredential,
			credDefID,
			didOwner.Address,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         TypeMsgPublishPublicVerifiableCredential,
			Context:         ctx,
			SimAccount:      didOwner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      credential.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		opMsg, fOp, err := simulation.GenAndDeliverTxWithRandFees(txCtx)

		return opMsg, fOp, err
	}
}
