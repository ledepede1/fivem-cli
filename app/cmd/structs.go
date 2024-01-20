package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ledepede1/forgestruct-cli/pkg/config"
	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

var Cname string
var label string
var url string

type Structures struct {
	StructData map[string]JsonInfo `json:"structures"`
}

type JsonInfo struct {
	Key string `json:"key"`
}

var structEdit = &cobra.Command{
	Use:   "struct",
	Short: "Want to add a new structures?",
	Long: `The structs command is used for adding structures
to the config.json`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Need arguments: `delete` `add` `edit`")
		} else if args[0] == "delete" {
			if len(Cname) == 0 {
				fmt.Println("Wrong usage of command needs to be `fstruct delete --name {name}`")
			} else {
				if checkJson(Cname) {
					fmt.Println("found this")
					config.DeleteStructure(Cname)
				}
			}
		} else if args[0] == "add" {
			if len(Cname) == 0 || len(label) == 0 || len(url) == 0 {
				fmt.Println("Wrong usage of command needs to be `fstruct add --name {name} --label {label} --url {url}`")
			} else {
				structures.CreateStruct(Cname, label, url)
			}
		} else if args[0] == "edit" {
			if len(args) <= 1 {
				fmt.Println("Wrong usage of commands needs to be `fstruct edit {nameofstructure} optional(--name <newName> --label <newLabel> --url <newUrl>`)")
			} else {
				config.EditStructure(args[1], Cname, url, label)
			}
		}
	},
}

func init() {
	structEdit.Flags().StringVarP(&Cname, "name", "n", "", "Name that will be used in the command")
	structEdit.Flags().StringVarP(&label, "label", "l", "", "Label of the command")
	structEdit.Flags().StringVarP(&url, "url", "u", "", "Url/Link to the github repo")

	rootCmd.AddCommand(structEdit)
}

func checkJson(name string) bool {
	var jsonInfo Structures

	file, err := os.OpenFile(`C:\Program Files\ForgeStruct-CLI\config.json`, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return false
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonInfo)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for k := range jsonInfo.StructData {
		if string(k) == name {
			return true
		}
	}

	return false
}
