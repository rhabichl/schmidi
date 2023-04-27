package cmd

import (
	"fmt"
	"strings"

	"github.com/rhabichl/schmidi/cmd/helper"
)

type Class struct {
	Name        string
	PackageName string
	Imports     []Import
	Annotaions  []string
	Variables   []Variable
	ID          Variable
	hasOne      []*Class
	Functions   []Function
}

type Record struct {
	PackageName string
	Name        string
	Imports     []Import
	Parameter   []FunctionParameter
}

type Function struct {
	Name       string
	ReturnType string
	Annotaions []string
	Parameter  []FunctionParameter
	Code       []string
	Security   string
}

type FunctionParameter struct {
	Name       string
	DataType   string
	Annotaions []string
}

func (r *Record) Get() (string, error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package %s.DTO;\n\n", packageName))

	if len(r.Imports) != 0 {
		for _, v := range r.Imports {
			sb.WriteString(fmt.Sprintf("import %s;\n", v.Name))
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("public record %s(", r.Name))
	for i := 0; i < len(r.Parameter)-1; i++ {
		for _, v := range r.Parameter[i].Annotaions {
			sb.WriteString(v + " ")
		}
		sb.WriteString(fmt.Sprintf("%s %s, ", r.Parameter[i].DataType, r.Parameter[i].Name))
	}

	for _, v := range r.Parameter[len(r.Parameter)-1].Annotaions {
		sb.WriteString(v + " ")
	}
	sb.WriteString(fmt.Sprintf("%s %s", r.Parameter[len(r.Parameter)-1].DataType, r.Parameter[len(r.Parameter)-1].Name))
	for _, v := range r.Parameter[len(r.Parameter)-1].Annotaions {
		sb.WriteString(v + " ")
	}
	sb.WriteString("){}")
	return sb.String(), nil
}

func (f *Function) Get() (string, error) {
	var sb strings.Builder
	for _, v := range f.Annotaions {
		sb.WriteString("\t" + v + "\n")
	}

	sb.WriteString(fmt.Sprintf("\t%s %s %s", f.Security, f.ReturnType, f.Name))
	sb.WriteString("(")

	for i := 0; i < len(f.Parameter)-1; i++ {
		for _, v := range f.Parameter[i].Annotaions {
			sb.WriteString(v + " ")
		}
		sb.WriteString(fmt.Sprintf("%s %s, ", f.Parameter[i].DataType, f.Parameter[i].Name))
	}
	for _, v := range f.Parameter[len(f.Parameter)-1].Annotaions {
		sb.WriteString(v + " ")
	}
	sb.WriteString(fmt.Sprintf("%s %s", f.Parameter[len(f.Parameter)-1].DataType, f.Parameter[len(f.Parameter)-1].Name))

	sb.WriteString("){\n")
	for _, v := range f.Code {
		sb.WriteString("\t\t" + v + "\n")
	}
	sb.WriteString("\t}\n")
	return sb.String(), nil
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

	if len(v.Security) != 0 {
		sb.WriteString(fmt.Sprintf("\t%s %s %s;\n", v.Security, v.DataType, v.Name))
	} else {
		sb.WriteString(fmt.Sprintf("\t%s %s;\n", v.DataType, v.Name))
	}

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

	if c.ID.Name != "" {
		tmpId, _ := c.ID.Get()
		sb.WriteString("\n" + tmpId + "\n")
	}

	// write vars
	for _, v := range c.Variables {
		tmp, _ := v.Get()
		sb.WriteString("\n" + tmp + "\n")
	}

	if len(c.Functions) != 0 {
		for _, v := range c.Functions {
			tmp, _ := v.Get()
			sb.WriteString(tmp)
		}
	}

	sb.WriteString("}")

	return sb.String(), nil
}

func (r *Record) Import() {
	for _, v := range r.Parameter {
		if (v.DataType == "Date" || v.DataType == "Timestamp" || v.DataType == "UUID") && !r.IsImported(helper.GetImport(v.DataType)) {
			r.Imports = append(r.Imports, Import{Name: helper.GetImport(v.DataType)})
		}
	}

}

func (r *Record) IsImported(name string) bool {
	for _, v := range r.Imports {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (c *Class) getVarImports() []Import {
	var result []Import

	for _, v := range c.Imports {
		if v.Name == "java.sql.Date" || v.Name == "java.sql.Timestamp" || v.Name == "java.util.UUID" {
			result = append(result, v)
		}
	}

	return result
}

func (c *Class) getVarsWithoutNtoOne() []Variable {
	var result []Variable

	for _, v := range c.Variables {
		if !contains(v.Annotaions, "@JoinColumn") && !contains(v.Annotaions, "@OneToMany") {
			result = append(result, v)
		}
	}
	return result
}

// check if the string starts with the other string
func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if strings.HasPrefix(v, needle) {
			return true
		}
	}
	return false
}

func (c *Class) isNonNullVarPresent() bool {
	for _, variable := range c.Variables {
		for _, annotation := range variable.Annotaions {
			if annotation == "@NonNull" {
				return true
			}
		}
	}

	return false
}

type Printable interface {
	Get() (string, error)
}
