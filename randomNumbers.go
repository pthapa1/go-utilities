package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(length int) int {
	randomNumber := rand.Intn(length)
	return randomNumber
}

func GenerateRandomNumber(maxRange int, numberOfItemsRequired int) []int {
	var randomNumberList []int

	for i := 0; i < numberOfItemsRequired; i++ {
		randomNumberList = append(randomNumberList, randomNumber(maxRange))
	}

	fmt.Printf("%v \n", randomNumberList)

	return randomNumberList
}

