package credential

var (
	// CredentialDefinitionKey prefix for each key of a PublicCredential
	CredentialDefinitionKey = []byte{0x71}
	// PublicCredentialKey prefix for each key of a PublicCredential
	PublicCredentialKey = []byte{0x72}
	// PublicCredentialAllowed prefix
	PublicCredentialAllowKey = []byte{0x73}
)

const (
	// ModuleName defines the module name
	ModuleName = "credential"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability_credentials"
)
