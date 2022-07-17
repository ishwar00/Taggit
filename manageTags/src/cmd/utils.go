package cmd

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/fatih/color"
)

type PathToTag struct {
	Table map[string]map[string]bool
}

func (pt *PathToTag) Print() {
	for path, tags := range pt.Table {
		fmt.Printf("%v: %v\n", path, tags)
	}
}

type TagToPath struct {
	Table map[string]map[string]bool
}

func (tp *TagToPath) print() {
	for tag, paths := range tp.Table {
		fmt.Printf("%v: %v\n", tag, paths)
	}
}

func isValidTag(tag string) bool {

	if len(tag) == 0 {
		return false
	}
	for _, c := range tag {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func printTagsGreen(tags []string) {
	maxLength := 0

	for _, tag := range tags {
		if len(tag) > maxLength {
			maxLength = len(tag)
		}
	}

	for i, tag := range tags {
		fmt.Printf(" %v", color.GreenString(tag))
		for i := 0; i < maxLength-len(tag)+1; i++ {
			fmt.Print(" ")
		}
		if i+1 < len(tags) {
			fmt.Print(",")
		}

		if (i+1)%5 == 0 {
			fmt.Println("")
		}
	}
}

func printTagsRed(tags []string) {
	maxLength := 0

	for _, tag := range tags {
		if len(tag) > maxLength {
			maxLength = len(tag)
		}
	}

	for i, tag := range tags {
		fmt.Printf(" %v", color.RedString(tag))
		for i := 0; i < maxLength-len(tag)+1; i++ {
			fmt.Print(" ")
		}
		if i+1 < len(tags) {
			fmt.Print("|")
		}

		if (i+1)%5 == 0 {
			fmt.Println("")
		}
	}
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

func writeTable(path string, table interface{}) error {
	source, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer source.Close()

	enc := gob.NewEncoder(source)
	if err := enc.Encode(table); err != nil {
		return err
	}
	return nil
}

// this can be used to debug, to see if both tables are consistent
func checkConsistency(tpTable TagToPath, ptTable PathToTag) {
	for path, tags := range ptTable.Table {
		for tag := range tags {
			_, ok := tpTable.Table[tag][path]
			if !ok {
				panic("tables are inconsistent!!!")
			}
		}
	}

	fmt.Println("tables are consistent :)")
}

func ManageError(err error) {
	color.Red(err.Error())
	Hold()
	os.Exit(1)
}
