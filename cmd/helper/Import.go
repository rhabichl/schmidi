package helper

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

	return result
}
