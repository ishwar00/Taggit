package main

import (
	"fmt"

	"github.com/ishwar00/Taggit/searchTags/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		errMsg := fmt.Errorf("sorry, encountered an error: %w", err)
		fmt.Println(errMsg)
	}
	fmt.Println("Please press ENTER to exit...")
	var input string
	fmt.Scanln(&input)
}
