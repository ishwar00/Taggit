package main

import (
	"fmt"

	"github.com/ishwar00/Taggit/searchTags/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		errMsg := fmt.Errorf("sorry, encountered an error: %w", err)
		cmd.ManageError(errMsg)
	}
	// cmd.Hold()
}
