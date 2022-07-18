package main

import "github.com/ishwar00/Taggit/theWatcher/watcher"

func main() {
	if err := watcher.Execute(); err != nil {
		watcher.ManageError(err)
	}
}
