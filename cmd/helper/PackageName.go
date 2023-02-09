package helper

import (
	"fmt"
	"os"
	"strings"
)

func GetPackageName(pom PomXml) string {
	return strings.ToLower(fmt.Sprintf("%s.%s", pom.GroupId.Text, pom.ArtifactId.Text))
}

func GetPathForSaving(pom PomXml, path string) string {
	return fmt.Sprintf("%s%csrc%cmain%cjava%c%s", path,
		os.PathSeparator, os.PathSeparator, os.PathSeparator, os.PathSeparator,
		strings.ReplaceAll(GetPackageName(pom), ".", string(os.PathSeparator)))
}
