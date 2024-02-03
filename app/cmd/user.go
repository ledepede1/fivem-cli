package cmd

import (
	"fmt"

	"github.com/ledepede1/forgestruct-cli/pkg/config"
	"github.com/spf13/cobra"
)

var username string
var usename bool

var userConf = &cobra.Command{
	Use:   "user",
	Short: "The 'create' command is used for creating a new structured project.",
	Long: `The 'create' commad is used to create structured projects it
will ask you about the project name and what kind of project 
youd like to get`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Need arguments: `username`, `usename`")
		} else if args[0] != "username" && args[0] != "usename" {
			fmt.Println("Valid command arguments: \nusername\nusename")
		} else if args[0] == "username" {
			if len(username) >= 1 {
				config.EditUsername(username)
			} else {
				fmt.Println("Username needs to be with a length of over 0")
			}
		} else if args[0] == "usename" {
			if usename || !usename {
				config.UseName(usename)
			} else {
				fmt.Println("Need to be a boolean")
			}
		}
	},
}

func init() {
	userConf.Flags().StringVarP(&username, "username", "u", "", "Specify your new Username (Using argument `username`)")
	userConf.Flags().BoolVarP(&usename, "bool", "b", false, "Specifiy wether to use the Username in the fxmanifest.lua/__resource.lua, (Using argument `usename --b=Boolean`)")
	rootCmd.AddCommand(userConf)
}
