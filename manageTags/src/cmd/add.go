package cmd

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var input string

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

func AddTags(path string) error {
	prompt := promptui.Prompt{
		Label:     "Hey there!, please enter tags",
		Default:   " eg: tag1, tag2, tag3",
		AllowEdit: false,
		Validate: func(input string) error {
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	tags := strings.Split(result, ",")

	validTags := []string{}
	invalideTags := []string{}
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if isValidTag(tag) {
			validTags = append(validTags, tag)
		} else {
			invalideTags = append(invalideTags, tag)
		}
	}
	if len(invalideTags) > 0 {
		fmt.Println("\nonly " + color.YellowString("alphanumeric") + " characters are allowed")
		fmt.Println(color.RedString("Found invalid tags :(, following tags are not allowed:"))
		printTagsRed(invalideTags)
		fmt.Println("")
	}

	updateTable(validTags, path)
	fmt.Printf("\nSuccefully added following tags to %v: \n", path)
	printTagsGreen(validTags)
	fmt.Println("")
	return nil
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
			fmt.Print(",")
		}

		if (i+1)%5 == 0 {
			fmt.Println("")
		}
	}
}

func printTagsBlue(tags []string) {
	maxLength := 0

	for _, tag := range tags {
		if len(tag) > maxLength {
			maxLength = len(tag)
		}
	}

	for i, tag := range tags {
		fmt.Printf(" %v", color.BlueString(tag))
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

func updateTable(tags []string, path string) {
	ptTable, err := RecoverPathToTagTable(Pt)
	if err != nil {
		fmt.Scanln(&input)
		panic(err) 
	}
	
	tpTable, err := RecoverTagToPathTable(Tp)
	if err != nil {
		fmt.Scanln(&input)
		panic(err) 
	}
	
	// checkConsistency(tpTable, ptTable)

	// fmt.Println(ptTable, tpTable)
	
	_, ok := ptTable.Table[path]
	if ok {
		for _, tag := range tags {
			v, ok := ptTable.Table[path][tag]
			if !ok || v == false {
				ptTable.Table[path][tag] = true
				_, ok := tpTable.Table[tag]
				if !ok {
					tpTable.Table[tag] = make(map[string]bool, 0)
				}
				tpTable.Table[tag][path] = true
			}

		}
	} else {
		// first time we are tagging file
		ptTable.Table[path] = make(map[string]bool, 0)
		for _, tag := range tags {
			ptTable.Table[path][tag] = true

			_, ok := tpTable.Table[tag]
			if !ok {
				tpTable.Table[tag] = make(map[string]bool, 0)
			}
			tpTable.Table[tag][path] = true
		}
	}

	if err := writeTable(Tp, tpTable); err != nil {
		fmt.Scanln(&input)
		panic(err)
	}
	if err := writeTable(Pt, ptTable); err != nil {
		fmt.Scanln(&input)
		panic(err)
	}
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
