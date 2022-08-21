package types

import (
	"fmt"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(params Params, bootstrapDate string, bootstrapDateCanary bool) *GenesisState {
	return &GenesisState{
		BootstrapDate:       bootstrapDate,
		BootstrapDateCanary: bootstrapDateCanary,
		Params:              params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	// default genesis state has empty bootstrap date
	return NewGenesisState(DefaultParams(), "", false)
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return fmt.Errorf("mint genesis validation failed, %w", err)
	}

	return nil
}
