package internal

import (
	"testing"
)

func TestCollectingDirItems(t *testing.T) {
	watcher := NewWatchdog("../..")
	if len(watcher.Objects) < 1 {
		t.Error("the watch dog is unable to pick up dir entries")
	}
}
