package main

import (
	"math/rand"
)

// randomNumber generates a random integer in the range [0, length).
// It takes an integer 'length' as input and returns a random integer.
func randomNumber(length int) int {
	randomNumber := rand.Intn(length)
	return randomNumber
}

// GenerateRandomNumber generates a list of random integers within a given range.
// It takes two parameters:
// - maxRange: the upper bound (exclusive) for the random numbers.
// - numberOfItemsRequired: the number of random integers to generate.
// It returns a slice of random integers.
func GenerateRandomNumber(maxRange int, numberOfItemsRequired int) []int {
	var randomNumberList []int

	for i := 0; i < numberOfItemsRequired; i++ {
		randomNumberList = append(randomNumberList, randomNumber(maxRange))
	}

	return randomNumberList
}
