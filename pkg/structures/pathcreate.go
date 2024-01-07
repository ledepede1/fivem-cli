package structures

import (
	"fmt"
	"os"
)

func createDir(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Printf("Error in creating project folder %s check if folder already exist\n", path)
	}
}

func createFile(path string) {
	if _, err := os.Create(path); err != nil {
		fmt.Println("Error in creating file")
	}
}

func writeFile(path string, text []byte) {
	if err := os.WriteFile(path, text, os.ModePerm); err != nil {
		fmt.Println("Error in writing file: ", err)
	}
}
