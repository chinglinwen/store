// Store is for storing arbitrary data.
// It support oss now
package store

type Reader interface {
	Read(key string) ([]byte, error)
}

type Writer interface {
	Write(key string, value []byte) error
}

type ReadWriter interface {
	Reader
	Writer
}

type Store struct {
	ReadWriter
}

// Create a new store for read and write
func New(rw ReadWriter) Store {
	return Store{rw}
}
