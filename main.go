package main

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	GenerateRandomNumber(10, 3)

}

// make sure the numbers aren't repeated on the same array.
// make sure the numbers aren't repeated on the different array for up to 5 times.
// You can only request
