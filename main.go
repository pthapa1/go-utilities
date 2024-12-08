package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFilePath(filePath string) (string, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", err
	}
	finalFilePath := filepath.Join(projectRoot, filePath)
	return finalFilePath, nil
}

// a user would provide me with a name -f test
// in the current working dir "./" list all the file paths

func listAllFiles(path ...string) error {
	initAbsFp, err := CreateFilePath("")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var startPath string
	if len(path) == 0 {
		startPath = initAbsFp
	} else {
		startPath = path[0]
	}

	dirContents, err := os.ReadDir(startPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, val := range dirContents {
		filePattern := [2]string{".yaml", ".yml"}

		// when it's a file
		if !val.IsDir() {
			// ensure that the file matches the pattern above
			for _, ptrn := range filePattern {
				if strings.Contains(val.Name(), ptrn) {
					yamlFiles := strings.Split(val.Name(), ptrn)[0]
					fmt.Println("Split", yamlFiles)
				}
			}
		}

		// recurse into subdirectories
		if val.IsDir() {
			subDirPath := filepath.Join(startPath, val.Name()) // Use `startPath`
			err := listAllFiles(subDirPath)                    // Recurse with new path
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}

func main() {
	err := listAllFiles()
	if err != nil {
		fmt.Println(err)
	}
}
