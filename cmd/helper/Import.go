package helper

import (
	"bytes"
	"strings"
)

var JavaVersion = 0

func GetImport(input string) string {
	var result = ""
	switch input {
	case "Date":
		result = "java.sql.Date"
	case "Timestamp":
		result = "java.sql.Timestamp"
	case "Jpa":
		if JavaVersion > 11 {
			result = "jakarta.persistence.*"
		} else {
			result = "javax.persistence.*"
		}
	case "UUID":
		result = "java.util.UUID"
	case "GenericGenerator":
		result = "org.hibernate.annotations.GenericGenerator"
	case "lombok":
		result = "lombok.*"
	}

	return "import " + result + ";\n"
}

func CheckIfIsImported(b bytes.Buffer, imp string) bool {
	return strings.Contains(b.String(), imp)
}
