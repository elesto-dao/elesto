package credentials

import (
	"fmt"

	"github.com/elesto-dao/elesto/x/did"
)

// NewCredentialDefinitionFromFile create a credential definition by reading the data from a file
func NewCredentialDefinitionFromFile(did, publisherDID did.DID,
	name, description string,
	isPublic, isActive bool,
	schemaFile, vocabFile string) (*CredentialDefinition, error) {

	def := &CredentialDefinition{
		Id:          did.String(),
		PublisherId: publisherDID.String(),
		Name:        name,
		Description: description,
		IsPublic:    isPublic,
		IsActive:    isActive,
	}

	var err error

	if def.Schema, err = CompactJSON(schemaFile); err != nil {
		err = fmt.Errorf("error reading schema file: %w ", err)
		return nil, err
	}
	if def.Vocab, err = CompactJSON(vocabFile); err != nil {
		err = fmt.Errorf("error reading vocab file: %w ", err)
		return nil, err
	}

	return def, nil
}
