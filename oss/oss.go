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
	*oss.Bucket
}

// Provide the client for convience.
func GetClient() (*oss.Client, error) {
	return oss.New(endpoint, apikey, secret)
}

type newer struct{}

func (*newer) New(bucket string) (s store.Backend, err error) {
	c, err := GetClient()
	if err != nil {
		return
	}
	b, err := c.Bucket(bucket)
	if err != nil {
		return
	}
	return &S{b}, nil
}

// Write write any bytes to oss.
func (s *S) Write(key string, value []byte) (err error) {
	return s.PutObject(key, bytes.NewReader(value))
}

// Read read bytes from oss.
func (s *S) Read(key string) ([]byte, error) {
	body, err := s.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	return ioutil.ReadAll(body)
}

func init() {
	store.Register("oss", &newer{})
}

func init() {
	// Set a default for convenience.
	// It can be set again in somewhere else too.
	endpoint := "http://oss-cn-zhangjiakou.aliyuncs.com"
	key := "LTAIiJAOgdjxTITy"
	secret := "KMHRfsN7nyC74HGBvp2M5xSSjuVhym"
	SetKeySecret(endpoint, key, secret)
}
