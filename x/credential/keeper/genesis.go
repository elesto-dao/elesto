package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

func InitGenesis(ctx sdk.Context, k Keeper, data *credential.GenesisState) {
	//TODO: validate data in genesis
	for i := range data.CredentialDefinitions {
		k.SetCredentialDefinition(ctx, &data.CredentialDefinitions[i])
	}

	for i := range data.PublicVerifiableCredentials {
		k.SetPublicCredential(ctx, &data.PublicVerifiableCredentials[i])
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) *credential.GenesisState {
	var genState credential.GenesisState
	cdsIterator := k.GetAll(ctx, credential.CredentialDefinitionKey)
	defer cdsIterator.Close()

	for ; cdsIterator.Valid(); cdsIterator.Next() {
		var cd credential.CredentialDefinition
		ok := k.Unmarshal(cdsIterator.Value(), &cd)
		if !ok {
			panic(fmt.Errorf("cannot unmarshal %s", cdsIterator.Value()))
		}

		genState.CredentialDefinitions = append(genState.CredentialDefinitions, cd)
	}

	pvcIterator := k.GetAll(ctx, credential.PublicCredentialKey)
	defer pvcIterator.Close()

	for ; pvcIterator.Valid(); pvcIterator.Next() {
		var pvc credential.PublicVerifiableCredential
		ok := k.Unmarshal(pvcIterator.Value(), &pvc)
		if !ok {
			panic(fmt.Errorf("cannot unmarshal %s", pvcIterator.Value()))
		}

		genState.PublicVerifiableCredentials = append(genState.PublicVerifiableCredentials, pvc)
	}

	return &genState
}
