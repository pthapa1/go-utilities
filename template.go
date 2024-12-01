package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
)

func replaceVariables(
	strToChange string,
	mapWithVars map[string]string,
) (string, error) {
	if len(strToChange) == 0 {
		return "", errors.New("input string is empty")
	}
	tmpl, err := template.New("template").Option("missingkey=error").Parse(strToChange)
	if err != nil {
		return "", fmt.Errorf("template parsing error: %w", err)
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, mapWithVars)
	if err != nil {
		return "", fmt.Errorf("%v", err) // Simplified error handling
	}
	return result.String(), nil
}

// Replace the template {{ }} in the variable map itself.
// Sometimes, we have a variables map that references some other variable in itself.
func prepareMap(varsMap map[string]string) (map[string]string, error) {
	for key, val := range varsMap {
		changedStr, err := replaceVariables(val, varsMap)
		if err != nil {
			return nil, fmt.Errorf("error while preparing variables in map: %v", err)
		}
		varsMap[key] = changedStr
	}
	return varsMap, nil
}

func SubstituteVariables(
	strToChange string,
	mapWithVars map[string]string,
) (string, error) {
	finalMap, err := prepareMap(mapWithVars)
	if err != nil {
		return "", err
	}

	result, err := replaceVariables(strToChange, finalMap)
	if err != nil {
		return "", err
	}
	return result, nil
}

// func main() {
// 	str := "Hello, {{.name}}! Welcome to {{.place}}. This is a {{.project}}. And a comment {{/* a comment */}} is ignored. Finally, an {{.pratik}}"
// 	vars := map[string]string{
// 		"name":    "John",
// 		"place":   "Gopherland",
// 		"project": "{{.hulak}}",
// 		"hulak":   "Hulak v1",
// 		"third":   "{{.name}}",
// 	}
//
// 	finalStr, err := SubstituteVariables(str, vars)
// 	if err != nil {
// 		fmt.Println("Error:", err.Error())
// 		return
// 	}
// 	fmt.Println(finalStr)
// }
