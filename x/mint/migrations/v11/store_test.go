package v11_test

import (
	"testing"

	"github.com/elesto-dao/elesto/app"
	v11 "github.com/elesto-dao/elesto/x/mint/migrations/v11"
	"github.com/elesto-dao/elesto/x/mint/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestMigrateParams(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	v11.MigrateParams(ctx, app.MintKeeper)

	require.Equal(t, types.DefaultParams(), app.MintKeeper.GetParams(ctx))
}
