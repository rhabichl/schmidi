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

var (
	classes map[string]*Class
)

func ActionLoop() {
	fmt.Println(packageName)
	fmt.Println(helper.GetPathForSaving(pom, path))
	classes = make(map[string]*Class)

	h := helper.NewPromtContent("what do you want to do:", "please choose one of the things below")
	for {
		d := helper.GetActionPromt(h, []string{CREATE_CLASS, RELATIONSHIPS, SAVE})
		switch d {
		case CREATE_CLASS:
			c := CreateClass()
			classes[c.Name] = &c

		case RELATIONSHIPS:
			CreateRelationship()

		case SAVE:
			Save()
		}

	}

}
