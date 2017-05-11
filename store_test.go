package store_test

import (
	"fmt"
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
		err := readwrite(test.backend, test.bucket)
		if err != nil {
			fmt.Print(err)
		}
	}
	fmt.Print("everything ok")
}

func readwrite(backend, bucket string) error {
	s, err := store.New(backend, bucket)
	if err != nil {
		return fmt.Errorf("new store err: %v", err)
	}

	// Folder need to create at the console first
	// Any path will be okay, the path must exist
	k, v := "hello", []byte("test")
	if err := s.Write(k, v); err != nil {
		return fmt.Errorf("write err: %v", err)
	}

	_, err = s.Read(k)
	if err != nil {
		return fmt.Errorf("read err: %v", err)
	}
	return nil
}
