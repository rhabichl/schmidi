package helper

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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

type FileContent struct {
	Import    bytes.Buffer
	Variables bytes.Buffer
	Functions bytes.Buffer
}

func (f *FileContent) BytesClass(name string) []byte {
	//f.Variables.Write(f.Import.Bytes())
	b := GenerateClass(f.Import, *bytes.NewBufferString(name), f.Variables)
	return b.Bytes()
}

func (f *FileContent) BytesRepo(name, idDataType string) []byte {
	b := GenerateRepo(name, idDataType)
	return b.Bytes()
}

// abstraction for file
type Fi struct {
	Name    string
	IdType  string
	Content FileContent
}

func PathRepo(name string) string {
	return fmt.Sprintf("%s%crepository%c%sRepository.java", GetPathForSaving(pom, path), os.PathSeparator, os.PathSeparator, name)
}

func PathModel(name string) string {
	return fmt.Sprintf("%s%cmodel%c%s.java", GetPathForSaving(pom, path), os.PathSeparator, os.PathSeparator, name)
}

func (f *Fi) Save(p string, body []byte) {

	err := os.Mkdir(filepath.Dir(p), 0777)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(p, body, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateClass(im, name, content bytes.Buffer) bytes.Buffer {
	var sb strings.Builder
	// grow the stringbuilder once
	sb.Grow(18 + len(packageName) + im.Len() + 26 + name.Len() + content.Len() + 3)
	// print the name of the package
	sb.WriteString(fmt.Sprintf("package %s.model;\n\n", packageName))
	// write the imports
	sb.WriteString(im.String())
	// write the class
	sb.WriteString(fmt.Sprintf("\n@Entity\n@Getter\n@Setter\n@RequiredArgsConstructor\n@NoArgsConstructor\npublic class %s {\n\n", &name))
	// write content
	sb.WriteString(content.String())
	// close the class
	sb.WriteString("\n}")
	return *bytes.NewBuffer([]byte(sb.String()))
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
