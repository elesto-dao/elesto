package types_test

import (
	"testing"

	"github.com/elesto-dao/elesto/v4/app"
	"github.com/elesto-dao/elesto/v4/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	app.Setup(false)

	for _, tc := range []struct {
		name        string
		malleate    func(genState *types.GenesisState)
		expectedErr string
	}{
		{
			"default is valid",
			func(genState *types.GenesisState) {},
			"",
		},
		{
			"invalid params",
			func(genState *types.GenesisState) {
				genState.Params.MintDenom = ""
			},
			"mint genesis validation failed, mint denom cannot be blank",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			genState := types.NewGenesisState(types.DefaultParams())
			tc.malleate(genState)
			err := types.ValidateGenesis(*genState)
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
