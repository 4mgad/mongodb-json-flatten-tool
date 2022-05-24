package util

import (
	"encoding/json"
)

func traverse(jsonIn *map[string]interface{}, jsonOut *map[string]interface{}, parentPath string) {
	keyPath := ""
	if parentPath != "" {
		keyPath = parentPath + "."
	}
	for k, v := range *jsonIn {
		switch v.(type) {
		case map[string]interface{}:
			val := v.(map[string]interface{})
			traverse(&val, jsonOut, keyPath+k)
		default:
			(*jsonOut)[keyPath+k] = v
		}
	}
}

func Flatten(jsonInStr string) string {
	jsonIn := map[string]interface{}{}
	json.Unmarshal([]byte(jsonInStr), &jsonIn)

	jsonOut := map[string]interface{}{}
	traverse(&jsonIn, &jsonOut, "")
	jsonOutStr, _ := json.MarshalIndent(&jsonOut, "", "  ")
	return string(jsonOutStr)
}
