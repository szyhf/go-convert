package convert

import (
	"encoding/json"
)

func MustJson(i interface{}) []byte {
	if d, err := json.Marshal(i); err == nil {
		return d
	} else {
		panic(err)
	}
}

func MustJsonString(i interface{}) string {
	return string(MustJson(i))
}
