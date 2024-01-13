package cmd

import (
	"fmt"

	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

var Cname string
var label string
var url string

var editConf = &cobra.Command{
	Use:   "structcreate",
	Short: "Want to add a new structures?",
	Long: `The structs command is used for adding structures
to the config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(Cname) == 0 || len(label) == 0 || len(url) == 0 {
			fmt.Println("Wrong usage of command needs to be `fstruct --name {name} --label {label} --url {url}`")
		} else {
			structures.CreateStruct(Cname, label, url)
		}
	},
}

func init() {
	editConf.Flags().StringVarP(&Cname, "name", "n", "", "Name that will be used in the command")
	editConf.Flags().StringVarP(&label, "label", "l", "", "Label of the command")
	editConf.Flags().StringVarP(&url, "url", "u", "", "Url/Link to the github repo")

	rootCmd.AddCommand(editConf)
}
