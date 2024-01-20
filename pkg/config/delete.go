package config

import (
	"encoding/json"
	"fmt"
	"os"
)

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
