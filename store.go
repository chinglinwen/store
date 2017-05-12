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

type Backend interface {
	Reader
	Writer
}

type Store struct {
	B Backend
	C Compression
}

func (s *Store) Write(key string, value []byte) (err error) {
	if s.C != nil {
		value, err = s.C.Compress(value)
		if err != nil {
			return err
		}
	}
	return s.B.Write(key, value)
}

func (s *Store) Read(key string) ([]byte, error) {
	v, err := s.B.Read(key)
	if err != nil {
		return nil, err
	}
	if s.C != nil {
		return s.C.Decompress(v)
	}
	return v, err
}

// Create a new store for read and write,
// Backend is one of registered backends.
func New(backend, bucket string) (*Store, error) {
	c := NewGzipCompression()
	s, err := backends[backend].New(bucket)
	if err != nil {
		return nil, err
	}
	return &Store{s, c}, nil
}

type Newer interface {
	New(string) (Backend, error)
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
