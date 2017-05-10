package store_test

import (
	"bytes"
	"fmt"

	"g.haodai.com/golang/common/store"
	"g.haodai.com/golang/common/store/oss"
)

func init() {
	endpoint := "http://oss-cn-zhangjiakou.aliyuncs.com"
	key := "LTAIj8XauZDqhzLz"
	secret := "0bvzEIzPktdVVmIVIGIeylGhUCLxil"
	oss.SetKeySecret(endpoint, key, secret)
}

func Example() {
	// Init oss key and secret first
	// oss.SetKeySecret(endpoint, key, secret)

	ss, err := oss.New("bigprove-dev")
	if err != nil {
		fmt.Print("new client err:", err)
		return
	}
	s := store.New(&ss)

	// Folder need to create at the console first
	// Any path will be okay, the path must exist
	k, v := "folder1/folder2/hello", []byte("test")
	err = s.Write(k, v)
	if err != nil {
		fmt.Print("write err:", err)
		return
	}

	b, err := s.Read(k)
	if err != nil {
		fmt.Print("read err:", err)
		return
	}
	if !bytes.Equal(b, v) {
		fmt.Errorf("read error, got %v, want %v\n", string(b), string(v))
		return
	}
	fmt.Print("everything ok, result:", string(b))

	// Output:
	// everything ok, result:test
}
