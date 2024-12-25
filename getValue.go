package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GetValue(myKey string, dict map[string]interface{}) interface{} {
	var result interface{}

	for key, value := range dict {
		if key == myKey {
			result = value
			if va, ok := value.(map[string]interface{}); ok {
				jsonString, err := json.Marshal(va)
				if err != nil {
					return fmt.Sprintf("Error converting to JSON: %v", err)
				}
				return string(jsonString)
			}
			return result
		}

		// Recursively check nested maps
		if val, ok := value.(map[string]interface{}); ok {
			innerMap := GetValue(myKey, val)
			if va, ok := innerMap.(map[string]interface{}); ok {
				jsonString, err := json.Marshal(va)
				if err != nil {
					return fmt.Sprintf("Error converting to JSON: %v", err)
				}
				return string(jsonString)
			}
			if innerMap != nil {
				return innerMap
			}
		}
	}
	return result
}

func LookUpValuePath(key string, data map[string]interface{}) (interface{}, error) {
	if value, exists := data[key]; exists {
		return value, nil
	}

	pathSeparator := "." // You can change this to "/"
	segments := strings.Split(key, pathSeparator)
	current := data

	for i, segment := range segments {
		value, exists := current[segment]
		if !exists {
			return nil, errors.New("key not found: " + strings.Join(segments[:i+1], pathSeparator))
		}

		if i == len(segments)-1 {
			return value, nil
		}

		if nestedMap, ok := value.(map[string]interface{}); ok {
			current = nestedMap
		} else {
			return nil, errors.New("invalid path, segment is not a map: " + strings.Join(segments[:i+1], pathSeparator))
		}
	}

	return nil, errors.New("unexpected error")
}

type Person struct {
	Name  string
	Age   int
	Years int
}

// return json
// stringfy the response if it's json
// skip the .
// use of [0] should be accepted as well for arrays
