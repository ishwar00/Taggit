package main

import (
	"fmt"
	"os"

	"github.com/ishwar00/Taggit/manageTags/cmd"
)

const doc = `
manageTags takes first argument as filePath, to which we are trying manage tags
use:$ manageTags.exe [filepath]

eg: manageTag example/file/path`

func main() {
	if len(os.Args) != 2 {
		errMsg := fmt.Errorf(doc)
		cmd.ManageError(errMsg)
	}

	cmd.Execute(os.Args[1])
}
