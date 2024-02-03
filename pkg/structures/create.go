package structures

import (
	"bufio"
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
	data, _ := config.GetStructData(structType)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	// Cloning the repo
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      data.Link,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Print("Error: ", err)
		return false
	}

	// Removing the .git folder
	err = os.RemoveAll(path + "/.git")
	if err != nil {
		fmt.Println("Error in deleting .git folder")
		return false
	}

	var configData = config.GetUserConfig()
	if configData.UseUserName {
		fmt.Println("hey")

		file, err := os.OpenFile(path+"/fxmanifest.lua", os.O_RDWR, 0644)
		if err != nil {
			if err == os.ErrNotExist {
				file, err = os.OpenFile(path+"__resource.lua", os.O_RDWR, 0644)
				if err == os.ErrNotExist {
					fmt.Println("There is no __resource.lua or fxmanifest.lua in this structure!")
					return false
				}
			} else {
				fmt.Println(err)
				return false
			}
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var modifiedLines []string
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, "author") {
				line = fmt.Sprintf("author '%s'", configData.Username)
			}
			if strings.Contains(line, "description") {
				line = "description 'new_description'"
			}

			modifiedLines = append(modifiedLines, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file:", err)
			return false
		}

		// Truncate the file before writing
		err = file.Truncate(0)
		if err != nil {
			fmt.Println("Error truncating file:", err)
			return false
		}

		// Move the file offset to the beginning
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Println("Error seeking file:", err)
			return false
		}

		// Join the modified lines into a single string
		modifiedContent := strings.Join(modifiedLines, "\n")

		// Write the modified content to the file
		_, err = file.WriteString(modifiedContent)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return false
		}

		return true
	} else {
		return true
	}
}
