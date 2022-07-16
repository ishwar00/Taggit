package cmd

import "fmt"

func ShowTags(path string) error {

	ptTable, err := RecoverPathToTagTable(Pt)
	if err != nil {
		return err
	}

	tags := []string{}
	maxLength := 0
	for tag := range ptTable.Table[path] {
		tags = append(tags, tag)
		if maxLength < len(tag) {
			maxLength = len(tag)
		}
	}

	for i, tag := range tags {
		fmt.Printf(" %v", tag)
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

	return nil
}
