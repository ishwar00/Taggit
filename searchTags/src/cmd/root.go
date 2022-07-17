package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/skratchdot/open-golang/open"
)

var Tp string

func init() {
	buf, err := os.UserHomeDir()
	if err != nil {
		ManageError(err)
	}
	Tp = filepath.Join(buf, "Taggit/tagToPath.db")
	createDB(Tp)
}

func Execute() error {
	prompt := promptui.Prompt{
		Label: "Please enter tag to search",
	}

	tag, err := prompt.Run()
	if err != nil {
		ManageError(err)
	}

	tag = strings.TrimSpace(tag)
	if !isValidTag(tag) {
		return fmt.Errorf(color.RedString("Please enter a valid Tag :("))
	}

	paths, err := SearchByTags(tag)
	if err != nil {
		return err
	}
	paths = append(paths, "Exit")
	if len(paths) > 0 {
		prompt := promptui.Select{
			Label: fmt.Sprintf("Found %v results :) ", len(paths) - 1),
			Items: paths,
		}

		index, _, err := prompt.Run()

		if err != nil {
			return err
		}
		if index == len(paths)-1 {
			return nil
		}
		if err := open.Start(paths[index]); err != nil {
			return err
		}

	} else {
		fmt.Printf("Sorry, no files are tagged by %v yet :(\n", color.YellowString(tag))
	}

	return nil
}
