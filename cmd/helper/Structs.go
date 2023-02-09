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

type FileContent struct {
	Import    bytes.Buffer
	Variables bytes.Buffer
	Functions bytes.Buffer
}

func (f *FileContent) Bytes(name string) []byte {
	//f.Variables.Write(f.Import.Bytes())
	b := GenerateClass(f.Import, *bytes.NewBufferString(name), f.Variables)
	return b.Bytes()
}

// abstraction for file
type Fi struct {
	Name    string
	Content FileContent
}

func (f *Fi) Path() string {
	return fmt.Sprintf("%s%cmodel%c%s.java", GetPathForSaving(pom, path), os.PathSeparator, os.PathSeparator, f.Name)
}

func (f *Fi) Save() {

	err := os.WriteFile(f.Path(), f.Content.Bytes(f.Name), 0644)
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
