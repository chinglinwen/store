// Store is for storing arbitrary data.
// It support oss as backend for now.
package store

import (
	"encoding/json"
	"sync"
)

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
	c := NewGzipCompression()
	return backends[backend].New(&Options{bucket, c})
}

type Newer interface {
	New(*Options) (Store, error)
}

type Options struct {
	BucketName  string
	Compression Compression
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

func AppendItem(result []byte, key string, value interface{}) (newresult []byte, err error) {
	var m map[string]interface{}
	err = json.Unmarshal(result, &m)
	if err != nil {
		return
	}
	m[key] = value
	newresult, err = json.Marshal(m)
	return
}
