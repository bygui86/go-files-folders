package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	folderPath = "data/"
	fileName   = "write-test.txt"
	filePath   = folderPath + fileName
	text       = "some sample data\n"
	// mode       = 0666
	mode = 0755
)

func main() {
	checkAndCreateFolder()
	writeToFile()
	listFiles()
}

func checkAndCreateFolder() {
	_, statErr := os.Stat(folderPath)
	if os.IsNotExist(statErr) {
		mkErr := os.MkdirAll(folderPath, mode)
		if mkErr != nil {
			fmt.Printf("[ERROR] Folder %v creation failed: %v\n", folderPath, mkErr.Error())
		}
	}
}

func writeToFile() {
	// open file
	file, fileErr := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)
	if fileErr != nil {
		fmt.Printf("[ERROR] Open file %v failed: %v\n", filePath, fileErr.Error())
		return
	}

	// write to file
	_, writeErr := file.Write(
		[]byte(text),
	)
	if writeErr != nil {
		fmt.Printf("[ERROR] Write to file %v failed: %v\n", filePath, writeErr.Error())
	}

	// close file
	file.Close()
}

func listFiles() {
	listErr := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Printf("Found file %v in folder %v\n", info.Name(), folderPath)
		}
		return nil
	})
	if listErr != nil {
		fmt.Printf("[ERROR] List files in folder %v failed: %v\n", folderPath, listErr.Error())
	}
}
