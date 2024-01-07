package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forgestruct-cli",
	Short: "Forge Struct CLI is used for creating FiveM scripts",
	Long:  `A CLI that helps with creating FiveM script project structures`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Forge Struct CLI is used for creating FiveM script structures")
		fmt.Print("Get help by using 'fstruct --help'")
	},
}

func Execute() {
	rootCmd.Aliases = []string{"fstruct"}

	// Execute the Cobra CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
