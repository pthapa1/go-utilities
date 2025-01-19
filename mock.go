package main

// Define an interface `ValueProvider` with a single method `GetValue`
type ValueProvider interface {
	// This method signature defines a contract:
	// any type that implements `ValueProvider` must provide a `GetValue()` method.
	GetValue() string
}

// Implement the `ValueProvider` interface by creating a `DefaultProvider` struct
type DefaultProvider struct{}

// Implement the `GetValue()` method for `DefaultProvider`
// This method satisfies the `ValueProvider` interface and returns a simple string.
func (dp DefaultProvider) GetValue() string {
	return "hello there" // Returning a hardcoded string
}

// `setValue` function takes a `ValueProvider` interface as a parameter
// This function can accept any type that implements the `ValueProvider` interface.
func setValue(provider ValueProvider) string {
	userName := " xaaha"
	// Calling the `GetValue()` method of the `ValueProvider` interface,
	// which will invoke the `GetValue` implementation of the specific provider passed in.
	greetings := provider.GetValue() + userName // Concatenate the value from `GetValue()` and the username
	return greetings                            // Return the final greeting string
}
