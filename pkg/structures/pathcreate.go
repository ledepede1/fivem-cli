package structures

import (
	"fmt"
	"os"
)

var Worked = true

func createDir(path string) {
	if Worked {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			Worked = false
			fmt.Printf("Error in creating project folder %s check if folder already exist\n", path)
		} else {
			Worked = true
		}
	}
}

func createFile(path string) {
	if Worked {
		if _, err := os.Create(path); err != nil {
			Worked = false
			fmt.Println("Error in creating file")
		} else {
			Worked = true
		}
	}
}

func writeFile(path string, text []byte) {
	if Worked {
		if err := os.WriteFile(path, text, os.ModePerm); err != nil {
			Worked = false
			fmt.Println("Error in writing file: ", err)
		} else {
			Worked = true
		}
	}
}
