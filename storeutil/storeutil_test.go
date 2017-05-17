package storeutil

import (
	"encoding/json"
	"testing"
)

var data = map[string]interface{}{
	"hello":  1,
	"hellob": 2,
}

func TestBytes2map(t *testing.T) {
	b, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	d, err := Bytes2map(b)
	if err != nil {
		t.Error(err)
	}
	t.Log(d)
}
