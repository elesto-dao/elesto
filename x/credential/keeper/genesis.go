package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

func InitGenesis(ctx sdk.Context, k Keeper, genState *credential.GenesisState) {
	for _, id := range genState.AllowedCredentialIds {
		//_, found := k.GetCredentialDefinition(ctx, id)
		//if !found {
		//	panic(fmt.Sprintf("credential id %s not found", id))
		//}
		//
		//allowed := k.IsPublicCredentialIDAllowed(ctx, id)
		//if allowed {
		//	panic(fmt.Sprintf("credential id %s already allowed", id))
		//}
		k.AllowPublicCredential(ctx, id)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) *credential.GenesisState {
	var genState credential.GenesisState

	allowedCdIterator := k.GetAll(ctx, credential.PublicCredentialAllowKey)
	defer allowedCdIterator.Close()

	for ; allowedCdIterator.Valid(); allowedCdIterator.Next() {
		id := string(allowedCdIterator.Value())
		genState.AllowedCredentialIds = append(genState.AllowedCredentialIds, id)
	}

	return &genState
}
