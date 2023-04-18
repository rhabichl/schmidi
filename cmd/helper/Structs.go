package helper

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	packageName string
	pom         PomXml
	path        string
)

func SetPath(p string) {
	path = p
}

func SetPomXML(p PomXml) {
	pom = p
}

func SetPackageName(i string) {
	packageName = i
}

func PathRepo(name string) string {
	return fmt.Sprintf("%s%crepository%c%sRepository.java", GetPathForSaving(pom, path), os.PathSeparator, os.PathSeparator, name)
}

func PathModel(name string) string {
	return fmt.Sprintf("%s%cmodel%c%s.java", GetPathForSaving(pom, path), os.PathSeparator, os.PathSeparator, name)
}

func GenerateRepo(name, idDataType string) bytes.Buffer {
	var sb strings.Builder
	// grow the stringbuilder once
	// print the name of the package
	sb.WriteString(fmt.Sprintf("package %s.repository;", packageName))
	sb.WriteString("\n\nimport org.springframework.data.jpa.repository.JpaRepository;")
	sb.WriteString("\nimport java.util.UUID;")
	sb.WriteString(fmt.Sprintf("\n\nimport %s.model.%s;", packageName, name))
	// write the imports
	sb.WriteString(fmt.Sprintf("\n\npublic interface %sRepository extends JpaRepository<%s, %s> {}", name, name, idDataType))
	// write content
	return *bytes.NewBuffer([]byte(sb.String()))
}
