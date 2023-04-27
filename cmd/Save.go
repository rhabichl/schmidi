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
		Label:     "Make repositories for the entities",
		IsConfirm: true,
	}
	result, _ := prompt.Run()

	if result == "y" {
		for _, v := range classes {
			b := helper.GenerateRepo(v.Name, v.ID.DataType)
			writeFile(helper.PathRepo(v.Name), b.Bytes())
		}
	}

	prompt2 := promptui.Prompt{
		Label:     "Make web-controller for the entities",
		IsConfirm: true,
	}
	result2, _ := prompt2.Run()

	if result2 == "y" {
		for _, v := range classes {
			cont, dto := CreateWebController(v)
			if dto.Name != "" {
				bDto, _ := dto.Get()
				writeFile(helper.PathDTO(v.Name), []byte(bDto))
			}

			b, _ := cont.Get()
			writeFile(helper.PathController(v.Name), []byte(b))
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
