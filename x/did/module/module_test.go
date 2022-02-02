package did_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	elestoapp "github.com/elesto-dao/elesto/app"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp"

	dbm "github.com/tendermint/tm-db"
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
		cosmoscmd.MakeEncodingConfig(elestoapp.ModuleBasics),
		simapp.EmptyAppOptions{},
	)

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	require.NotNil(t, app.(*elestoapp.App).DidKeeper)
}
