package util

import (
	"encoding/json"
)

func ToJson(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
