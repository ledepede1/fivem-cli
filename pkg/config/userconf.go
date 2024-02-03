package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func EditUsername(userName string) {
	if len(userName) >= 1 {
		file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\userconfig.json`, os.O_RDWR, 0660)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		var userNameStruct UserName

		userNameStruct.Username = userName
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&userNameStruct.Username)
		if err != nil {
			fmt.Println(err)
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
		err = encoder.Encode(userNameStruct)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func UseName(useName bool) {
	fmt.Println(useName)
	file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\userconfig.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var userNameStruct UserName

	userNameStruct.UseUserName = useName
	userNameStruct.Username = GetUserConfig().Username

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(userNameStruct)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userNameStruct)
}

func GetUserConfig() UserName {
	var userNameStruct UserName

	file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\userconfig.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return userNameStruct
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&userNameStruct)
	if err != nil {
		fmt.Println(err)
	}

	return userNameStruct
}
