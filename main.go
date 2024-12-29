package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Years int
}

func main() {
	myDict := map[string]interface{}{
		"company.inc": "Test Company",
		"name":        "Pratik",
		"age":         32,
		"years":       111,
		"marathon":    false,
		"profession": map[string]interface{}{
			"company.info": "Earth Based Human Led",
			"title":        "Senior Test SE",
			"years":        5,
		},
		"myArr": []Person{
			{Name: "xaaha", Age: 22, Years: 11},
			{Name: "pt", Age: 35, Years: 88},
		}, "myArr2": []interface{}{
			Person{Name: "xaaha", Age: 22, Years: 11},
			Person{Name: "pt", Age: 35, Years: 88},
		},
	}

	result, err := GetValueOf("age", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'age':", result)
	}

	result, err = GetValueOf("marathon", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'marathon':", result)
	}

	result, err = GetValueOf("profession", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'profession':", result)
	}

	result, err = GetValueOf("myArr2", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'myArr2':", result)
	}

	result, err = GetValueOf("{company.inc}", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for '{company.inc}':", result)
	}

	result, err = GetValueOf("profession.{company.info}", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'company.info':", result)
	}

	result, err = GetValueOf("myArr[1]", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'myArr':", result)
	}

	result, err = GetValueOf("myArr[1].Name", myDict)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result for 'myArr.name':", result)
	}
}
