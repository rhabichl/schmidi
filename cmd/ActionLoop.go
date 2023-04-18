package cmd

import (
	"fmt"

	"github.com/rhabichl/schmidi/cmd/helper"
)

const (
	CREATE_CLASS  = "create classes"
	RELATIONSHIPS = "relationships"
	SAVE          = "save"
)

func ActionLoop() {
	fmt.Println(packageName)
	fmt.Println(helper.GetPathForSaving(pom, path))
	h := helper.NewPromtContent("what do you want to do:", "please choose one of the things below")
	for {
		d := helper.GetActionPromt(h, []string{CREATE_CLASS, RELATIONSHIPS, SAVE})
		switch d {
		case CREATE_CLASS:
			name, im, va, fu, idDataType := CreateClass()
			files = append(files, &helper.Fi{
				Name:   name,
				IdType: idDataType,
				Content: helper.FileContent{
					Import:    im,
					Variables: va,
					Functions: fu,
				},
			})

		case RELATIONSHIPS:
			CreateRelationship()

		case SAVE:
			Save()
		}

	}

}
