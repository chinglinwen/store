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
	*store.Options
}

type Newer struct{}

func (*Newer) New(option *store.Options) (s store.Store, err error) {
	d := diskv.New(diskv.Options{
		BasePath:     option.BucketName,
		CacheSizeMax: 1024 * 1024,
	})
	return &D{d, option}, nil
}

// Write write any bytes to diskv.
func (d *D) Write(key string, value []byte) (err error) {
	if d.Compression != nil {
		value, err = d.Compression.Compress(value)
		if err != nil {
			return err
		}
	}
	return d.Disk.Write(key, value)
}

// Read read bytes from diskv.
func (d *D) Read(key string) ([]byte, error) {
	b, err := d.Disk.Read(key)
	if err != nil {
		return nil, err
	}
	if d.Compression != nil {
		return d.Compression.Decompress(b)
	}
	return b, nil
}

func init() {
	store.Register("diskv", &Newer{})
}
