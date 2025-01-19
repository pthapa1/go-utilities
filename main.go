package main

import "fmt"

func main() {
	// Create a `DefaultProvider` instance, which implements `ValueProvider`
	provider := DefaultProvider{}
	// Pass the provider into the `setValue` function and print the result
	// The `setValue` function will use `DefaultProvider`'s `GetValue()` method,
	// which returns "hello there", and concatenate it with " xaaha".
	fmt.Println(setValue(provider)) // Expected output: "hello there xaaha"
}
