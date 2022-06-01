package types_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/elesto-dao/elesto/app"
	"github.com/elesto-dao/elesto/x/mint/types"
)

func TestParams(t *testing.T) {
	app.Setup(false);
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `mint_denom:"stake" inflation_rates:"1" inflation_rates:"0.5" inflation_rates:"0.25" inflation_rates:"0.125" inflation_rates:"0.0625" inflation_rates:"0.03125" inflation_rates:"0.02" inflation_rates:"0.02" inflation_rates:"0.02" inflation_rates:"0.02" blocks_per_year:6308000 max_supply:1000000000000000 team_address:"elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684" team_reward:"0.1" `
	require.Equal(t, paramsStr, defaultParams.String())
}

func TestParamsValidate(t *testing.T) {
	app.Setup(false);

	require.NoError(t, types.DefaultParams().Validate())

	testCases := []struct {
		name        string
		configure   func(*types.Params)
		expectedErr string
	}{
		{
			"valid default params",
			func(params *types.Params) {},
			"",
		},
		{
			"empty mint denom",
			func(params *types.Params) {
				params.MintDenom = ""
			},
			"mint denom cannot be blank",
		},
		{
			"invalid mint denom",
			func(params *types.Params) {
				params.MintDenom = "a"
			},
			"invalid denom: a",
		},
		{
			"nil inflation rates",
			func(params *types.Params) {
				params.InflationRates = nil
			},
			"inflation rates must be provided",
		},
		{
			"inflation rates has element with less than 0 value value",
			func(params *types.Params) {
				params.InflationRates = []string{"-1"}
			},
			"inflation must be a value greather than 0, got: -1.000000000000000000",
		},
		{
			"inflation rates has element with non-numerical value",
			func(params *types.Params) {
				params.InflationRates = []string{"random"}
			},
			"failed to set decimal string: random000000000000000000",
		},
		{
			"max supply is 0",
			func(params *types.Params) {
				params.MaxSupply = 0
			},
			"max supply must be greater than zero, got 0",
		},
		{
			"blocks per year are zero",
			func(params *types.Params) {
				params.BlocksPerYear = 0
			},
			"blocks per year must be positive, got 0",
		},
		{
			"team reward is not a valid number",
			func(params *types.Params) {
				params.TeamReward = "not valid"
			},
			"failed to set decimal string: not valid000000000000000000",
		},
		{
			"team reward is less than 0",
			func(params *types.Params) {
				params.TeamReward = "-1"
			},
			"team reward must be a value between 0 and 1, got: -1",
		},
		{
			"team reward is greater than 1",
			func(params *types.Params) {
				params.TeamReward = "1.1"
			},
			"team reward must be a value between 0 and 1, got: 1.1",
		},
		{
			"team address has non-elesto prefix",
			func(params *types.Params) {
				params.TeamAddress = "nonelesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu82xzwuk"
			},
			"invalid Bech32 prefix; expected elesto, got nonelesto",
		},
		{
			"malformed address",
			func(params *types.Params) {
				// here, the last six characters are wrong (the checksum, that is)
				params.TeamAddress = "elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu82xzwuk"
			},
			"decoding bech32 failed: invalid checksum (expected pwm684 got 2xzwuk)",
		},
		{
			"a fine address",
			func(params *types.Params) {
				params.TeamAddress = "elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684"
			},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := types.DefaultParams()
			tc.configure(&params)
			err := params.Validate()

			var err2 error
			for _, p := range params.ParamSetPairs() {
				err := p.ValidatorFn(reflect.ValueOf(p.Value).Elem().Interface())
				if err != nil {
					err2 = err
					break
				}
			}
			if tc.expectedErr != "" {
				require.EqualError(t, err, tc.expectedErr)
				require.EqualError(t, err2, tc.expectedErr)
			} else {
				require.Nil(t, err)
				require.Nil(t, err2)
			}
		})
	}
}
