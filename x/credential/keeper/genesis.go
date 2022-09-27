package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

func InitGenesis(ctx sdk.Context, k Keeper, genState *credential.GenesisState) {
	for _, id := range genState.AllowedCredentialIds {
		found := k.Exists(ctx, credential.PublicCredentialAllowKey, []byte(id))
		if !found {
			panic("credential id not found")
		}

		k.SetAllowedPublicCredential(ctx, id)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) *credential.GenesisState {
	var genState credential.GenesisState

	allowedCds := k.GetAll(ctx, credential.PublicCredentialAllowKey)
	defer allowedCds.Close()

	for ; allowedCds.Valid(); allowedCds.Next() {
		id := string(allowedCds.Value())
		genState.AllowedCredentialIds = append(genState.AllowedCredentialIds, id)
	}

	return &genState
}
