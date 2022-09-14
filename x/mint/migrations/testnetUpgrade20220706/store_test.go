package testnetUpgrade20220706_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elesto-dao/elesto/v3/app"
	"github.com/elesto-dao/elesto/v3/x/mint/migrations/testnetUpgrade20220706"
	mintTypes "github.com/elesto-dao/elesto/v3/x/mint/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func (b *bogusMigrationKeeper) SetParams(ctx sdk.Context, params mintTypes.Params) {
	if b.failSetParams {
		panic("cannot set params")
	}
}

func (b *bogusMigrationKeeper) MintCoins(ctx sdk.Context, amt sdk.Coins) error {
	if b.failMint {
		return fmt.Errorf("cannot mint")
	}

	return nil
}

func (b *bogusMigrationKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	if b.failGetSupply {
		c := sdk.Coin{}
		c.Denom = "a"

		return c
	}
	for _, t := range b.supply {
		if t.Denom == denom {
			return t
		}
	}

	return sdk.Coin{}
}

type bogusMigrationKeeper struct {
	failSetParams bool
	failMint      bool
	failBurn      bool
	failGetSupply bool
	supply        sdk.Coins
}

func TestMigrate(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	testnetUpgrade20220706.Migrate(ctx, app.MintKeeper)

	require.Equal(t, "utsp", app.MintKeeper.GetParams(ctx).MintDenom)
	require.Equal(t, "stake", mintTypes.DefaultParams().MintDenom)
}

func TestFailSetParams(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	keeper := &bogusMigrationKeeper{
		failSetParams: true,
	}

	require.Panics(t, func() {
		_ = testnetUpgrade20220706.Migrate(ctx, keeper)
	})
}
