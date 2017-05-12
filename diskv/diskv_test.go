package diskv

import (
	"bytes"
	"testing"

	"g.haodai.com/golang/common/store"
)

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
