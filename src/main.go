package main

import (
	"os"
	"fmt"

	"github.com/ishwar00/Taggit/cmd"
)

const doc = `
use: Taggit.exe filepath
`


func main() {
	
	if len(os.Args) != 2 {
		err := fmt.Errorf(doc)
		panic(err)
	}

	cmd.Execute(os.Args[1])
}

