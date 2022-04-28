package types_test

import (
	"github.com/elesto-dao/elesto/x/mint/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesisState_Validate(t *testing.T) {
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
