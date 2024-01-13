package cmd

import (
	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

var editConf = &cobra.Command{
	Use:   "structs",
	Short: "Want to add new structures or edit existing structures?",
	Long: `The structs command is used for editing and adding structures
to the config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		structures.CreateStruct("test3", "Test 3", "https://github.com/ledepede1/default-structure")
	},
}

func init() {
	cobra.OnlyValidArgs(editConf, []string{"edit", "add"})
	rootCmd.AddCommand(editConf)
}
