package utils

import (
	"encoding/json"
	"github.com/bytedance/sonic"
)

func JsonMarshal(v any) string {
	res, err := sonic.MarshalString(v)
	if err != nil {
		return ""
	}
	return res
}

func marshal(v any) string {
	bytes, err := json.MarshalIndent(v, " ", "")
	if err != nil {
		return ""
	}
	return string(bytes)
}
