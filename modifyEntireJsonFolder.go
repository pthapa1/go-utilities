package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ModifyEntireJsonFolder(){
	
folderPath := "./addresses"

filelist, err :=  ioutil.ReadDir(folderPath)
CheckError(err)

for _, file := range filelist{

	if filepath.Ext(file.Name()) == ".json" {

		path := filepath.Join(folderPath, file.Name())
		ModifyJson(path)
		fmt.Printf(file.Name(), "Successfully Modified")
	}
}
}