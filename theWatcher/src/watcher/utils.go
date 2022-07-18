package watcher

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

type PathToTag struct {
	Table map[string]map[string]bool
}

type TagToPath struct {
	Table map[string]map[string]bool
}

func Hold() {
	var input string
	fmt.Println("\nPlease press ENTER to continue...")
	fmt.Scanln(&input)
}

func ManageError(err error) {
	color.Red(err.Error())
	Hold()
	os.Exit(1)
}
func RecoverPathToTagTable(path string) (PathToTag, error) {
	source, err := os.Open(path)
	if err != nil {
		return PathToTag{}, err
	}
	defer source.Close()
	dec := gob.NewDecoder(source)

	var table PathToTag
	if err := dec.Decode(&table); err == io.EOF {
		table = PathToTag{
			Table: map[string]map[string]bool{},
		}
	} else if err != nil {
		return PathToTag{}, err
	}
	return table, nil
}

func RecoverTagToPathTable(path string) (TagToPath, error) {
	source, err := os.Open(path)
	if err != nil {
		return TagToPath{}, err
	}
	defer source.Close()
	dec := gob.NewDecoder(source)

	var table TagToPath
	if err := dec.Decode(&table); err == io.EOF {
		table = TagToPath{
			Table: map[string]map[string]bool{},
		}
	} else if err != nil {
		return TagToPath{}, err
	}
	return table, nil
}
