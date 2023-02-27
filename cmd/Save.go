package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/rhabichl/schmidi/cmd/helper"
)

var (
	files []*helper.Fi
)

func Save() {

	prompt := promptui.Prompt{
		Label:     "Make repositories for the classes (only supports UUID at the time)",
		IsConfirm: true,
	}
	result, _ := prompt.Run()

	if result == "y" {
		for _, v := range files {
			v := helper.Fi{
				Name: v.Name,
			}
			v.Save(v.PathRepo(), v.Content.BytesRepo(v.Name))
		}
	}

	for _, v := range files {
		v.Save(v.Path(), v.Content.BytesClass(v.Name))
	}

	fmt.Println("Make sure you have lombok as a dependency")
	os.Exit(0)
}
