// Store backend for oss ( Aliyun Object Storage Service)
package oss

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint string
	apikey   string
	secret   string
)

// SetClient will set the default client
// A helper function for initialization
func SetKeySecret(e, k, sec string) {
	endpoint = e
	apikey = k
	secret = sec
}

type S struct {
	BucketName string
	Bucket     *oss.Bucket
}

// Provide the client for convience
func GetClient() (*oss.Client, error) {
	return oss.New(endpoint, apikey, secret)
}

func New(bucket string) (s S, err error) {
	c, err := GetClient()
	if err != nil {
		return
	}
	b, err := c.Bucket(bucket)
	if err != nil {
		return
	}
	return S{bucket, b}, nil
}

// Write write any bytes to oss
func (s *S) Write(key string, value []byte) error {
	return s.Bucket.PutObject(key, bytes.NewReader(value))
}

func (s *S) WriteString(key, value string) error {
	return s.Bucket.PutObject(key, strings.NewReader(value))
}

//func (s *S) WriteR(key string,r io.Reader) error {
//	return s.bucket.PutObject(key, r)
//}

// Read read bytes from oss
func (s *S) Read(key string) ([]byte, error) {
	body, err := s.Bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	return ioutil.ReadAll(body)
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
