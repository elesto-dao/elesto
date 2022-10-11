package module_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	elestoapp "github.com/elesto-dao/elesto/v4/app"
)

func TestCreateModuleInApp(t *testing.T) {
	app := elestoapp.New(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		make(map[int64]bool),
		elestoapp.DefaultNodeHome,
		0,
		elestoapp.MakeEncodingConfig(elestoapp.ModuleBasics),
		simapp.EmptyAppOptions{},
	)

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	require.NotNil(t, app.DidKeeper)
}
