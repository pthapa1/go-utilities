package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// add error saying that we don't have file with matching name if the result is empty
// add the function description here
// and make it efficient
func createFilePath(filePath string) (string, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", err
	}
	finalFilePath := filepath.Join(projectRoot, filePath)
	return finalFilePath, nil
}

func listAllFiles(matchFile string, path ...string) ([]string, error) {
	matchFile = strings.ToLower(matchFile)
	var result []string
	initAbsFp, err := createFilePath("")
	if err != nil {
		fmt.Println(err)
		return []string{}, err
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
		return []string{}, err
	}

	for _, val := range dirContents {
		filePattern := [2]string{".yaml", ".yml"}

		// when it's a file
		if !val.IsDir() {
			// ensure that the file matches the pattern above
			for _, ptrn := range filePattern {
				if strings.Contains(val.Name(), ptrn) {
					yamlFile := strings.Split(val.Name(), ptrn)[0]
					matchFile = strings.Split(matchFile, ptrn)[0]
					if matchFile == yamlFile {
						matchingFp := filepath.Join(startPath, val.Name())
						result = append(result, matchingFp)
					}
				}
			}
		}

		// recurse into subdirectories
		if val.IsDir() {
			subDirPath := filepath.Join(startPath, val.Name())  // Use `startPath`
			matches, err := listAllFiles(matchFile, subDirPath) // Recurse with new path
			result = append(result, matches...)
			if err != nil {
				fmt.Println(err)
				return []string{}, err
			}
		}
	}
	return result, nil
}

func main() {
	files, err := listAllFiles("test.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)
}
