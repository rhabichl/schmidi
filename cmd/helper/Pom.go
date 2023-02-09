package helper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Read the pom.xml file and parses it into a struct
*/
func ReadPomXml(path string) PomXml {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s%cpom.xml", path, os.PathSeparator))
	if err != nil {
		panic(fmt.Sprintf("couldn't read the pom.xml file:%s", err))
	}
	var p PomXml
	err = xml.Unmarshal(content, &p)
	if err != nil {
		panic(fmt.Sprintf("couldn't parse pom.xml file:%s", err))
	}
	return p
}

func WritePomXml(path string, p PomXml) {

	b, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s%cpom.xml.bak", path, os.PathSeparator), b, 0664)
	if err != nil {
		fmt.Println(err)
	}
}
