package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Structures struct {
	StructData map[string]StructData `json:"structures"`
}

type StructData struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type StructDataWithKey struct {
	Key  string     `json:"key"`
	Data StructData `json:"data"`
}

func GetStructData(structType string) (string, string, bool) {
	var structures Structures

	cfgFile, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		return "", "", false
	}
	decoder := json.NewDecoder(cfgFile)
	decoder.Decode(&structures)

	data, found := structures.StructData[structType]
	if found {
		return data.Link, data.Link, false
	}

	return "", "", false
}

func GetAllStructs() []StructDataWithKey {
	var structures Structures
	var structs = []StructDataWithKey{}

	data, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return structs
	}

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&structures)
	if err != nil {
		return structs
	}

	for k, v := range structures.StructData {
		structs = append(structs, StructDataWithKey{Key: k, Data: v})
	}

	return structs
}

func DeleteStructure(name string) {
	file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return
	}

	delStruct(data, name)

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

func delStruct(data map[string]interface{}, name string) {
	if structures, ok := data["structures"].(map[string]interface{}); ok {
		delete(structures, name)
	} else {
		fmt.Println("Error: 'structures' key not found in the JSON data.")
	}
}
