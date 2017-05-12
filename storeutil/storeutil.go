package storeutil

import (
	"bytes"
	"encoding/json"
	"strings"
)

func AppendItem(result []byte, key string, value interface{}) (newresult []byte, err error) {
	var m map[string]interface{}
	err = json.Unmarshal(result, &m)
	if err != nil {
		return
	}
	m[key] = value
	newresult, err = json.Marshal(m)
	return
}

// A helper function convert string to the map
func Str2map(s string) (data map[string]interface{}, err error) {
	dec := json.NewDecoder(strings.NewReader(s))
	err = dec.Decode(&data)
	return
}

// A helper function convert string to the map
func Bytes2map(b []byte) (data map[string]interface{}, err error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(&data)
	return
}
