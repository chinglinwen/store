// Store backend for diskv.
//
// Using this as the backend of store package,
// If for registering backend only (it can import as blank identifier).
package diskv

import (
	"g.haodai.com/golang/common/store"
	"github.com/peterbourgon/diskv"
)

type D struct {
	Disk *diskv.Diskv
}

type Newer struct{}

func (*Newer) New(path string) (s store.Store, err error) {
	d := diskv.New(diskv.Options{
		BasePath:     path,
		CacheSizeMax: 1024 * 1024,
	})
	return &D{d}, nil
}

// Write write any bytes to diskv.
func (d *D) Write(key string, value []byte) error {
	return d.Disk.Write(key, value)
}

func (d *D) WriteString(key, value string) error {
	return d.Disk.Write(key, []byte(value))
}

// Read read bytes from diskv.
func (d *D) Read(key string) ([]byte, error) {
	return d.Disk.Read(key)
}

func init() {
	store.Register("diskv", &Newer{})
}
