package main

import (
	"fmt"
	"reflect"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var test Test
	test = Test{
		Name: "Pratik",
		Age:  31,
	}

	numberOfKeys := reflect.TypeOf(test).NumField()
	fmt.Println(numberOfKeys)

}

// make sure the numbers aren't repeated on the same array.
// make sure the numbers aren't repeated on the different array for up to 5 times.
// You can only request
