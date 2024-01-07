package cmd

import (
	"fmt"

	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

// forgeCmd represents the forge command
var forgeCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("List of avaible structures: \nDefault: default\nReactjs: reactjs")
		} else {
			structureType := string(args[0])
			structures.StrucutreCreate(structureType)
		}
	},
}

func create(structuretype string) {
	fmt.Println(structuretype)
}

func init() {
	rootCmd.AddCommand(forgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
