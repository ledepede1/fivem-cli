package structures

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/ledepede1/forgestruct-cli/pkg/config"
)

func StrucutreCreate(structure string, path string) {
	var lowercaseStruct = strings.ToLower(structure)

	if CreateStructure(lowercaseStruct, path) {
		fmt.Println("Successfully created structure")
	} else {
		fmt.Println("That structure type doesn't exist!")
	}
}

func CreateStructure(structType string, path string) bool {
	repoUrl, _, error := config.GetStructData(structType)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	// Cloning the repo
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Print("Error: ", err)
	}

	// Removing the .git folder
	err = os.RemoveAll(path + "/.git")
	if err != nil {
		fmt.Println("Error in deleting .git folder")
	}

	return error
}
