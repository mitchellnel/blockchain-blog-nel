package types

const (
	// ModuleName defines the module name
	ModuleName = "blockchainblognel"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blockchainblognel"

	// PostKey defines the key for the store of Posts
	PostKey = "Post-value-"

	// PostCountKey defines the count of Posts in the store
	PostCountKey = "Post-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
