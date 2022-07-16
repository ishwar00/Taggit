package main

import (
	"os"

	"github.com/ishwar00/Taggit/cmd"
)

func main() {
	cmd.Execute(os.Args[1])
}
