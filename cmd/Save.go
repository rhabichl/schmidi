package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/rhabichl/schmidi/cmd/helper"
)

func Save() {

	prompt := promptui.Prompt{
		Label:     "Make repositories for the classes",
		IsConfirm: true,
	}
	result, _ := prompt.Run()

	if result == "y" {
		for _, v := range classes {
			b := helper.GenerateRepo(v.Name, v.ID.DataType)
			writeFile(helper.PathRepo(v.Name), b.Bytes())
		}
	}

	for _, v := range classes {
		b, _ := v.Get()
		writeFile(helper.PathModel(v.Name), []byte(b))
	}

	fmt.Println("Make sure you have lombok as a dependency")
	os.Exit(0)
}

func writeFile(p string, body []byte) {

	err := os.Mkdir(filepath.Dir(p), 0777)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(p, body, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
