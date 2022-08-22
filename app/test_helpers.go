package app

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// used here because we don't need app options in testing
type emptyAppOption struct{}

func (*emptyAppOption) Get(i string) interface{} {
	return nil
}

func setup(withGenesis bool, invCheckPeriod uint) (*App, GenesisState) {
	if config := sdk.GetConfig(); config.GetBech32AccountAddrPrefix() != AccountAddressPrefix {
		cosmoscmd.SetPrefixes(AccountAddressPrefix)
	}

	db := dbm.NewMemDB()
	encCdc := MakeEncodingConfig(ModuleBasics)
	app := New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, invCheckPeriod, encCdc, &emptyAppOption{})

	if withGenesis {
		return app, NewDefaultGenesisState(encCdc.Marshaler)
	}

	return app, GenesisState{}
}

// Setup initializes a new App. A Nop logger is set in App.
func Setup(isCheckTx bool) *App {
	app, genesisState := setup(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators: []abci.ValidatorUpdate{},
				ConsensusParams: &abci.ConsensusParams{
					Block: &abci.BlockParams{
						MaxBytes: 200000,
						MaxGas:   2000000,
					},
					Evidence: &tmproto.EvidenceParams{
						MaxAgeNumBlocks: 302400,
						MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
						MaxBytes:        10000,
					},
					Validator: &tmproto.ValidatorParams{
						PubKeyTypes: []string{
							tmtypes.ABCIPubKeyTypeEd25519,
						},
					},
				},
				AppStateBytes: stateBytes,
			},
		)
	}

	return app
}
