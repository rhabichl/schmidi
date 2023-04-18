package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rhabichl/schmidi/cmd/helper"
)

var (
	isImported = make(map[string]bool)
)

func CreateClass() Class {
	var c Class

	tmpName := helper.PromptGetInput(helper.NewPromtContent("Name of the class", "enter a correct classname"))
	// make every classname uppercase to make teacher happy
	c.Name = strings.Title(tmpName)
	c.PackageName = helper.GetPackageName(pom) + ".model"
	c.Imports = append(c.Imports, Import{Name: helper.GetImport("Jpa")})
	c.Imports = append(c.Imports, Import{Name: helper.GetImport("lombok")})

	id, im := getId()
	c.Imports = append(c.Imports, im...)
	c.ID = id

	var finished = false
	for !finished {

		prompt := promptui.Prompt{
			Label:     "New variable",
			IsConfirm: true,
		}
		result, _ := prompt.Run()
		finished = true
		if result == "y" {
			finished = false
			// ceate new var
			va, im := askForVariable(c)
			c.Variables = append(c.Variables, va)
			c.Imports = append(c.Imports, im...)
		}
	}
	return c
}

func getId() (Variable, []Import) {
	var v Variable
	var i []Import

	v.Annotaions = append(v.Annotaions, "@Id")
	v.Name = "id"
	v.Security = "private"

	ty := helper.PromptGetSelect(helper.NewPromtContent("Choose the datatype of the id", "please select one of the items"), []string{"Long", "String", "UUID"})
	switch ty {
	case "Long":
		v.DataType = "Long"
		v.Annotaions = append(v.Annotaions, "@GeneratedValue(strategy = GenerationType.IDENTITY)")
		v.Annotaions = append(v.Annotaions, "@Setter(AccessLevel.NONE)")

	case "String":
		v.DataType = "String"
		v.Annotaions = append(v.Annotaions, "@GeneratedValue(strategy = GenerationType.IDENTITY)")
		v.Annotaions = append(v.Annotaions, "@Setter(AccessLevel.NONE)")

	case "UUID":

		i = append(i, Import{Name: helper.GetImport("UUID")})
		if helper.JavaVersion > 11 {
			v.Annotaions = append(v.Annotaions, "@GeneratedValue(strategy = GenerationType.UUID)")
			v.Annotaions = append(v.Annotaions, "@Column(name = \"id\", nullable = false)")

		} else {
			i = append(i, Import{Name: helper.GetImport("GenericGenerator")})

			v.Annotaions = append(v.Annotaions, "@GeneratedValue(generator = \"uuid2\")")
			v.Annotaions = append(v.Annotaions, "@GenericGenerator(name = \"uuid2\", strategy = \"uuid2\")")
			v.Annotaions = append(v.Annotaions, "@Column(name = \"id\", updatable = false, nullable = false, columnDefinition = \"VARCHAR(36)\")")
			v.Annotaions = append(v.Annotaions, "@Type(type = \"uuid-char\")")
		}
		v.Annotaions = append(v.Annotaions, "@Setter(AccessLevel.NONE)")
		v.DataType = "UUID"
	}
	return v, i
}

func askForVariable(c Class) (Variable, []Import) {
	var v Variable
	var i []Import
	v.Security = "private"

	t := helper.PromptGetSelect(helper.NewPromtContent("Type of class", "select a valid type"), []string{"String", "Integer", "Double", "Date", "Timestamp"})
	switch t {
	case "Date":

		if !c.isImported(helper.GetImport("Date")) {
			i = append(i, Import{Name: helper.GetImport("Date")})
		}

	case "Timestamp":
		if !c.isImported(helper.GetImport("Timestamp")) {
			i = append(i, Import{Name: helper.GetImport("Timestamp")})
		}

	}

	name := helper.PromptGetInput(helper.NewPromtContent("name:", "enter a valid name"))
	// make sure the var is lower case
	nameNew := strings.ToLower(string(name[0])) + name[1:]
	v.Name = nameNew
	v.DataType = t
	// add annotaion
	v.Annotaions = append(v.Annotaions, fmt.Sprintf("@Column(name = \"%s\")", nameNew))
	v.Annotaions = append(v.Annotaions, "@NonNull")

	return v, i
}
