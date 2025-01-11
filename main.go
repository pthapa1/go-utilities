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

func replaceInPlace(
	beforeMap map[string]interface{},
	parentKey string,
) []string {
	var cmprt []string
	for bKey, bValue := range beforeMap {
		currentKey := bKey

		// if the parentKey has something,
		if parentKey != "" {
			currentKey = parentKey + " -> " + bKey
		}

		switch bTypeVal := bValue.(type) {
		case string:
			// Example: Check if value contains a "." and modify accordingly
			if strings.Contains(bTypeVal, ".") {
				cmprt = append(cmprt, currentKey)
			}
		case map[string]interface{}:
			// cmprt = mergeMaps(cmprt, replaceInPlace(bTypeVal, currentKey))
			cmprt = append(cmprt, replaceInPlace(bTypeVal, currentKey)...)
		case []map[string]interface{}:
			for idx, val := range bTypeVal {
				key := fmt.Sprintf("%s[%d]", currentKey, idx)
				// cmprt = mergeMaps(cmprt, replaceInPlace(val, key))
				cmprt = append(cmprt, replaceInPlace(val, key)...)
			}
		default:
			fmt.Println("uncovered type")
			cmprt = append(cmprt, fmt.Sprintf("%s: %v", currentKey, bTypeVal))
		}
	}
	return cmprt
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
		"users": []map[string]interface{}{
			{"person": map[string]interface{}{
				"name": "Jane Doe",
				"age":  "{{.Age}}",
			}},
		},
	}

	result := replaceInPlace(beforeMap, "")

	prt, _ := marshalToJSON(result)
	fmt.Println(prt)
}

// func main() {
// 	beforeMap := map[string]interface{}{
// 		"foo":   "bar",
// 		"miles": "{{.distance}}",
// 		"age":   "29",
// 		"person": map[string]interface{}{
// 			"name": "Jane Doe",
// 			"age":  "{{.Age}}",
// 		},
// 	}
// 	afterMap := map[string]interface{}{
// 		"foo":   "bar",
// 		"miles": "26.2",
// 		"age":   "30",
// 		"person": map[string]interface{}{
// 			"name": "Jane Doe",
// 			"age":  "50",
// 		},
// 	}
// 	modifiedMap := replaceInPlace(beforeMap, afterMap)

// 	jsonString, _ := marshalToJSON(modifiedMap)
// 	fmt.Println(jsonString)
// }
