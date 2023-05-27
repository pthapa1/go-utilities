package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CreateMap() {

	var countryMap map[string]string
	countryMap = make(map[string]string)

	countries, err := os.Open("./countries.json")
	CheckError(err)
	defer countries.Close()

	data, err := ioutil.ReadAll(countries)

	var country []CountryAndCode

	err1 := json.Unmarshal(data, &country)
	CheckError(err1)

	for i := 0; i < len(country); i++ {
		countryMap[country[i].Code] = country[i].Country
		// country[i].Code = country[i].Country
	}

	newData, err := json.Marshal(countryMap)
	CheckError(err)

	err2 := ioutil.WriteFile("countryMap.json", newData, 0644)
	CheckError(err2)
}
