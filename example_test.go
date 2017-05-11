package store_test

import (
	"fmt"

	"g.haodai.com/golang/common/store"
	_ "g.haodai.com/golang/common/store/oss"
)

func Example() {
	// import (
	//	"g.haodai.com/golang/common/store"
	//	_ "g.haodai.com/golang/common/store/oss"
	// )

	s, err := store.New("oss", "bigprove-dev")
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
