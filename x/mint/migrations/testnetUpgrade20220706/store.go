package testnetUpgrade20220706

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	mintTypes "github.com/elesto-dao/elesto/v2/x/mint/types"
)

type ExpectedKeeper interface {
	SetParams(ctx sdk.Context, params mintTypes.Params)
	MintCoins(ctx sdk.Context, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

func Migrate(ctx sdk.Context, keeper ExpectedKeeper) error {
	// reset params, which now hardcodes the mint denom as "utsp"
	keeper.SetParams(ctx, mintTypes.DefaultParams())

	// get the "stake" minted tokens
	totalStakeSupply := keeper.GetSupply(ctx, "stake")
	if !totalStakeSupply.IsValid() {
		return fmt.Errorf("supply returned by GetSupply is not valid")
	}

	if totalStakeSupply.IsZero() {
		return nil // no "stake" tokens to migrate
	}

	toBeMintedUtsp := sdk.NewCoin("utsp", totalStakeSupply.Amount)

	// burn "stake" tokens
	if err := keeper.BurnCoins(ctx, sdk.NewCoins(totalStakeSupply)); err != nil {
		return fmt.Errorf("cannot burn stake coins, %w", err)
	}

	// mint utsp for the same amount
	if err := keeper.MintCoins(ctx, sdk.NewCoins(toBeMintedUtsp)); err != nil {
		return fmt.Errorf("cannot mint utsp coins, %w", err)
	}

	return nil
}
