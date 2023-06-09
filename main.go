package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}

	directory := os.Args[1]

	rawDir := filepath.Join(directory, "raw")

	rawFiles, err := filepath.Glob(filepath.Join(rawDir, "*.RAF"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawFilesNum := len(rawFiles)
	deletedNum := 0
	for _, rawFile := range rawFiles {
		jpgFile := strings.Replace(rawFile, rawDir, directory, -1)
		jpgFile = strings.Replace(jpgFile, ".RAF", ".JPG", -1)

		// If JPG file exists, skip
		if _, err := os.Stat(jpgFile); err == nil {
			continue
		}

		// If JPG file does not exist, delete RAW file
		if err := os.Remove(rawFile); err != nil {
			fmt.Printf("Failed to delete %s: %s\n", rawFile, err)
			continue
		}

		deletedNum++
		fmt.Printf("Deleted %s\n", rawFile)
	}

	fmt.Printf("Delete %d files done. Remaining %d files\n", deletedNum, rawFilesNum-deletedNum)
}
