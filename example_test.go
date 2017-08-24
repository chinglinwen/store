package store_test

import (
	"fmt"

	"hdgit.com/golib/store"
	oss "hdgit.com/golib/store/oss"
)

func Example() {
	// import (
	//	"hdgit.com/golib/store"
	//	_ "hdgit.com/golib/store/oss"
	// )

	// Can be no compress by adding a option here.
	s, err := store.New("oss", bucket)
	if err != nil {
		fmt.Print("new store err:", err)
		return
	}

	// Folder need to create at the console first
	// Any path will be okay, the path must exist
	k, v := "folder1/folder2/hello", []byte("test")
	if err := s.Write(k, v); err != nil {
		fmt.Print("write err:", err)
		return
	}

	b, err := s.Read(k)
	if err != nil {
		fmt.Print("read err:", err)
		return
	}
	fmt.Print("everything ok, result:", string(b))

	// Output:
	// everything ok, result:test
}

var bucket string

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
