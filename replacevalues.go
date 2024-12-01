package main

import "fmt"

func ReplaceValues(dict map[string]interface{}) map[string]interface{} {
	// Initialize a new map to store modified values
	changedMap := make(map[string]interface{})

	// Iterate over each key-value pair in the input map
	for key, val := range dict {
		switch valTyped := val.(type) {
		// If the value is another map, recursively call ReplaceValues
		case map[string]interface{}:
			changedMap[key] = ReplaceValues(valTyped)
		// If the value is a string, print and add it to the changed map
		case string:
			fmt.Println("This is a string:", valTyped)
			changedMap[key] = valTyped // Add to changedMap
		// If the value is a map of strings, iterate over it and print values
		case map[string]string:
			innerMap := make(map[string]interface{})
			for k, v := range valTyped {
				fmt.Println("This is a string within map[string]string:", v)
				innerMap[k] = v // Add string values from map[string]string to a new map
			}
			changedMap[key] = innerMap
		// Default case for unexpected types
		default:
			fmt.Println("Unexpected type:", val)
			changedMap[key] = val // Add unmodified
		}
	}
	return changedMap
}
