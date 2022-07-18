package cmd

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"unicode"
	"runtime"

	"github.com/fatih/color"
)

type PathToTag struct {
	Table map[string]map[string]bool
}

func (pt *PathToTag) Print() {
	for path, tags := range pt.Table {
		fmt.Printf("%v: %v\n", path, tags)
	}
}

type TagToPath struct {
	Table map[string]map[string]bool
}

func (tp *TagToPath) print() {
	for tag, paths := range tp.Table {
		fmt.Printf("%v: %v\n", tag, paths)
	}
}

func RecoverPathToTagTable(path string) (PathToTag, error) {
	source, err := os.Open(path)
	if err != nil {
		return PathToTag{}, err
	}
	defer source.Close()
	dec := gob.NewDecoder(source)

	var table PathToTag
	if err := dec.Decode(&table); err == io.EOF {
		table = PathToTag{
			Table: map[string]map[string]bool{},
		}
	} else if err != nil {
		return PathToTag{}, err
	}
	return table, nil
}

func RecoverTagToPathTable(path string) (TagToPath, error) {
	source, err := os.Open(path)
	if err != nil {
		return TagToPath{}, err
	}
	defer source.Close()
	dec := gob.NewDecoder(source)

	var table TagToPath
	if err := dec.Decode(&table); err == io.EOF {
		table = TagToPath{
			Table: map[string]map[string]bool{},
		}
	} else if err != nil {
		return TagToPath{}, err
	}
	return table, nil
}

func createDB(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		directory := filepath.Dir(path)
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			panic(err)
		}

		if _, err := os.Create(path); err != nil {
			panic(err)
		}

	} else if err != nil {
		panic(err)
	}
}

func isValidTag(tag string) bool {

	if len(tag) == 0 {
		return false
	}
	for _, c := range tag {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func Hold() {
	var input string
	fmt.Println("\nPlease press ENTER to exit...")
	fmt.Scanln(&input)
}
func ManageError(err error) {
	color.Red(err.Error())
	Hold()
	os.Exit(1)
}

func runCmd(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func clearScreen() {
	switch runtime.GOOS {
    case "darwin":
        runCmd("clear")
    case "linux":
        runCmd("clear")
    case "windows":
        runCmd("cmd", "/c", "cls")
    default:
        runCmd("clear")
    }
}
