package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(length int) int  {
	randomNumber := rand.Intn(length)
	return randomNumber;
} 

func GenerateRandomNumber(length int) []int  {
	firstRandomNumber := randomNumber(length)
	
	fmt.Println(firstRandomNumber)

	secondRandomNumber := randomNumber(length)
	
	if firstRandomNumber == secondRandomNumber {
		secondRandomNumber = randomNumber(length)
	}
	fmt.Println(secondRandomNumber)

	var randomNumberList []int;

	randomNumberList = 	append(randomNumberList, firstRandomNumber)
	randomNumberList = 	append(randomNumberList, secondRandomNumber)

	return randomNumberList
}