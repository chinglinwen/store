// Store backend for oss (Aliyun Object Storage Service).
// It comes with default key and secret (it can be init again).
//
// Using this as the backend of store package,
// If for registering backend only (it can import as blank identifier).
package oss

import (
	"bytes"
	"io/ioutil"

	"g.haodai.com/golang/common/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint string
	apikey   string
	secret   string
)

// SetClient will set the default client
// A helper function for initialization.
func SetKeySecret(e, k, sec string) {
	endpoint = e
	apikey = k
	secret = sec
}

type S struct {
	Bucket *oss.Bucket
	*store.Options
}

// Provide the client for convience.
func GetClient() (*oss.Client, error) {
	return oss.New(endpoint, apikey, secret)
}

type Newer struct{}

func (*Newer) New(option *store.Options) (s store.Store, err error) {
	c, err := GetClient()
	if err != nil {
		return
	}
	b, err := c.Bucket(option.BucketName)
	if err != nil {
		return
	}
	return &S{b, option}, nil
}

// Write write any bytes to oss.
func (s *S) Write(key string, value []byte) (err error) {
	if s.Compression != nil {
		value, err = s.Compression.Compress(value)
		if err != nil {
			return err
		}
	}
	return s.Bucket.PutObject(key, bytes.NewReader(value))
}

// Read read bytes from oss.
func (s *S) Read(key string) ([]byte, error) {
	body, err := s.Bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	if s.Compression != nil {
		return s.Compression.Decompress(b)
	}
	return b, err
}

func init() {
	store.Register("oss", &Newer{})
}

func init() {
	// Set a default for convenience.
	// It can be set again in somewhere else too.
	endpoint := "http://oss-cn-zhangjiakou.aliyuncs.com"
	key := "LTAIj8XauZDqhzLz"
	secret := "0bvzEIzPktdVVmIVIGIeylGhUCLxil"
	SetKeySecret(endpoint, key, secret)
}
