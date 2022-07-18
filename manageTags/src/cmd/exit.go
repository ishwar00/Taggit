package cmd

import (
	"time"
	"fmt"

	"github.com/fatih/color"
)

func Exit(_ string) error {
	clearScreen()
	color.Green("\n\n\n...................THANK YOU FOR USING TAGGIT :)\n\n\n")
	time.Sleep(time.Second)
	return fmt.Errorf("closing")
}
