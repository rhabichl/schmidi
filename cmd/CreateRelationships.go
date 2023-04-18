package cmd

import (
	"fmt"
	"strings"

	"github.com/rhabichl/schmidi/cmd/helper"
)

const (
	ONE_TO_N = "1 - n"
	N_TO_ONE = "n - 1"
	N_TO_M   = "n - m"
)

func CreateRelationship() {
	var tmp []string
	for _, v := range classes {
		tmp = append(tmp, v.Name)
	}
	if len(tmp) < 2 {
		return
	}
	first := helper.PromptGetSelect(helper.NewPromtContent("first classe for the relationship:", "choose on class"), tmp)

	second := helper.PromptGetSelect(helper.NewPromtContent("second classe for the relationship:", "choose on class"), helper.RemoveValue(tmp, first))

	firstClass := classes[first]
	secondClass := classes[second]

	if firstClass == nil || secondClass == nil {
		return
	}

	tmpOtN := fmt.Sprintf("%s (1*%s - n*%s)", ONE_TO_N, first, second)
	tmpNtO := fmt.Sprintf("%s (n*%s - 1*%s)", N_TO_ONE, first, second)
	tmpNtM := fmt.Sprintf("%s (n*%s - m*%s)", N_TO_M, first, second)

	r := helper.PromptGetSelect(helper.NewPromtContent("what type of relationship should it be:", "choose a type"), []string{
		tmpOtN, tmpNtO, tmpNtM,
	})

	switch r {
	case tmpOtN:
		firstClass.Variables = append(firstClass.Variables, Variable{
			Annotaions: []string{"@JsonIgnore", fmt.Sprintf("@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)", strings.ToLower(firstClass.Name))},
			Security:   "private",
			DataType:   fmt.Sprintf("Set<%s>", secondClass.Name),
			Name:       strings.ToLower(secondClass.Name),
		})

		secondClass.Variables = append(secondClass.Variables, Variable{
			Annotaions: []string{"@ManyToOne(fetch = FetchType.LAZY)", fmt.Sprintf("@JoinColumn(name = \"%s_id\")", strings.ToLower(firstClass.Name))},
			Security:   "private",
			DataType:   firstClass.Name,
			Name:       strings.ToLower(firstClass.Name),
		})

		if !firstClass.isImported("java.util.Set") {
			firstClass.Imports = append(firstClass.Imports, Import{Name: "java.util.Set"})
		}
		if !firstClass.isImported("com.fasterxml.jackson.annotation.JsonIgnore") {
			firstClass.Imports = append(firstClass.Imports, Import{Name: "com.fasterxml.jackson.annotation.JsonIgnore"})
		}

	case tmpNtO:

		secondClass.Variables = append(secondClass.Variables, Variable{
			Annotaions: []string{"@JsonIgnore", fmt.Sprintf("@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)", strings.ToLower(secondClass.Name))},
			Security:   "private",
			DataType:   fmt.Sprintf("Set<%s>", firstClass.Name),
			Name:       strings.ToLower(firstClass.Name),
		})

		firstClass.Variables = append(firstClass.Variables, Variable{
			Annotaions: []string{"@ManyToOne(fetch = FetchType.LAZY)", fmt.Sprintf("@JoinColumn(name = \"%s_id\")", strings.ToLower(secondClass.Name))},
			Security:   "private",
			DataType:   secondClass.Name,
			Name:       strings.ToLower(secondClass.Name),
		})

		if !secondClass.isImported("java.util.Set") {
			secondClass.Imports = append(secondClass.Imports, Import{Name: "java.util.Set"})
		}
		if !secondClass.isImported("com.fasterxml.jackson.annotation.JsonIgnore") {
			secondClass.Imports = append(secondClass.Imports, Import{Name: "com.fasterxml.jackson.annotation.JsonIgnore"})
		}

	case tmpNtM:
		c := CreateClass()
		//fFirst.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fFirst.Name), name, strings.ToLower(name)))
		firstClass.Variables = append(firstClass.Variables, Variable{
			Annotaions: []string{"@JsonIgnore", fmt.Sprintf("@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)", strings.ToLower(firstClass.Name))},
			Security:   "private",
			DataType:   fmt.Sprintf("Set<%s>", c.Name),
			Name:       strings.ToLower(c.Name),
		})

		//fSecond.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fSecond.Name), name, strings.ToLower(name)))
		secondClass.Variables = append(secondClass.Variables, Variable{
			Annotaions: []string{"@JsonIgnore", fmt.Sprintf("@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)", strings.ToLower(secondClass.Name))},
			Security:   "private",
			DataType:   fmt.Sprintf("Set<%s>", c.Name),
			Name:       strings.ToLower(c.Name),
		})

		if !firstClass.isImported("java.util.Set") {
			firstClass.Imports = append(firstClass.Imports, Import{Name: "java.util.Set"})
		}
		if !firstClass.isImported("com.fasterxml.jackson.annotation.JsonIgnore") {
			firstClass.Imports = append(firstClass.Imports, Import{Name: "com.fasterxml.jackson.annotation.JsonIgnore"})
		}
		if !secondClass.isImported("java.util.Set") {
			secondClass.Imports = append(secondClass.Imports, Import{Name: "java.util.Set"})
		}
		if !secondClass.isImported("com.fasterxml.jackson.annotation.JsonIgnore") {
			secondClass.Imports = append(secondClass.Imports, Import{Name: "com.fasterxml.jackson.annotation.JsonIgnore"})
		}

		c.Variables = append(c.Variables, Variable{
			Annotaions: []string{"@ManyToOne(fetch = FetchType.LAZY)", fmt.Sprintf("@JoinColumn(name = \"%s_id\")", strings.ToLower(firstClass.Name))},
			Security:   "private",
			DataType:   firstClass.Name,
			Name:       strings.ToLower(firstClass.Name),
		})

		c.Variables = append(c.Variables, Variable{
			Annotaions: []string{"@ManyToOne(fetch = FetchType.LAZY)", fmt.Sprintf("@JoinColumn(name = \"%s_id\")", strings.ToLower(secondClass.Name))},
			Security:   "private",
			DataType:   secondClass.Name,
			Name:       strings.ToLower(secondClass.Name),
		})

		classes[c.Name] = &c
	}

}
