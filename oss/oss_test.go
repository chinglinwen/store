package oss

import (
	"bytes"
	"os"
	"testing"

	"g.haodai.com/golang/common/store"
)

func TestMain(m *testing.M) {
	endpoint := "http://oss-cn-zhangjiakou.aliyuncs.com"
	key := "LTAIj8XauZDqhzLz"
	secret := "0bvzEIzPktdVVmIVIGIeylGhUCLxil"
	SetKeySecret(endpoint, key, secret)

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestPutGet(t *testing.T) {
	n := &Newer{}
	s, err := n.New(&store.Options{"bigprove-dev", store.NewGzipCompression()})
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
