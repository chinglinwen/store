package store_test

import (
	"bytes"
	"testing"

	"g.haodai.com/golang/common/store"
	_ "g.haodai.com/golang/common/store/diskv"
	_ "g.haodai.com/golang/common/store/oss"
)

func Test(t *testing.T) {
	tests := []struct {
		backend string
		bucket  string
	}{
		{"oss", "bigprove-dev"},
		{"diskv", "bigprove-dev"},
	}
	for _, test := range tests {
		s, err := store.New(test.backend, test.bucket)
		if err != nil {
			t.Errorf("new store err: %v", err)
		}
		k, v := "hello.txt", []byte("test")
		if err := s.Write(k, v); err != nil {
			t.Errorf("write err: %v", err)
		}
		b, err := s.Read(k)
		if err != nil {
			t.Errorf("read err: %v", err)
		}
		if !bytes.Equal(b, v) {
			t.Errorf("read and write not equal err: %v", err)
		}
	}
	t.Log("everything ok")
}

func TestNoDecompress(t *testing.T) {
	s, err := store.New("oss", "bigprove-dev", store.NoDecompress)
	if err != nil {
		t.Errorf("new store err: %v", err)
	}
	k, v := "hello.txt", []byte("test")
	if err := s.Write(k, v); err != nil {
		t.Errorf("write err: %v", err)
	}
	b1, err := s.Read(k)
	if err != nil {
		t.Errorf("read err: %v", err)
	}
	b, err := s.C.Decompress(b1)
	if err != nil {
		t.Errorf("decompress err: %v", err)
	}
	if !bytes.Equal(b, v) {
		t.Errorf("read and write not equal err: %v", err)
	}
}
