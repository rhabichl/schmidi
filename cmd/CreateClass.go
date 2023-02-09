package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rhabichl/schmidi/cmd/helper"
)

var (
	isImported = make(map[string]bool)
)

func CreateClass() (string, bytes.Buffer, bytes.Buffer, bytes.Buffer) {
	tmpName := helper.PromptGetInput(helper.NewPromtContent("Name of the class", "enter a correct classname"))
	// make every classname uppercase to make teacher happy

	var im, va, fu bytes.Buffer

	im.WriteString(helper.GetImport("Jpa"))
	im.WriteString(helper.GetImport("lombok"))

	tim, tva := getId()
	im.Write(tim.Bytes())
	va.Write(tva.Bytes())
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
			tim, tva := askForVariable()
			im.Write(tim.Bytes())
			va.Write(tva.Bytes())
		}
	}
	return strings.Title(tmpName), im, va, fu
}

func getId() (bytes.Buffer, bytes.Buffer) {
	var im, va bytes.Buffer
	va.WriteString("\t@Id\n")
	ty := helper.PromptGetSelect(helper.NewPromtContent("Choose the datatype of the id", "please select one of the items"), []string{"Long", "String", "UUID"})
	switch ty {
	case "Long":
		va.WriteString("\t@GeneratedValue(strategy = GenerationType.IDENTITY)\n")
		va.WriteString("\t@Setter(AccessLevel.NONE)\n")
		va.WriteString("\tprivate Long id;\n\n")
	case "String":
		va.WriteString("\t@GeneratedValue(strategy = GenerationType.IDENTITY)\n")
		va.WriteString("\t@Setter(AccessLevel.NONE)\n")
		va.WriteString("\tprivate String id;\n\n")
	case "UUID":
		im.WriteString(helper.GetImport("UUID"))
		if helper.JavaVersion > 11 {
			va.WriteString("\t@GeneratedValue(strategy = GenerationType.UUID)\n")
			va.WriteString("\t@Column(name = \"id\", nullable = false)\n")
		} else {
			im.WriteString(helper.GetImport("GenericGenerator"))
			va.WriteString("\t@GeneratedValue(generator = \"uuid2\")\n")
			va.WriteString("\t@GenericGenerator(name = \"uuid2\", strategy = \"uuid2\")\n")
			va.WriteString("\t@Column(name = \"id\", updatable = false, nullable = false, columnDefinition = \"VARCHAR(36)\")\n")
			va.WriteString("\t@Type(type = \"uuid-char\")\n")
		}
		va.WriteString("\t@Setter(AccessLevel.NONE)\n")
		va.WriteString("\tprivate UUID id;\n\n")
	}
	return im, va
}

func askForVariable() (bytes.Buffer, bytes.Buffer) {
	var im, variables bytes.Buffer
	variables.WriteString("\tprivate ")

	t := helper.PromptGetSelect(helper.NewPromtContent("Type of class", "select a valid type"), []string{"String", "Integer", "Double", "Date", "Timestamp"})
	switch t {
	case "Date":
		if _, ok := isImported["Date"]; !ok {
			im.WriteString(helper.GetImport("Date"))
			isImported["Date"] = true
		}

	case "Timestamp":
		if _, ok := isImported["Timestamp"]; !ok {
			im.WriteString(helper.GetImport("Timestamp"))
			isImported["Timestamp"] = true
		}
	}
	variables.WriteString(t + " ")

	name := helper.PromptGetInput(helper.NewPromtContent("name:", "enter a valid name"))
	// make sure the var is lower case
	nameNew := strings.ToLower(string(name[0])) + name[1:]

	variables.WriteString(nameNew + ";\n")
	// add annotaion
	var varPlusAnnotaion bytes.Buffer
	varPlusAnnotaion.WriteString(fmt.Sprintf("\t@Column(name = \"%s\")\n", nameNew))
	varPlusAnnotaion.WriteString("\t@NonNull\n")
	varPlusAnnotaion.Write(variables.Bytes())
	varPlusAnnotaion.WriteString("\n")

	return im, varPlusAnnotaion
}
