package main

import "fmt"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello World")
}
