package main

import "fmt"

func main() {
	myDict := map[string]interface{}{
		"name":  "Pratik",
		"age":   32,
		"years": 111,
		"profession": map[string]interface{}{
			"title": "Senior Test SE",
			"years": 5,
		},
		"myArr": []Person{
			{Name: "xaaha", Age: 22, Years: 11},
			{Name: "pt", Age: 35, Years: 88},
		},
	}
	result, err := LookUpValuePath("myArr", myDict)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
