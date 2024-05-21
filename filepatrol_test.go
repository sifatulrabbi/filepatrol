package filepatrol

import (
	"fmt"
	"testing"
)

func TestCollectingDirItems(t *testing.T) {
	watcher := NewWatchdog("../..")
	if len(watcher.Objects) < 1 {
		t.Error("the watch dog is unable to pick up dir entries")
	}
}

func TestWatchingFiles(t *testing.T) {
	watcher := NewWatchdog("../..")
	go watcher.StartWatching(func(files []string) {
		fmt.Println("files changed", files)
	})

	if len(watcher.Objects) < 1 {
		t.Error("the watch dog is unable to pick up dir entries")
	}
}
