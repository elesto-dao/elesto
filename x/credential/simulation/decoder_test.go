package simulation_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/elesto-dao/elesto/v4/x/credential"
	"github.com/elesto-dao/elesto/v4/x/credential/simulation"
)

var (
	//go:embed testdata/dummy.schema.json
	dummySchemaOk string
	//go:embed testdata/dummy.vocab.json
	dummyVocabOk string
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	credDef := &credential.CredentialDefinition{
		Id:           "test-id",
		PublisherId:  "test-published",
		Schema:       []byte(dummySchemaOk),
		Vocab:        []byte(dummyVocabOk),
		Name:         "Credential Definition 1",
		Description:  "This is a sample credential",
		SupersededBy: "",
		IsActive:     true,
	}

	credBz, err := cdc.Marshal(credDef)
	require.NoError(t, err)
	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: credential.CredentialDefinitionKey, Value: credBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectErr   bool
		expectedLog string
	}{
		{"Create Credential definition", false, fmt.Sprintf("%v\n%v", credDef, credDef)},
		{"other", true, ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectErr {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
