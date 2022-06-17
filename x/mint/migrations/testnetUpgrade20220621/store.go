package testnetUpgrade20220621

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/x/mint/types"
)

type ParamsKeeper interface {
	SetParams(ctx sdk.Context, params types.Params)
}

func MigrateParams(ctx sdk.Context, keeper ParamsKeeper) error {
	keeper.SetParams(ctx, types.DefaultParams())
	return nil
}
