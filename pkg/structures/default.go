package structures

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ledepede1/forgestruct-cli/pkg/config"
)

type Files struct {
	Files map[string][]byte
}

type Folders struct {
	Folders []string
}

var files = Files{
	Files: map[string][]byte{
		"Client/Main.lua":       config.DefaultStructure.ClientMain,
		"Server/Main.lua":       config.DefaultStructure.ServerMain,
		"Configs/Shared.lua":    config.DefaultStructure.SharedConfig,
		"Configs/Cl_Config.lua": config.DefaultStructure.ClConfig,
		"Configs/Sv_Config.lua": config.DefaultStructure.SvConfig,
		"/fxmanifest.lua":       config.DefaultStructure.FxManifest,
	},
}

var folders = Folders{
	Folders: []string{
		"/Client",
		"/Server",
		"/Configs",
	},
}

func DefaultStrucute(path string) {
	if path == "" {
		curPath, err := os.Getwd()
		if err != nil {
			fmt.Println("Error: Couldnt get current path!")
		}
		path = curPath
	}

	fmt.Println("Ctrl + C to cancel")
	fmt.Print("Project name: ")
	var name string
	fmt.Scan(&name)

	createProject(path, name)
}

func createProject(userPath string, projname string) {
	// Creating the Project Folder
	var path = filepath.Join(userPath, projname)
	createDir(path)

	// Creating the Folders inside the Project Folder
	for _, v := range folders.Folders {
		var path = filepath.Join(userPath, projname, v)
		createDir(path)
	}

	// Creating the files inside the Folders
	for fileInfo := range files.Files {
		var path = filepath.Join(userPath, projname, fileInfo)
		createFile(path)
	}

	// Writing to the files
	for fileInfo, fileContent := range files.Files {
		var path = filepath.Join(userPath, projname, string(fileInfo))
		writeFile(path, fileContent)
	}

	if Worked {
		fmt.Println("Successfully created your new project!")
	} else {
		fmt.Println("Some errors occured when trying to create your project! Check if folder already exists!")
	}
}
