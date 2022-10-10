package credential

import "fmt"

// NewGenesisState creates a new GenesisState object
func NewGenesisState(allowedIds ...string) *GenesisState {
	return &GenesisState{
		AllowedCredentialIds: allowedIds,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState()
}

func (gs *GenesisState) ValidateGenesis() error {
	exists := map[string]bool{}
	for _, id := range gs.AllowedCredentialIds {
		if id == "" {
			return fmt.Errorf("invalid id")
		}

		if _, found := exists[id]; found {
			return fmt.Errorf("id %s already present", id)
		}

		exists[id] = true
	}

	return nil
}
