package main

import (
	"encoding/json"
	"fmt"
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

func replaceInPlace(beforeMap, afterMap map[string]interface{}) map[string]interface{} {
	// loop over the first one
	for bKey, bValue := range beforeMap {
		switch bTypeVal := bValue.(type) {
		case map[string]interface{}:
			replaceInPlace(bTypeVal, afterMap)
		case string:
			if strings.Contains(bTypeVal, ".") {
				afterMap[bKey] = 26.2
			}
		default:
			fmt.Println("uncovered type")
		}
	}
	return afterMap
}

func main() {
	beforeMap := map[string]interface{}{
		"foo":   "bar",
		"miles": "{{.distance}}",
		"age":   "29",
		"person": map[string]interface{}{
			"name": "Jane Doe",
			"age":  "{{.Age}}",
		},
	}
	afterMap := map[string]interface{}{
		"foo":   "bar",
		"miles": "26.2",
		"age":   "30",
		"person": map[string]interface{}{
			"name": "Jane Doe",
			"age":  "50",
		},
	}

	modifiedMap := replaceInPlace(beforeMap, afterMap)

	jsonString, _ := marshalToJSON(modifiedMap)
	fmt.Println(jsonString)
}
