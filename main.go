package main

func CheckError(e error)  {
	if e != nil {
		panic(e)
	}
}

func main(){

	AddCountryToAllFiles()
	
}