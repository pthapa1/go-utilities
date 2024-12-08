package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// createFilePath generates an absolute file path by joining the project's root directory
// with the provided relative filePath. Returns an error if the project root cannot be determined.
func createFilePath(filePath string) (string, error) {
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
func listAllFiles(matchFile string, initialPath ...string) ([]string, error) {
	matchFile = strings.ToLower(matchFile)
	var result []string

	// Get the initial path to start the search
	initAbsFp, err := createFilePath("")
	if err != nil {
		return nil, fmt.Errorf("error getting initial file path: %w", err)
	}

	var startPath string
	if len(initialPath) == 0 {
		startPath = initAbsFp
	} else {
		startPath = initialPath[0]
	}

	// Read the contents of the starting directory
	dirContents, err := os.ReadDir(startPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", startPath, err)
	}

	filePattern := [2]string{".yaml", ".yml"}

	for _, val := range dirContents {
		// Skip `.git` or `.vscode` directory
		if val.IsDir() && strings.HasPrefix(val.Name(), ".") {
			continue
		}

		// Process files
		if !val.IsDir() {
			lowerName := strings.ToLower(
				val.Name(),
			) // Convert name to lowercase for consistent comparison
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

		// Recurse into subdirectories, skipping `.git`
		if val.IsDir() {
			subDirPath := filepath.Join(startPath, val.Name())
			matches, err := listAllFiles(matchFile, subDirPath)
			if err != nil {
				dir := path.Base(subDirPath)
				return nil, fmt.Errorf("\n error processing subdirectory %s: %w", dir, err)
			}
			result = append(result, matches...)
		}
	}

	// Return an error if no matching files are found
	if len(result) == 0 {
		return nil, fmt.Errorf(
			"\n no files with matching name '%s' found in path '%s'",
			matchFile,
			initAbsFp,
		)
	}

	return result, nil
}

func main() {
	matches, err := listAllFiles("test")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Matching files:", matches)
	}
}
