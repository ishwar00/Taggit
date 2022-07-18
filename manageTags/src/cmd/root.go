package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

var Pt string
var Tp string

func init() {
	buf, err := os.UserHomeDir()
	if err != nil {
		ManageError(err)
	}
	Pt = filepath.Join(buf, "Taggit/pathToTag.db")
	Tp = filepath.Join(buf, "Taggit/tagToPath.db")
	createDB(Pt)
	createDB(Tp)
}

func Execute(path string) error {
	for {
		fmt.Printf("file name: %v\n", path)
		prompt := promptui.Select{
			Label: "Please select operation",
			Items: []string{
				"Add tags",
				"Show tags",
				"Remove tags",
				"Exit",
			},
		}
		i, _, err := prompt.Run()

		if err != nil {
			return err
		}

		operations := []func(string) error{
			AddTags,
			ShowTags,
			RemoveTags,
			Exit,
		}

		if err := operations[i](path); err != nil {
			return err
		}
		Hold()
		clearScreen()
	}
}

func createDB(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		directory := filepath.Dir(path)
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			ManageError(err)
		}

		if _, err := os.Create(path); err != nil {
			ManageError(err)
		}
	} else if err != nil {
		ManageError(err)
	}
}
