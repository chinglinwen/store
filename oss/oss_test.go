package oss

import (
	"bytes"
	"testing"

	"hdgit.com/golib/store/oss"
)

var bucket string

func TestPutGet(t *testing.T) {
	n := &newer{}
	s, err := n.New(bucket)
	if err != nil {
		t.Error("new client err:", err)
	}
	k, v := "hello", []byte("test")
	err = s.Write(k, v)
	if err != nil {
		t.Error("write err:", err)
	}
	b, err := s.Read(k)
	if err != nil {
		t.Error("read err:", err)
	}
	if !bytes.Equal(b, v) {
		t.Errorf("read error, got %v, want %v\n", string(b), string(v))
	}
}

//dev
func init() {
	// Set a default for convenience.
	// It can be set again in somewhere else too.
	endpoint := "http://oss-cn-beijing.aliyuncs.com"
	key := "xxx"
	secret := "xxx"
	oss.SetKeySecret(endpoint, key, secret)

	bucket = "xxx"
}
