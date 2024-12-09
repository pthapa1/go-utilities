package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CreateFilePath generates an absolute file path by joining the project's root directory
// with the provided relative filePath. Returns an error if the project root cannot be determined.
func CreateFilePath(filePath string) (string, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", err
	}
	finalFilePath := filepath.Join(projectRoot, filePath)
	return finalFilePath, nil
}

// listAllFiles searches for files matching the `matchFile` name (case-insensitive, .yaml/.yml only)
// in the specified directory and its subdirectories. If no directory is specified, it starts from the project root.
// Skips the `.git` folder during traversal. Returns an error if no matching files are found or if there are file system errors.
func ListMatchingFiles(matchFile string, initialPath ...string) ([]string, error) {
	matchFile = strings.ToLower(matchFile)
	var result []string

	initAbsFp, err := CreateFilePath("")
	if err != nil {
		return nil, fmt.Errorf("error getting initial file path: %w", err)
	}

	var startPath string
	if len(initialPath) == 0 {
		startPath = initAbsFp
	} else {
		startPath = initialPath[0]
	}

	dirContents, err := os.ReadDir(startPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", startPath, err)
	}

	filePattern := [2]string{".yaml", ".yml"}

	for _, val := range dirContents {
		// Skip hidden directories
		if val.IsDir() && strings.HasPrefix(val.Name(), ".") {
			continue
		}

		// Process files
		if !val.IsDir() {
			lowerName := strings.ToLower(val.Name())
			for _, ext := range filePattern {
				if strings.HasSuffix(lowerName, ext) {
					yamlFile := strings.TrimSuffix(lowerName, ext)
					if matchFile == yamlFile {
						matchingFp := filepath.Join(startPath, val.Name())
						result = append(result, matchingFp)
					}
				}
			}
		}

		// Process subdirectories
		if val.IsDir() {
			subDirPath := filepath.Join(startPath, val.Name())
			matches, err := ListMatchingFiles(matchFile, subDirPath)
			if err != nil && !isNoMatchingFileError(err) {
				fmt.Printf("Skipping subdirectory %s due to error: %v\n", val.Name(), err)
				continue
			}
			result = append(result, matches...)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"no files with matching name '%s' found in: %s", matchFile, path.Base(initAbsFp),
		)
	}

	return result, nil
}

// isNoMatchingFileError determines if the error is related to no matching files found.
func isNoMatchingFileError(err error) bool {
	return strings.Contains(err.Error(), "no files with matching name")
}

func main() {
	matches, err := ListMatchingFiles("test111")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Matching files:", matches)
	}
}
