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
