package main

import (
	"testing"
)

// Mock provider for testing
type MockProvider struct{}

func (mp MockProvider) GetValue() string {
	return "mocked hello"
}

func TestSetValue(t *testing.T) {
	mockProvider := MockProvider{}

	result := setValue(mockProvider)
	expected := "mocked hello xaaha"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
