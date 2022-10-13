package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/elesto-dao/elesto/v3/x/credential"
)

// NewDecodeStore returns a decoder function closure that umarshals the KVPair's
// Value to the corresponding did type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], credential.CredentialDefinitionKey):
			var credA, credB credential.CredentialDefinition
			cdc.MustUnmarshal(kvA.Value, &credA)
			cdc.MustUnmarshal(kvB.Value, &credB)
			return fmt.Sprintf("%v\n%v", credA.String(), credB.String())
		default:
			panic(fmt.Sprintf("invalid credential definition key %X", kvA.Key))
		}
	}
}
