package main

import (
	"errors"
	"fmt"
	"strconv"
)

func GetEnvVarGeneric(key string) (interface{}, error) {
	envVarMap := make(map[string]string)

	envVarMap["myBool"] = "tru"
	envVarMap["myId"] = "19393"
	envVarMap["myFloat"] = "19393.9229"
	envVarMap["myStr"] = "rubyOnTram"

	strValue, ok := envVarMap[key]
	if !ok {
		return nil, errors.New("Key " + key + "does not exist")
	}

	// Attempt to parse as bool
	if valueBool, err := strconv.ParseBool(strValue); err == nil {
		return valueBool, nil
	}

	// Attempt to parse as int
	if valueInt, err := strconv.Atoi(strValue); err == nil {
		return valueInt, nil
	}

	// Attempt to parse as float
	if valueFloat, err := strconv.ParseFloat(strValue, 64); err == nil {
		return valueFloat, nil
	}

	return strValue, nil
}

func main() {
	parsedVal, err := GetEnvVarGeneric("myBool")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("Type of parsedVal is %T\n", parsedVal)
	fmt.Println(parsedVal)
}
