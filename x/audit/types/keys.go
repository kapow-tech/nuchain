package types

const (
	// ModuleName defines the module name
	ModuleName = "audit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_audit"
)

var (
	ParamsKey = []byte("p_audit")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
