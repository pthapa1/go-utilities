package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := ".value"
	str2 := `getValueOf "foo" "bar"`

	res := strings.Split(str2, " ")
	var key string
	for i, value := range res {
		value = strings.ReplaceAll(value, `"`, "")
		value = strings.ReplaceAll(value, `'`, "")
		if res[0] == "getValueOf" && i == 1 {
			key = value
		}
	}
	fmt.Println(key)

	fmt.Println(string(str1[0]) == ".")
}

// if length is zero, return ""
// if length is not zero && it only has whitespace, return ""
// if the length is not zero, and !onlyHasEmptySpace, trim whitespace, and return content
