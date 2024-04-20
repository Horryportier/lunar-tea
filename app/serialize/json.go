package serialize

import (
	"encoding/json"
)

type JsonMapFn func(T interface{}, m map[string]string) (map[string]string, error)

func JsonMap(T any, fn JsonMapFn) ([]byte, error) {
	_map := make(map[string]string)
	_map, err := fn(T, _map)
	if err != nil {
		return []byte{}, err
	}

	buf, err := json.MarshalIndent(_map, " ", "    ")
	return buf, err
}
