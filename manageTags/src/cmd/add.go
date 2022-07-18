package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func AddTags(path string) error {
	prompt := promptui.Prompt{
		Label:     "Hey there!, please enter tags",
		Default:   " eg: tag1, tag2, tag3",
		AllowEdit: false,
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
		fmt.Printf("\nOnly %v characters are allowed", color.YellowString("alphanumeric"))
		color.Red("\nFound invalid tags :(, following tags are not allowed:")
		printTagsRed(invalideTags)
		fmt.Println("")
	}

	updateTable(validTags, path)
	fmt.Printf("\nSuccessfully added following tags to %v: \n", path)
	printTagsGreen(validTags)
	fmt.Println("")
	return nil
}

func updateTable(tags []string, path string) error {
	ptTable, err := RecoverPathToTagTable(Pt)
	if err != nil {
		return err
	}

	tpTable, err := RecoverTagToPathTable(Tp)
	if err != nil {
		return err
	}

	_, ok := ptTable.Table[path]
	if ok {
		for _, tag := range tags {
			v, ok := ptTable.Table[path][tag]
			if !ok || !v {
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
		return err
	}
	if err := writeTable(Pt, ptTable); err != nil {
		return err
	}
	return nil
}
