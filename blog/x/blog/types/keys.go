package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"
)

var (
	ParamsKey = []byte("p_blog")

	// PostKey is used to uniquely identify posts within the system.
	// It will be used as the beginning of the key for each post, followed by their unique ID
	PostKey = "Post/value/"
	PoslKey = "Posl/value/"

	// This key will be used to keep track of the ID of the latest post added to the store.
	PostCountKey = "Post/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}