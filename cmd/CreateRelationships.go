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
	for _, v := range files {
		tmp = append(tmp, v.Name)
	}
	if len(tmp) < 2 {
		return
	}
	first := helper.PromptGetSelect(helper.NewPromtContent("first classe for the relationship:", "choose on class"), tmp)

	second := helper.PromptGetSelect(helper.NewPromtContent("second classe for the relationship:", "choose on class"), helper.RemoveValue(tmp, first))

	fFirst := helper.GetFileByName(files, first)
	fSecond := helper.GetFileByName(files, second)

	if fFirst == nil || fSecond == nil {
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
		fFirst.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fFirst.Name), fSecond.Name, strings.ToLower(fSecond.Name)))
		fSecond.Content.Variables.WriteString(fmt.Sprintf("\t@ManyToOne(fetch = FetchType.LAZY)\n\t@JoinColumn(name = \"%s_id\")\n\tprivate %s %s;\n\n", strings.ToLower(fFirst.Name), fFirst.Name, strings.ToLower(fFirst.Name)))

		if !helper.CheckIfIsImported(fFirst.Content.Import, "java.util.Set") {
			fFirst.Content.Import.WriteString("import java.util.Set;\n")
		}
		if !helper.CheckIfIsImported(fFirst.Content.Import, "com.fasterxml.jackson.annotation.JsonIgnore") {
			fFirst.Content.Import.WriteString("import com.fasterxml.jackson.annotation.JsonIgnore;\n")
		}

		//fFirst.Content.Import.WriteString(fmt.Sprintf("import %s.model.%s;\n", packageName, fSecond.Name))
		//fSecond.Content.Import.WriteString(fmt.Sprintf("import %s.model.%s;\n", packageName, fFirst.Name))

	case tmpNtO:
		fSecond.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fSecond.Name), fFirst.Name, strings.ToLower(fFirst.Name)))
		fFirst.Content.Variables.WriteString(fmt.Sprintf("\t@ManyToOne(fetch = FetchType.LAZY)\n\t@JoinColumn(name = \"%s_id\")\n\tprivate %s %s;\n\n", strings.ToLower(fSecond.Name), fSecond.Name, strings.ToLower(fSecond.Name)))

		if !helper.CheckIfIsImported(fSecond.Content.Import, "java.util.Set") {
			fSecond.Content.Import.WriteString("import java.util.Set;\n")
		}
		if !helper.CheckIfIsImported(fSecond.Content.Import, "com.fasterxml.jackson.annotation.JsonIgnore") {
			fSecond.Content.Import.WriteString("import com.fasterxml.jackson.annotation.JsonIgnore;\n")
		}
		//fFirst.Content.Import.WriteString(fmt.Sprintf("import %s.model.%s;\n", packageName, fSecond.Name))
		//fSecond.Content.Import.WriteString(fmt.Sprintf("import %s.model.%s;\n", packageName, fFirst.Name))

	case tmpNtM:
		name, im, va, fu, idDataType := CreateClass()
		fFirst.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fFirst.Name), name, strings.ToLower(name)))
		fSecond.Content.Variables.WriteString(fmt.Sprintf("\t@JsonIgnore\n\t@OneToMany(mappedBy = \"%s\", fetch = FetchType.LAZY)\n\tprivate Set<%s> %s;\n\n", strings.ToLower(fSecond.Name), name, strings.ToLower(name)))

		if !helper.CheckIfIsImported(fFirst.Content.Import, "java.util.Set") {
			fFirst.Content.Import.WriteString("import java.util.Set;\n")
		}
		if !helper.CheckIfIsImported(fFirst.Content.Import, "com.fasterxml.jackson.annotation.JsonIgnore") {
			fFirst.Content.Import.WriteString("import com.fasterxml.jackson.annotation.JsonIgnore;\n")
		}
		if !helper.CheckIfIsImported(fSecond.Content.Import, "java.util.Set") {
			fSecond.Content.Import.WriteString("import java.util.Set;\n")
		}
		if !helper.CheckIfIsImported(fSecond.Content.Import, "com.fasterxml.jackson.annotation.JsonIgnore") {
			fSecond.Content.Import.WriteString("import com.fasterxml.jackson.annotation.JsonIgnore;\n")
		}

		va.WriteString(fmt.Sprintf("\t@ManyToOne(fetch = FetchType.LAZY)\n\t@JoinColumn(name = \"%s_id\")\n\tprivate %s %s;\n\n", strings.ToLower(fFirst.Name), fFirst.Name, strings.ToLower(fFirst.Name)))
		va.WriteString(fmt.Sprintf("\t@ManyToOne(fetch = FetchType.LAZY)\n\t@JoinColumn(name = \"%s_id\")\n\tprivate %s %s;\n\n", strings.ToLower(fSecond.Name), fSecond.Name, strings.ToLower(fSecond.Name)))

		files = append(files, &helper.Fi{
			Name:   name,
			IdType: idDataType,
			Content: helper.FileContent{
				Import:    im,
				Variables: va,
				Functions: fu,
			},
		})

	}

}
