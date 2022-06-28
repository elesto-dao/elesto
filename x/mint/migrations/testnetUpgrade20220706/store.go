package testnetUpgrade20220706

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	mintTypes "github.com/elesto-dao/elesto/v2/x/mint/types"
)

type ExpectedKeeper interface {
	SetParams(ctx sdk.Context, params mintTypes.Params)
	MintCoins(ctx sdk.Context, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

func Migrate(ctx sdk.Context, keeper ExpectedKeeper) error {
	// reset params, which now hardcodes the mint denom as "utsp"
	keeper.SetParams(ctx, mintTypes.DefaultParams())

	return nil
}
