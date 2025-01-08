package main

import (
	"fmt"
	"strings"
)

// func marshalToJSON(value interface{}) (interface{}, error) {
// 	switch val := value.(type) {
// 	case string, bool, int, float64:
// 		return val, nil
// 	case nil:
// 		return nil, nil
// 	default:
// 		if arr, ok := value.([]interface{}); ok {
// 			var jsonArray []string
// 			for _, item := range arr {
// 				jsonStr, err := json.Marshal(item)
// 				if err != nil {
// 					return "", err
// 				}
// 				jsonArray = append(jsonArray, string(jsonStr))
// 			}
// 			return fmt.Sprintf("[%s]", strings.Join(jsonArray, ",")), nil
// 		}
// 		jsonStr, err := json.Marshal(value)
// 		if err != nil {
// 			return "", err
// 		}
// 		return string(jsonStr), nil
// 	}
// }

func replaceInPlace(
	beforeMap map[string]interface{},
	afterMap map[string]interface{},
	parentKey string,
) map[string]interface{} {
	cmprt := make(map[string]interface{})
	for bKey, bValue := range beforeMap {
		// Construct the current path key
		currentKey := bKey
		if parentKey != "" {
			currentKey = parentKey + " -> " + bKey
		}

		switch bTypeVal := bValue.(type) {
		case string:
			// Example: Check if value contains a "." and modify accordingly
			if strings.Contains(bTypeVal, ".") {
				cmprt[currentKey] = "modified_value"
			} else {
				cmprt[currentKey] = bValue
			}
		case map[string]interface{}:
			// Recurse into nested maps
			cmprt = mergeMaps(cmprt, replaceInPlace(bTypeVal, afterMap, currentKey))
		default:
			fmt.Println("uncovered type")
			cmprt[currentKey] = bValue
		}
	}
	return cmprt
}

// Helper function to merge two maps
func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
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

	afterMap := map[string]interface{}{} // Unused but included for completeness
	result := replaceInPlace(beforeMap, afterMap, "")
	for k, v := range result {
		fmt.Printf("%s: %v\n", k, v)
	}
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
