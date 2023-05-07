package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CompleteAddress struct {
	Address
	CountryCode    string `json:"country-code"`

}

func addCountryCode(countryNamePath string, alpha2 string)  {	

 jsonFile, err  := os.Open(countryNamePath);
	CheckError(err)

	fmt.Println(" File successfully opened");

	//defer the closing. 
 defer jsonFile.Close();

	data, err1 := ioutil.ReadAll(jsonFile);
	CheckError(err1);

	var completeAddress []CompleteAddress;

	err2 := json.Unmarshal(data, &completeAddress)
	CheckError(err2)

	for i := 0; i < len(completeAddress); i++ {
		completeAddress[i].CountryCode = alpha2;
	}

	modifiedData, err3 := json.MarshalIndent(completeAddress, "", " ")
	CheckError(err3);

 err4 := ioutil.WriteFile(countryNamePath, modifiedData, 0644)
	CheckError(err4)

	fmt.Println(countryNamePath, "modified successfully.")
}

func ModifyEntireJsonFolderToAddCountryCode()  {

	countryCode := []string{"MK", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PL", "PT", "QA", "RO", "RU", "RW", "KN", "LC", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SK", "SI", "SB", "SO", "ZA", "KR", "SS", "ES", "LK", "SD", "SR", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TO", "TT", "TN", "TR", "TM", "TV", "UG", "UA", "AE", "GB", "US", "UY", "UZ", "VU", "VA", "VE", "VN", "YE", "ZM", "ZW"}


	counter := 0;

	files, err := ioutil.ReadDir("./addresses")
	CheckError(err)

	for _, file := range files{

		if filepath.Ext(file.Name()) == ".json" {

			countryNamePath := filepath.Join(".", "addresses", os.FileInfo.Name(file))
			alpha2 := countryCode[counter]
			addCountryCode(countryNamePath, alpha2)
			counter++
		}
	}
}
