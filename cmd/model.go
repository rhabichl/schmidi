package cmd

import (
	"fmt"
	"strconv"

	"github.com/rhabichl/schmidi/cmd/helper"
	"github.com/spf13/cobra"
)

var (
	pom         helper.PomXml
	packageName string
	path        string
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "To ineract with models",
	Long:  "Please provide a path where the src folder is",
	Run: func(cmd *cobra.Command, args []string) {

		// get the pom.xml file for the configuration
		path = cmd.Flag("path").Value.String()
		pom = helper.ReadPomXml(path)
		version, err := strconv.Atoi(pom.Properties.JavaVersion.Text)
		if err != nil {
			fmt.Println(err)
		}

		helper.JavaVersion = version
		// get the package name
		packageName = helper.GetPackageName(pom)
		helper.SetPackageName(packageName)
		helper.SetPomXML(pom)
		helper.SetPath(path)
		// enter the Action Loop
		ActionLoop()
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
