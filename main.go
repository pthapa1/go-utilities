package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createFilePath(filePath string) (string, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", err
	}
	finalFilePath := filepath.Join(projectRoot, filePath)
	return finalFilePath, nil
}

// a user would provide me with a name -f test
// in the current working dir "./" list all the file paths

func listAllFiles(matchFile string, path ...string) ([]string, error) {
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
					result = append(result, yamlFile)
					// if matchFile == yamlFile {
					// }
					// lowercase it
					// remove yaml and yml
					// find it's path
					// and return the relative path in result
					// error saying that we don't have file with matching name error
					fmt.Println("Split", yamlFile)
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
