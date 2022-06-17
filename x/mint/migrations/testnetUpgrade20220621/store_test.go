package testnetUpgrade20220621_test

import (
	"testing"

	"github.com/elesto-dao/elesto/app"
	"github.com/elesto-dao/elesto/x/mint/migrations/testnetUpgrade20220621"
	"github.com/elesto-dao/elesto/x/mint/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestMigrateParams(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	testnetUpgrade20220621.MigrateParams(ctx, app.MintKeeper)

	require.Equal(t, types.DefaultParams(), app.MintKeeper.GetParams(ctx))
}
