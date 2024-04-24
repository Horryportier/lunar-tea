package serialize

import (
	"encoding/json"
)

type JsonMapFn func(T interface{}, m map[string]string) (map[string]string, error)

func JsonMap(T any, type_id string, fn JsonMapFn) ([]byte, error) {
	m := make(map[string]string)
	m, err := fn(T, m)
	if err != nil {
		return []byte{}, err
	}
	m["type"] = type_id
	buf, err := json.Marshal(m)
	return buf, err
}
