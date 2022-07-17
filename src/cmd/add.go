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

type TagToPath struct {
	Table map[string]map[string]bool
}

func isValidTag(tag string) bool {
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
		Default:   "eg: tag1, tag2, tag3",
		AllowEdit: false,
		Validate: func(input string) error {
			return nil
		},
	}

	result, err := prompt.Run()
	fmt.Printf("entered values %v\n", result)
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
		fmt.Println("Invalid tags, which are not added:")
		for _, tag := range invalideTags {
			fmt.Println(color.RedString(tag))
		}
	}

	updateTable(validTags, path)
	return nil
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
			fmt.Printf("## %v\n", tag)
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
