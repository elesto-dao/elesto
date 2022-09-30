package types

const (
	// module name
	ModuleName = "mint"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the mint store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the mint querier
	QueryParameters = "parameters"
)

var (
	// BootstrapDateKey is the key for the bootstrap_date fields, used to determine the
	// current minting year.
	BootstrapDateKey       = []byte{0x42}
	BootstrapDateCanaryKey = []byte{0x43}
)
