package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/mint/keeper"
	"github.com/elesto-dao/elesto/x/mint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	if err := types.ValidateGenesis(*data); err != nil {
		panic(err)
	}
	// init to prevent nil slice, []types.InflationSchedule(nil)
	// TODO: empty inflation reates should fail the ValidateGenesis
	// this part should not be necessary
	if data.Params.InflationRates == nil {
		data.Params.InflationRates = make([]string, 0)
	}
	keeper.SetParams(ctx, data.Params)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	// init to prevent nil slice, []types.InflationSchedule(nil)
	if params.InflationRates == nil || len(params.InflationRates) == 0 {
		params.InflationRates = make([]string, 0)
	}
	return types.NewGenesisState(params)
}
