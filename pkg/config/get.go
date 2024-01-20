package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetStructData(structType string) (StructData, bool) {
	var structures Structures

	cfgFile, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		return StructData{}, false
	}
	decoder := json.NewDecoder(cfgFile)
	decoder.Decode(&structures)

	data, found := structures.StructData[structType]
	if found {
		return data, false
	}

	return StructData{}, false
}

func GetAllStructs() map[string]StructData {
	var structures Structures

	data, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&structures)
	if err != nil {
		fmt.Println(err)
	}

	return structures.StructData
}
