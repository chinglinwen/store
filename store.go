// Store is for storing arbitrary data.
// It support oss as backend for now.
package store

import "sync"

type Reader interface {
	Read(key string) ([]byte, error)
}

type Writer interface {
	Write(key string, value []byte) error
}

type Store interface {
	Reader
	Writer
}

// Create a new store for read and write,
// Backend is one of registered backends.
func New(backend, bucket string) (Store, error) {
	return backends[backend].New(bucket)
}

type Newer interface {
	New(string) (Store, error)
}

var (
	newerMu  sync.RWMutex
	backends = make(map[string]Newer)
)

// Register is for registering the backend (usually in init function).
func Register(name string, n Newer) {
	newerMu.Lock()
	defer newerMu.Unlock()
	if n == nil {
		panic("store: Register Newer is nil")
	}
	if _, dup := backends[name]; dup {
		panic("store: Register called twice for Newer " + name)
	}
	backends[name] = n
}
