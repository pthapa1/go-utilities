package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func marshalToJSON(value interface{}) (interface{}, error) {
	switch val := value.(type) {
	case string, bool, int, float64:
		return val, nil
	case nil:
		return nil, nil
	default:
		if arr, ok := value.([]interface{}); ok {
			var jsonArray []string
			for _, item := range arr {
				jsonStr, err := json.Marshal(item)
				if err != nil {
					return "", err
				}
				jsonArray = append(jsonArray, string(jsonStr))
			}
			return fmt.Sprintf("[%s]", strings.Join(jsonArray, ",")), nil
		}
		jsonStr, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(jsonStr), nil
	}
}

func ParseArrayKey(segment string) (bool, string, int) {
	if strings.HasSuffix(segment, "]") && strings.Contains(segment, "[") {
		openBracket := strings.LastIndex(segment, "[")
		closeBracket := strings.LastIndex(segment, "]")
		indexStr := segment[openBracket+1 : closeBracket]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return false, segment, -1 // Invalid index
		}
		return true, segment[:openBracket], index
	}
	return false, segment, -1
}

func parsePath(path string) ([]interface{}, error) {
	var keys []interface{}

	if len(path) == 0 {
		return keys, fmt.Errorf("path should not be empty")
	}

	rawKeys := strings.Split(path, "->")
	for _, segment := range rawKeys {
		trimmedKey := strings.TrimSpace(segment)
		isArrayKey, keyPart, index := ParseArrayKey(trimmedKey)
		if isArrayKey {
			keys = append(keys, keyPart)
			keys = append(keys, index)
		} else {
			keys = append(keys, trimmedKey)
		}
	}
	return keys, nil
}

func main() {
	keys, err := parsePath("employer[0] -> id")
	if err != nil {
		fmt.Println(err)
	}
	prt, _ := marshalToJSON(keys)
	fmt.Println(prt)
}
