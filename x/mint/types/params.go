package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyMintDenom     = []byte("MintDenom")
	KeyMaxSupply     = []byte("MaxSupply")
	KeyBlocksPerYear = []byte("BlocksPerYear")
	KeyTeamReward    = []byte("TeamReward")
	KeyTeamAddress   = []byte("TeamAddress")
)

// ParamKeyTable for mint module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams default mint module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:     sdk.DefaultBondDenom,
		BlocksPerYear: 6_308_000,
		MaxSupply:     1_000_000_000_000_000,
		TeamReward:    "0.1",
		TeamAddress:   "elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684",
	}
}

// Validate validate params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateMaxSupply(p.MaxSupply); err != nil {
		return err
	}
	if err := validateBlocksPerYear(p.BlocksPerYear); err != nil {
		return err
	}
	if err := validateTeamReward(p.TeamReward); err != nil {
		return err
	}
	if err := validateTeamAddress(p.TeamAddress); err != nil {
		return err
	}

	return nil

}

// ParamSetPairs Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyMaxSupply, &p.MaxSupply, validateMaxSupply),
		paramtypes.NewParamSetPair(KeyBlocksPerYear, &p.BlocksPerYear, validateBlocksPerYear),
		paramtypes.NewParamSetPair(KeyTeamAddress, &p.TeamAddress, validateTeamAddress),
		paramtypes.NewParamSetPair(KeyTeamReward, &p.TeamReward, validateTeamReward),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateMaxSupply(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("max supply must be greater than zero, got %d", v)
	}
	return nil
}

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("blocks per year must be positive, got %d", v)
	}
	return nil
}

func validateTeamReward(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	d, err := sdk.NewDecFromStr(v)
	if err != nil {
		return err
	}

	if d.LT(sdk.NewDec(0)) || d.GT(sdk.NewDec(1)) {
		return fmt.Errorf("team reward must be a value between 0 and 1, got: %s", v)
	}
	return nil
}

func validateTeamAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := sdk.AccAddressFromBech32(v)

	return err
}
