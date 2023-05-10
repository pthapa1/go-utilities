package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)
type CountryInAddress struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    Address      string `json:"address"`
    Phone        string `json:"phone"`
    CountryCode  string `json:"country-code"`
    Country      string `json:"country"`
}


func addCountryName(countryNamePath string)  {


	jsonFile, error :=  os.Open(countryNamePath);
	CheckError(error);

	defer jsonFile.Close()

	// get the country name. 

 filePathOfCountryName := path.Base(jsonFile.Name()) // returns the last part of name. users.json

	countryName := strings.TrimSuffix(filePathOfCountryName, path.Ext(jsonFile.Name())) // removes suffix or extension and returns only name. Here, users 

	// Read all the files. 

	data, err := ioutil.ReadAll(jsonFile);
	CheckError(err)

// unmarshall & loop
	var countryInAddress []CountryInAddress;

	err =  json.Unmarshal(data, &countryInAddress)
	CheckError(err)

	for i :=0; i < len(countryInAddress); i++ {
		countryInAddress[i].Country = countryName
	}
// Marshall
	modifiedData, err1 := json.MarshalIndent(countryInAddress, "", " ");	
	CheckError(err1);

// write 
	ioutil.WriteFile(countryNamePath, modifiedData, 0644)

	fmt.Println(countryNamePath, "modified successfully")

} 

func AddCountryToAllFiles(){
	
folderPath := "./addresses"

filelist, err :=  ioutil.ReadDir(folderPath)
CheckError(err)

for _, file := range filelist{

	if filepath.Ext(file.Name()) == ".json" {

		path := filepath.Join(folderPath, file.Name())
		addCountryName(path)
		fmt.Printf(file.Name(), "Successfully Modified")
	}
}
}