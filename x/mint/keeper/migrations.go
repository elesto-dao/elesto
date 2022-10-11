package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/mint/migrations/testnetUpgrade20220621"
	"github.com/elesto-dao/elesto/v4/x/mint/migrations/testnetUpgrade20220706"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return testnetUpgrade20220621.MigrateParams(ctx, m.keeper)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return testnetUpgrade20220706.Migrate(ctx, m.keeper)
}
