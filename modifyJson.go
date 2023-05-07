package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


type Address struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Address string `josn:"address"`
	Phone string `json:"phone"`
}


func ModifyJson(path string){
	
// open file
	jsonFile, err := os.Open(path)
	CheckError(err)

	fmt.Println("Successfully Opened user.json")

	// defer the closing of jsonFile so we can parse it later
	defer jsonFile.Close()

	// Read the file 
	data, err := ioutil.ReadAll(jsonFile);
	CheckError(err)

	// declare the array that will carry the changes. 
	var address []Address;

	// serialized data to go structure
	err = json.Unmarshal(data, &address)
	CheckError(err)

	// for all addresses. add the id value
	for i := 0; i < len(address); i++ {
		address[i].ID = i + 1
	}

	// Marshal the modified data back to JSON format
	modifiedData, err := json.MarshalIndent(address, "", " ")
	CheckError(err)

	// write the modified json data back to the file. 

	err = ioutil.WriteFile(path, modifiedData, 0644)
	CheckError(err)

fmt.Println("File modified successfully")
	

}
