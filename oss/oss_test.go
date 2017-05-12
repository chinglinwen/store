package oss

import (
	"bytes"
	"testing"
)

func TestPutGet(t *testing.T) {
	n := &newer{}
	s, err := n.New("bigprove-dev")
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
