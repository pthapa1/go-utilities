package main

import "fmt"

type Utility struct{}

func (u *Utility) greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	var u Utility
	message := u.greet("Pratik")
	fmt.Println(message)
}
