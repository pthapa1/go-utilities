package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type CountryAndCode struct{
	Country string `json:"country"`
	Code string `json:"code"`
}

func CreateJsonFile()  {
	
	data, err := ioutil.ReadFile("address.txt");
	CheckError(err);

	content := string(data);

	lines := strings.Split(content, "\n")

	var countries []CountryAndCode;
	var countryCode CountryAndCode;

	for i :=0; i < len(lines); i++ {
		combinedText := lines[i];
		splitArray := strings.Split(combinedText, ": ");
		if len(splitArray) !=2 {
			continue
		}

		country := splitArray[0];
		code := splitArray[1];

		countryCode = CountryAndCode{
			Country: country,
			Code: code,
		}

		countries = append(countries, countryCode)
	}
	jsonData, err := json.Marshal(countries);
	CheckError(err)

 err =	ioutil.WriteFile("countries.json", jsonData, 0644);
	CheckError(err)


}