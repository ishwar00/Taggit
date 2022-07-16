package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Execute(path string) error {
	fmt.Printf("file name: %v\n", path)
	prompt := promptui.Select{
		Label: "Please select operation",
		Items: []string{
			"Add tags",
			"show tags",
			"remove tags",
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
	}

	if err := operations[i](path); err != nil {
		return err
	}

	return nil
}
