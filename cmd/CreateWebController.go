package cmd

import (
	"fmt"
	"strings"

	"github.com/rhabichl/schmidi/cmd/helper"
)

func CreateWebController(c *Class) (*Class, *Record) {
	var dto Record

	cont := createWebControllerClass(c)
	if len(c.hasOne) != 0 {
		dto = *createWebControllerDTO(c)
	}

	return cont, &dto
}

func createWebControllerClass(c *Class) *Class {
	var webController Class

	if c.isImported(helper.GetImport("UUID")) {
		webController.Imports = append(webController.Imports, Import{Name: helper.GetImport("UUID")})
	}
	webController.PackageName = packageName + ".controller;"
	webController.Imports = append(webController.Imports, Import{Name: fmt.Sprintf("%s.model.%s", packageName, c.Name)})
	webController.Imports = append(webController.Imports, Import{Name: fmt.Sprintf("%s.repository.%sRepository", packageName, c.Name)})
	webController.Imports = append(webController.Imports, Import{Name: "org.springframework.beans.factory.annotation.Autowired"})
	webController.Imports = append(webController.Imports, Import{Name: "org.springframework.http.HttpStatus"})
	webController.Imports = append(webController.Imports, Import{Name: "org.springframework.http.ResponseEntity"})
	webController.Imports = append(webController.Imports, Import{Name: "org.springframework.web.bind.annotation.*"})

	webController.Annotaions = append(webController.Annotaions, "@RestController")
	webController.Annotaions = append(webController.Annotaions, fmt.Sprintf("@RequestMapping(\"/%s\")", strings.ToLower(c.Name)))
	webController.Name = fmt.Sprintf("%sRestController", c.Name)

	// add repository
	webController.Variables = append(webController.Variables, Variable{
		Name:       fmt.Sprintf("%sRepository", strings.ToLower(c.Name)),
		DataType:   fmt.Sprintf("%sRepository", c.Name),
		Annotaions: []string{"@Autowired"},
	})

	if len(c.hasOne) != 0 {
		webController.Imports = append(webController.Imports, Import{Name: fmt.Sprintf("%s.DTO.%sDTO", packageName, c.Name)})

		for _, v := range c.hasOne {
			webController.Imports = append(webController.Imports, Import{Name: fmt.Sprintf("%s.repository.%sRepository", packageName, v.Name)})
			webController.Variables = append(webController.Variables, Variable{
				Name:       fmt.Sprintf("%sRepository", strings.ToLower(v.Name)),
				DataType:   fmt.Sprintf("%sRepository", v.Name),
				Annotaions: []string{"@Autowired"},
			})
		}
	}

	// add function
	if len(c.hasOne) != 0 {
		f := Function{
			Name:       "add",
			Security:   "public",
			ReturnType: "ResponseEntity<?>",
			Parameter: []FunctionParameter{{
				Name:       strings.ToLower(c.Name) + "DTO",
				DataType:   c.Name + "DTO",
				Annotaions: []string{"@RequestBody"},
			}},
			Annotaions: []string{"@PostMapping(\"/create\")"},
		}

		for _, v := range c.hasOne {
			f.Code = append(f.Code, fmt.Sprintf("if (%sRepository.findById(%sDTO.%sId()).isEmpty()) {", strings.ToLower(v.Name), strings.ToLower(c.Name), strings.ToLower(v.Name)))
			f.Code = append(f.Code, "\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);")
			f.Code = append(f.Code, "}")
			f.Code = append(f.Code, "")
		}

		f.Code = append(f.Code, fmt.Sprintf("%s %s = new %s();", c.Name, strings.ToLower(c.Name), c.Name))
		for _, v := range c.hasOne {
			f.Code = append(f.Code, fmt.Sprintf("%s.set%s(%sRepository.findById(%sDTO.%sId()).get());", strings.ToLower(c.Name), v.Name, strings.ToLower(v.Name), strings.ToLower(c.Name), strings.ToLower(v.Name)))
		}
		va := c.getVarsWithoutNtoOne()
		for _, v := range va {
			f.Code = append(f.Code, fmt.Sprintf("%s.set%s(%sDTO.%s());", strings.ToLower(c.Name), strings.Title(v.Name), strings.ToLower(c.Name), v.Name))
		}
		f.Code = append(f.Code, "")
		f.Code = append(f.Code, fmt.Sprintf("%sRepository.save(%s);", strings.ToLower(c.Name), strings.ToLower(c.Name)), "")
		f.Code = append(f.Code, "return ResponseEntity.ok(null);")

		webController.Functions = append(webController.Functions, f)
	} else {
		webController.Functions = append(webController.Functions, Function{
			Name:       "add",
			Security:   "public",
			ReturnType: "ResponseEntity<?>",
			Parameter: []FunctionParameter{{
				Name:       strings.ToLower(c.Name),
				DataType:   c.Name,
				Annotaions: []string{"@RequestBody"},
			}},
			Annotaions: []string{"@PostMapping(\"/create\")"},
			Code: []string{
				fmt.Sprintf("%sRepository.save(%s);", strings.ToLower(c.Name), strings.ToLower(c.Name)),
				"",
				"return ResponseEntity.ok(null);",
			},
		})
	}
	if len(c.hasOne) != 0 {
		f := Function{
			Name:       "update",
			Security:   "public",
			ReturnType: "ResponseEntity<?>",
			Parameter: []FunctionParameter{{
				Name:       strings.ToLower(c.Name) + "DTO",
				DataType:   c.Name + "DTO",
				Annotaions: []string{"@RequestBody"},
			}},
			Annotaions: []string{"@PostMapping(\"/update\")"},
			Code: []string{
				fmt.Sprintf("if(%sRepository.findById(%sDTO.id()).isEmpty()){", strings.ToLower(c.Name), strings.ToLower(c.Name)),
				"\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);",
				"}",
				"",
			},
		}

		for _, v := range c.hasOne {
			f.Code = append(f.Code, fmt.Sprintf("if (%sRepository.findById(%sDTO.%sId()).isEmpty()) {", strings.ToLower(v.Name), strings.ToLower(c.Name), strings.ToLower(v.Name)))
			f.Code = append(f.Code, "\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);")
			f.Code = append(f.Code, "}")
			f.Code = append(f.Code, "")
		}

		f.Code = append(f.Code, fmt.Sprintf("%s %supdate = %sRepository.findById(%sDTO.id()).get();", c.Name, strings.ToLower(c.Name), strings.ToLower(c.Name), strings.ToLower(c.Name)))
		for _, v := range c.hasOne {
			f.Code = append(f.Code, fmt.Sprintf("%supdate.set%s(%sRepository.findById(%sDTO.%sId()).get());", strings.ToLower(c.Name), v.Name, strings.ToLower(v.Name), strings.ToLower(c.Name), strings.ToLower(v.Name)))
		}
		va := c.getVarsWithoutNtoOne()
		for _, v := range va {
			f.Code = append(f.Code, fmt.Sprintf("%supdate.set%s(%sDTO.%s());", strings.ToLower(c.Name), strings.Title(v.Name), strings.ToLower(c.Name), v.Name))
		}
		f.Code = append(f.Code, "")
		f.Code = append(f.Code, fmt.Sprintf("%sRepository.save(%supdate);", strings.ToLower(c.Name), strings.ToLower(c.Name)), "")
		f.Code = append(f.Code, "return ResponseEntity.ok(null);")

		webController.Functions = append(webController.Functions, f)
	} else {
		// update function
		f := Function{
			Name:       "update",
			Security:   "public",
			ReturnType: "ResponseEntity<?>",
			Parameter: []FunctionParameter{{
				Name:       strings.ToLower(c.Name),
				DataType:   c.Name,
				Annotaions: []string{"@RequestBody"},
			}},
			Annotaions: []string{"@PostMapping(\"/update\")"},
			Code: []string{
				fmt.Sprintf("if(%sRepository.findById(%s.getId()).isEmpty()){", strings.ToLower(c.Name), strings.ToLower(c.Name)),
				"\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);",
				"}",
				"",
				fmt.Sprintf("%s update%s = %sRepository.findById(%s.getId()).get();", c.Name, c.Name, strings.ToLower(c.Name), strings.ToLower(c.Name)),
			},
		}

		for _, v := range c.Variables {
			f.Code = append(f.Code, fmt.Sprintf("update%s.set%s(%s.get%s());", c.Name, strings.Title(v.Name), strings.ToLower(c.Name), strings.Title(v.Name)))
		}

		f.Code = append(f.Code, "", fmt.Sprintf("%sRepository.save(update%s);", strings.ToLower(c.Name), c.Name))
		f.Code = append(f.Code, "", "return ResponseEntity.ok(null);")
		webController.Functions = append(webController.Functions, f)
	}

	// delete function
	webController.Functions = append(webController.Functions, Function{
		Name:       "delete",
		Security:   "public",
		ReturnType: "ResponseEntity<?>",
		Parameter: []FunctionParameter{{
			Name:       "id",
			DataType:   c.ID.DataType,
			Annotaions: []string{"@PathVariable"},
		}},
		Annotaions: []string{"@DeleteMapping(\"/delete/{id}\")"},
		Code: []string{
			fmt.Sprintf("if(%sRepository.findById(id).isEmpty()){", strings.ToLower(c.Name)),
			"\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);",
			"}",
			"",
			fmt.Sprintf("%sRepository.deleteById(id);", strings.ToLower(c.Name)),
			"",
			"return ResponseEntity.ok(null);",
		},
	})

	// load function
	webController.Functions = append(webController.Functions, Function{
		Name:       "load",
		Security:   "public",
		ReturnType: "ResponseEntity<?>",
		Parameter: []FunctionParameter{{
			Name:       "id",
			DataType:   c.ID.DataType,
			Annotaions: []string{"@PathVariable(required = false)"},
		}},
		Annotaions: []string{"@GetMapping(value = {\"/load\", \"/load/{id}\"})"},
		Code: []string{
			"if(id != null){",
			fmt.Sprintf("\tif(%sRepository.findById(id).isEmpty()){", strings.ToLower(c.Name)),
			"\t\treturn ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);",
			"\t}",
			"",
			fmt.Sprintf("\treturn ResponseEntity.ok(%sRepository.findById(id).get());", strings.ToLower(c.Name)),
			"}else{",
			fmt.Sprintf("\treturn ResponseEntity.ok(%sRepository.findAll());", strings.ToLower(c.Name)),
			"}",
		},
	})

	return &webController
}

func createWebControllerDTO(c *Class) *Record {
	var r Record
	r.PackageName = packageName + ".DTO"
	r.Name = c.Name + "DTO"
	r.Parameter = append(r.Parameter, FunctionParameter{
		Name:     c.ID.Name,
		DataType: c.ID.DataType,
	})

	for _, v := range c.getVarsWithoutNtoOne() {
		r.Parameter = append(r.Parameter, FunctionParameter{
			Name:     v.Name,
			DataType: v.DataType,
		})
	}

	for _, v := range c.hasOne {
		r.Parameter = append(r.Parameter, FunctionParameter{
			Name:     strings.ToLower(v.Name) + "Id",
			DataType: v.ID.DataType,
		})
	}

	r.Import()
	return &r
}
