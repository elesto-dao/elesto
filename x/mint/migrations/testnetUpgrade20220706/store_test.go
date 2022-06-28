package testnetUpgrade20220706_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elesto-dao/elesto/v2/app"
	"github.com/elesto-dao/elesto/v2/x/mint/migrations/testnetUpgrade20220706"
	mintTypes "github.com/elesto-dao/elesto/v2/x/mint/types"
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

func (b *bogusMigrationKeeper) BurnCoins(ctx sdk.Context, amt sdk.Coins) error {
	if b.failBurn {
		return fmt.Errorf("cannot burn")
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
	tests := []struct {
		name           string
		mutate         func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins)
		existingSupply sdk.Coins
		wantErr        bool
		bogusKeeper    *bogusMigrationKeeper
	}{
		{
			"params had a different denom than stake",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, _ sdk.Coins) {
				p := mintTypes.DefaultParams()
				p.MintDenom = "other"
				k.SetParams(ctx, p)
			},
			sdk.NewCoins(),
			false,
			nil,
		},
		{
			"there are `stake` tokens and utsp tokens",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
				k.MintCoins(ctx, existingSupply)
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("stake", 100),
				sdk.NewInt64Coin("utsp", 100),
			),
			false,
			nil,
		},
		{
			"there are `stake` tokens only",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
				k.MintCoins(ctx, existingSupply)
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("stake", 100),
			),
			false,
			nil,
		},
		{
			"there are `utsp` tokens only",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
				k.MintCoins(ctx, existingSupply)
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("utsp", 100),
			),
			false,
			nil,
		},
		{
			"fail get supply",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("utsp", 100),
			),
			true,
			&bogusMigrationKeeper{
				failGetSupply: true,
			},
		},
		{
			"fail burn",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("stake", 100),
			),
			true,
			&bogusMigrationKeeper{
				failBurn: true,
			},
		},
		{
			"fail mint",
			func(ctx sdk.Context, k testnetUpgrade20220706.ExpectedKeeper, existingSupply sdk.Coins) {
			},
			sdk.NewCoins(
				sdk.NewInt64Coin("stake", 100),
			),
			true,
			&bogusMigrationKeeper{
				failMint: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := app.Setup(false)
			ctx := app.BaseApp.NewContext(false, tmproto.Header{})

			var keeper testnetUpgrade20220706.ExpectedKeeper
			keeper = app.MintKeeper

			tt.mutate(ctx, app.MintKeeper, tt.existingSupply)

			if tt.bogusKeeper != nil {
				tt.bogusKeeper.supply = tt.existingSupply
				keeper = tt.bogusKeeper
			}

			err := testnetUpgrade20220706.Migrate(ctx, keeper)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			params := app.MintKeeper.GetParams(ctx)
			require.Equal(t, "utsp", params.MintDenom)

			stakeSupply := app.MintKeeper.GetSupply(ctx, "stake")
			require.True(t, stakeSupply.IsZero())

			if !tt.existingSupply.Empty() {
				// check that the sum of all existingSupply tokens matches the supply of utsp
				sum := sdk.NewInt(0)
				for _, s := range tt.existingSupply {
					sum = sum.Add(s.Amount)
				}

				utspSupply := app.MintKeeper.GetSupply(ctx, "utsp")
				require.Equal(t, sum.Int64(), utspSupply.Amount.Int64())
			}
		})
	}
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