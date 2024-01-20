package structures

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

func CreateStruct(structurename string, structurelabel string, url string) {
	var structures Structures
	var structs = make(map[string]StructData)

	data, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer data.Close()

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&structures)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range structures.StructData {
		structs[k] = StructData{Label: v.Label, Link: v.Link}
	}

	structs[structurename] = StructData{Label: structurelabel, Link: url}

	structures.StructData = structs

	_, err = data.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = data.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewEncoder(data)
	err = encoder.Encode(structures)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Created new Structure")
}
