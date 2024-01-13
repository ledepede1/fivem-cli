package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ledepede1/forgestruct-cli/pkg/config"
	"github.com/ledepede1/forgestruct-cli/pkg/structures"
	"github.com/spf13/cobra"
)

var path string
var name string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "The 'create' command is used for creating a new structured project.",
	Long: `The 'create' commad is used to create structured projects it
will ask you about the project name and what kind of project 
youd like to get`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			var structures = config.GetAllStructs()

			fmt.Println("List of avaible structures:")
			for _, v := range structures {
				fmt.Println("- " + v.Data.Label + " | " + v.Key)
			}

		} else if len(args) >= 1 {
			curPath, err := os.Getwd()
			if err != nil {
				fmt.Println("Error: Couldnt get current path!")
			}
			structureType := string(args[0])

			if name == "" {
				fmt.Print("Project Name: ")
				fmt.Scan(&name)
			}

			if path == "" {
				path = curPath
			}

			finalPath, err := filepath.Abs(filepath.Join(path, name))
			if err != nil {
				fmt.Println(err)
			}

			structures.StrucutreCreate(structureType, finalPath)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&path, "path", "p", "", "Specify the project path")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Specify the project name")
	rootCmd.AddCommand(createCmd)
}
