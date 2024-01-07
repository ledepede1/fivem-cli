package cmd

import (
	"fmt"

	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "The 'create' command is used for creating a new structured project.",
	Long: `The 'create' commad is used to create structured projects it
will ask you about the project name and what kind of project 
youd like to get`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("List of avaible structures: \nDefault: default\nReactjs: reactjs")
		} else {
			structureType := string(args[0])
			structures.StrucutreCreate(structureType)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
