package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func RemoveTags(path string) error {

	prompt := promptui.Prompt{
		Label:     "please enter tags which are to be removed",
		Default:   " eg: friends, birthday, pets",
		AllowEdit: false,
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	tags := strings.Split(result, ",")
	validTags := []string{}
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if isValidTag(tag) {
			validTags = append(validTags, tag)
		}
	}

	tpTable, err := RecoverTagToPathTable(Tp)
	if err != nil {
		return err
	}

	ptTable, err := RecoverPathToTagTable(Pt)
	if err != nil {
		return err

	}

	deletedTags := []string{}
	unknownTags := []string{}

	_, ok := ptTable.Table[path]
	if ok {
		for _, tag := range validTags {
			_, ok := ptTable.Table[path][tag]
			if ok {
				delete(ptTable.Table[path], tag)
				delete(tpTable.Table[tag], path)
				deletedTags = append(deletedTags, tag)
			} else {
				unknownTags = append(unknownTags, tag)
			}
		}
	}

	if len(deletedTags) == 0 {
		fmt.Println("No valid tags to remove :(")
	}

	if len(deletedTags) > 0 {
		fmt.Println("Deleted following tags successfully :)")
		printTagsGreen(deletedTags)
		fmt.Println("")
	}

	if len(unknownTags) > 0 {
		fmt.Println("Following tags were not available to take action:")
		printTagsRed(unknownTags)
		fmt.Println("")
	}

	if err := writeTable(Pt, ptTable); err != nil {
		return err
	}

	if err := writeTable(Tp, tpTable); err != nil {
		return err
	}

	return nil
}
