package watchdog

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectingDirItems(t *testing.T) {
	watcher := NewWatchdog("../..")
	if len(watcher.Objects) < 1 {
		t.Error("the watch dog is unable to pick up dir entries")
	}
}

func TestWatchingFiles(t *testing.T) {
	watcher := NewWatchdog("../../tmp")
	assert.Greater(t, len(watcher.Objects), 0, "the watch dog is unable to pick up dir entries")
	for k := range watcher.Objects {
		fmt.Println(k)
	}
	watcher.StartWatching(func(files []string) {
		fmt.Println("files changed", files)
	})
}

func TestIgnoreFilenameMatching(t *testing.T) {
	ignoreList := []string{"*.log", "build/*", ".cache"}
	assert.Equal(t, true, shouldIgnore(ignoreList, "test.log"), "matching 'test.log'")
	assert.Equal(t, true, shouldIgnore(ignoreList, "./build"), "matching './build'")
	assert.Equal(t, true, shouldIgnore(ignoreList, "./build/app"), "matching './build/app'")
	assert.Equal(t, true, shouldIgnore(ignoreList, "app/build/app"), "matching 'app/build/app'")
	assert.Equal(t, true, shouldIgnore(ignoreList, "./app/.cache/somefile"), "matching './app/.cache/somefile'")
	assert.Equal(t, true, shouldIgnore(ignoreList, ".cache/somefile"), "matching '.cache/somefile'")
}

func TestFileTreeTraversingAccuracy(t *testing.T) {
	rootPath := "../../tmp"
	_, err := os.Stat(rootPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Panicln("the root path is invalid please provide a valid path for patrolling!", err)
		} else {
			log.Panicln(err)
		}
	}
	dir, err := os.ReadDir(rootPath)
	if err != nil {
		log.Panicln(err)
	}
	objects := sniffDirObjects(&WatchObjects{}, rootPath, dir, 0, ignoredFileList(rootPath))
	for k := range *objects {
		fmt.Println(k)
	}
}
