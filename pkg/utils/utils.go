package utils

import "encoding/json"

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ParseAsJSON[T any](body []byte, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil //nolint: nilnil
	}
	v := new(T)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
