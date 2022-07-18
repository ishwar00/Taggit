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
	for {
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
		if len(paths) > 0 {
			paths = append(paths, "Search Again", "Exit")
		listLoop:
			for {
				prompt := promptui.Select{
					Label: fmt.Sprintf("Found %v results :) ", len(paths)-2),
					Items: paths,
				}

				_, result, err := prompt.Run()
				if err != nil {
					return err
				}
				switch result {
				case "Exit":
					return Exit("")
				case "Search Again":
					break listLoop
				default:
					if err := open.Start(result); err != nil {
						return err
					}
				}
				clearScreen()
			}

		} else {
			fmt.Printf("Sorry, no files are tagged by %v yet :/ \n", color.YellowString(tag))
			prompt := promptui.Select{
				Label: "Please go with",
				Items: []string{
					"Search Again",
					"Exit",
				},
			}

			_, result, err := prompt.Run()
			if err != nil {
				return err
			}

			switch result {
			case "Exit":
				return Exit("")
			case "Search Again":
				// continue
				// do nothing
			default:
				return fmt.Errorf("sorry, something unexpected happened")
			}
		}
		// Hold()
		clearScreen()
	}
}
