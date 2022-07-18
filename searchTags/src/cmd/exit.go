package cmd

import (
	"time"

	"github.com/fatih/color"
)

func Exit(_ string) error {
	clearScreen()
	color.Green("\n\n\n...................THANK YOU FOR USING TAGGIT :)\n\n\n")
	time.Sleep(time.Second)
	return nil
}
