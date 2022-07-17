package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var Tp string

func init() {
	buf, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	Tp = filepath.Join(buf, "Taggit/tagToPath.db")
	createDB(Tp)
}

func Execute() error {
	prompt := promptui.Prompt{
		Label: "Please enter tag to search:",
	}

	tag, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	tag = strings.TrimSpace(tag)
	if !isValidTag(tag) {
		return fmt.Errorf(color.RedString("Please enter a valid Tag :("))
	}

	paths, err := SearchByTags(tag)
	if err != nil {
		return err
	}

	if len(paths) > 0 {
		prompt := promptui.Select{
			Label: fmt.Sprintf("Found %v results :) ", len(paths)),
			Items: paths,
		}

		_, selectedPath, err := prompt.Run()

		if err != nil {
			return err
		}

		fmt.Println("selected file: ", color.GreenString(selectedPath))
	} else {
		fmt.Printf("Sorry, no files are tagged by %v yet :(\n", color.YellowString(tag))
	}

	return nil
}
