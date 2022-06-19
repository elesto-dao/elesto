package mint_test

import (
	"testing"

	chain "github.com/elesto-dao/elesto/v2/app"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/elesto-dao/elesto/v2/x/mint"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

func TestInitGenesis(t *testing.T) {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// default gent state case
	genState := types.DefaultGenesisState()
	mint.InitGenesis(ctx, app.MintKeeper, app.AccountKeeper, genState)
	got := mint.ExportGenesis(ctx, app.MintKeeper)
	require.Equal(t, *genState, *got)
}

func TestInitInvalidGenesis(t *testing.T) {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// default gent state case
	genState := types.DefaultGenesisState()
	genState.Params.TeamAddress = ""

	require.Panics(t, func() {
		mint.InitGenesis(ctx, app.MintKeeper, app.AccountKeeper, genState)
	})
}

func TestImportExportGenesis(t *testing.T) {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	genState := mint.ExportGenesis(ctx, app.MintKeeper)
	bz := app.AppCodec().MustMarshalJSON(genState)

	var genState2 types.GenesisState
	app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	mint.InitGenesis(ctx, app.MintKeeper, app.AccountKeeper, &genState2)

	genState3 := mint.ExportGenesis(ctx, app.MintKeeper)
	require.Equal(t, *genState, genState2)
	require.Equal(t, genState2, *genState3)
}
