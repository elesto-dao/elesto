package did

import (
	"fmt"
)

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate validates all the did documents in the genesi file
func (s GenesisState) Validate() error {
	for _, didDoc := range s.DidDocuments {
		if !IsValidDIDDocument(didDoc) {
			return fmt.Errorf(
				"invalid did document in genesis state, %s", didDoc.Id,
			)
		}
	}

	return nil
}
