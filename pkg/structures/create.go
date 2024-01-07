package structures

import (
	"fmt"
	"strings"

	"github.com/ledepede1/forgestruct-cli/pkg/config"
)

func StrucutreCreate(structure string) {
	var lowercaseStruct = strings.ToLower(structure)

	if checkStructure(structure) {
		switch lowercaseStruct {
		case "default":
			{
				DefaultStrucute("")
			}
		}
	} else {
		fmt.Println("That structure type doesn't exist!")
	}
}

func checkStructure(structure string) bool {
	for _, v := range config.Structures {
		if structure == v {
			return true
		}
	}
	return false
}
