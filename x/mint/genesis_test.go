package mint_test

import (
	"testing"
	"time"

	chain "github.com/elesto-dao/elesto/v4/app"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/elesto-dao/elesto/v4/x/mint"
	"github.com/elesto-dao/elesto/v4/x/mint/types"
)

func (s *ModuleTestSuite) TestDefaultInitGenesis() {
	genState := *types.DefaultGenesisState()

	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState)
	// assign some values to the bootstrap date

	got := *mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(genState.Params, got.Params)
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
	ctx := app.BaseApp.NewContext(false, tmproto.Header{}).WithBlockHeight(12).WithBlockTime(time.Now())

	genState := mint.ExportGenesis(ctx, app.MintKeeper)
	bz := app.AppCodec().MustMarshalJSON(genState)

	var genState2 types.GenesisState
	app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	mint.InitGenesis(ctx, app.MintKeeper, app.AccountKeeper, &genState2)

	genState3 := mint.ExportGenesis(ctx, app.MintKeeper)
	require.Equal(t, *genState, genState2)
	require.Equal(t, genState2, *genState3)
}
