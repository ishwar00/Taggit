package watcher

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

var Pt string
var Tp string

func init() {
	buf, err := os.UserHomeDir()
	if err != nil {
		ManageError(err)
	}
	Pt = filepath.Join(buf, "Taggit/pathToTag.db")
	Tp = filepath.Join(buf, "Taggit/tagToPath.db")
	createDB(Pt)
	createDB(Tp)
}

func createDB(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		directory := filepath.Dir(path)
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			ManageError(err)
		}

		if _, err := os.Create(path); err != nil {
			ManageError(err)
		}
	} else if err != nil {
		ManageError(err)
	}
}
func Execute() error {
	windows.FindFirstChangeNotification(
		"/home/ishwar/stuff and stuff/Taggit/theWatcher/src/watcher/f2.txt",
		false,
		windows.FILE_NOTIFY_CHANGE_FILE_NAME,
	)
	return nil
}
