package testnetUpgrade20220706

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	mintTypes "github.com/elesto-dao/elesto/v3/x/mint/types"
)

type ExpectedKeeper interface {
	SetParams(ctx sdk.Context, params mintTypes.Params)
	MintCoins(ctx sdk.Context, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

func Migrate(ctx sdk.Context, keeper ExpectedKeeper) error {
	// reset params, and hardcode the mint denom as "utsp"
	p := mintTypes.DefaultParams()
	p.MintDenom = "utsp"
	keeper.SetParams(ctx, p)

	return nil
}
