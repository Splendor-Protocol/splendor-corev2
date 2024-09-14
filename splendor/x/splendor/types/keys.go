package types

const (
	// ModuleName defines the module name
	ModuleName = "splendor"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_splendor"
)

var (
	ParamsKey = []byte("p_splendor")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
