package cmd

import (
	"fmt"
	"strings"
)

type Class struct {
	Name        string
	PackageName string
	Imports     []Import
	Annotaions  []string
	Variables   []Variable
	ID          Variable
}

type Variable struct {
	Name       string
	DataType   string
	Annotaions []string
	Security   string
}

func (v *Variable) Get() (string, error) {
	var sb strings.Builder
	for _, v1 := range v.Annotaions {
		sb.WriteString("\t" + v1 + "\n")
	}
	fmt.Println("Datatype", v.DataType)

	sb.WriteString(fmt.Sprintf("\t%s %s %s;\n", v.Security, v.DataType, v.Name))

	return sb.String(), nil
}

type Import struct {
	Name string
}

func (c *Class) isImported(imp string) bool {
	for _, v := range c.Imports {
		if v.Name == imp {
			return true
		}
	}
	return false
}

func (c *Class) Get() (string, error) {
	var sb strings.Builder

	// write packagename
	sb.WriteString(fmt.Sprintf("package %s;\n", c.PackageName))
	sb.WriteString("\n")

	// write imports
	for _, v := range c.Imports {
		sb.WriteString(fmt.Sprintf("import %s;\n", v.Name))
	}
	sb.WriteString("\n")

	// write Annotaions
	for _, v := range c.Annotaions {
		sb.WriteString(v + "\n")
	}
	// write Class
	sb.WriteString(fmt.Sprintf("public class %s {\n", c.Name))

	tmpId, _ := c.ID.Get()
	sb.WriteString("\n" + tmpId + "\n")

	// write vars
	for _, v := range c.Variables {
		tmp, _ := v.Get()
		sb.WriteString("\n" + tmp + "\n")
	}

	sb.WriteString("}")

	return sb.String(), nil
}

type Printable interface {
	Get() (string, error)
}
