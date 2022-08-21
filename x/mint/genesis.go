package mint

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v2/x/mint/keeper"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	if err := types.ValidateGenesis(*data); err != nil {
		panic(err)
	}

	keeper.SetParams(ctx, data.Params)
	if err := keeper.SetBootstrapDateCanary(ctx, data.BootstrapDateCanary, true); err != nil {
		panic(err)
	}

	if keeper.BootstrapDateCanarySet(ctx) {
		var pt time.Time
		if err := pt.UnmarshalText([]byte(data.BootstrapDate)); err != nil {
			panic(err)
		}

		if err := keeper.SetBootstrapDate(ctx.WithBlockTime(pt), true); err != nil {
			panic(err)
		}
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	bd, err := keeper.BootstrapDate(ctx)
	if err != nil {
		panic(fmt.Errorf("cannot export mint genesis, %w", err))
	}

	bdText, err := bd.MarshalText()
	if err != nil {
		panic(fmt.Errorf("cannot export mint genesis, cannot format bootstrap date, %w", err))
	}

	return types.NewGenesisState(params, string(bdText), keeper.BootstrapDateCanarySet(ctx))
}
