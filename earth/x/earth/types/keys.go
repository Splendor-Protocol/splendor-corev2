package types

const (
	// ModuleName defines the module name
	ModuleName = "earth"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_earth"
)

var (
	ParamsKey = []byte("p_earth")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
