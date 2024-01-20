package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func EditStructure(oldName string, newName string, newUrl string, newLabel string) {
	file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var structures Structures
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&structures)
	if err != nil {
		fmt.Println(err)
		return
	}

	structData, ok := structures.StructData[oldName]
	if !ok {
		fmt.Printf("Structure with name %s not found.\n", oldName)
		return
	}

	delete(structures.StructData, oldName)

	if len(newLabel) >= 1 {
		structData.Label = newLabel
	}
	if len(newUrl) >= 1 {
		structData.Link = newUrl
	}

	if len(newName) >= 1 {
		structures.StructData[newName] = structData
	} else {
		structures.StructData[oldName] = structData
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(structures)
	if err != nil {
		fmt.Println(err)
	}
}
