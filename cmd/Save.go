package cmd

import (
	"fmt"
	"os"

	"github.com/rhabichl/schmidi/cmd/helper"
)

var (
	files []*helper.Fi
)

func Save() {
	for _, v := range files {
		v.Save()
	}

	fmt.Println("Make sure you have lombok as a dependency")
	os.Exit(0)
}
