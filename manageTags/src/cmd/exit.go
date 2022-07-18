package cmd

import (
	"os"
	"time"

	"github.com/fatih/color"
)

func Exit(_ string) error {
	color.Green("...................THANK YOU FOR USING TAGGIT :)")
	time.Sleep(time.Second)
	os.Exit(1)
	return nil
}
