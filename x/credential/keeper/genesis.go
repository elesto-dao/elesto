package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/credential"
)

func InitGenesis(ctx sdk.Context, k Keeper, genState *credential.GenesisState) {
	for i := range genState.CredentialDefinitions {
		k.SetCredentialDefinition(ctx, &genState.CredentialDefinitions[i])
	}

	for _, id := range genState.PublicCredentialDefinitionsIDs {
		_, found := k.GetCredentialDefinition(ctx, id)
		if !found {
			panic(fmt.Sprintf("credential id %s not found", id))
		}

		allowed := k.IsPublicCredentialDefinitionAllowed(ctx, id)
		if allowed {
			panic(fmt.Sprintf("credential id %s already allowed", id))
		}
		k.AllowPublicCredential(ctx, id)
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

	allowedCdIterator := k.GetAll(ctx, credential.PublicCredentialAllowKey)
	defer allowedCdIterator.Close()

	for ; allowedCdIterator.Valid(); allowedCdIterator.Next() {
		id := string(allowedCdIterator.Value())
		genState.PublicCredentialDefinitionsIDs = append(genState.PublicCredentialDefinitionsIDs, id)
	}

	return &genState
}
