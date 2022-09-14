package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/elesto-dao/elesto/v3/x/did"
)

// InitGenesis initializes the genesis state for the did module
func (k Keeper) InitGenesis(
	ctx types.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var genesisState did.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	for _, didDoc := range genesisState.DidDocuments {
		// set the did document in the store
		k.SetDidDocument(ctx, []byte(didDoc.Id), *didDoc)
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the did state to a genesis.json file
func (k Keeper) ExportGenesis(ctx types.Context, cdc codec.JSONCodec) *did.GenesisState {
	dids := k.GetAllDidDocuments(ctx)

	return &did.GenesisState{
		DidDocuments: dids,
	}
}
