package types_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/elesto-dao/elesto/v2/app"
	"github.com/elesto-dao/elesto/v2/x/mint/types"
)

func TestParams(t *testing.T) {
	app.Setup(false)
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `mint_denom:"stake" team_address:"elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684" `
	require.Equal(t, paramsStr, defaultParams.String())
}

func TestParamsValidate(t *testing.T) {
	app.Setup(false)

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
