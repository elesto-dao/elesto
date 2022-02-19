package credentials

var (
	// CredentialIssuerKey prefix for each key to a CredentialIssuer
	CredentialIssuerKey = []byte{0x61}
	// PublicCredentialKey prefix for each key of a PublicCredential
	PublicCredentialKey = []byte{0x62}
)

const (
	// ModuleName defines the module name
	ModuleName = "credentials"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability_credentials"
)
