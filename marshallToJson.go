package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func MarshalToJson(value interface{}) (interface{}, error) {
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
