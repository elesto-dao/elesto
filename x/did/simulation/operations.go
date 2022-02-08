package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

		didDoc, found := k.GetDidDocument(ctx, []byte(didID))
		d, _ := did.NewDidDocument(didID.String(), did.WithVerifications(auth))
		msg := did.NewMsgCreateDidDocument(
			didID.String(),
			did.Verifications{auth},
			did.Services{},
			ownerAddress,
		)

		if found {
			// check the store matches the input
			if !didDoc.Equal(d) {
				// return an error that will stop simulation
				return simtypes.NoOpMsg(
						did.ModuleName,
						TypeMsgCreateDidDocument,
						"did found but not matched",
					),
					nil,
					sdkerrors.Wrapf(sdkerrors.ErrNotFound, "did found but not matched")

			}
			return simtypes.NoOpMsg(did.ModuleName, TypeMsgCreateDidDocument, "dids found"), nil, nil
		}

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
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
