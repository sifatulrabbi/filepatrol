package watchdog

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type WatchObjects map[string]time.Time

type Watchdog struct {
	RootPath string
	Objects  WatchObjects
	Timeout  time.Duration
	Snapshot any
}

func NewWatchdog(rootPath string) *Watchdog {
	dir, err := os.ReadDir(rootPath)
	if err != nil {
		if os.IsNotExist(err) {
			panic("the root path is invalid please provide a valid path for patrolling!")
		} else {
			panic(err)
		}
	}

	objects := sniffDirObjects(&WatchObjects{}, rootPath, dir, 0, ignoredFileList(rootPath))
	wdog := Watchdog{
		RootPath: rootPath,
		Objects:  *objects,
		Timeout:  time.Second * 1,
	}
	return &wdog
}

func (wdog *Watchdog) StartWatching(bark func(files []string)) {
	fmt.Println("Watch dog has started patrolling files...")

	for {
		info, err := os.Stat(wdog.RootPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Panicln("the root path is invalid please provide a valid path for patrolling!", err)
			} else {
				log.Panicln(err)
			}
		}

		if !info.IsDir() {
			time.Sleep(time.Second * 2)
			if t, ok := wdog.Objects[wdog.RootPath]; ok && t.Before(time.Now()) {
				bark([]string{wdog.RootPath})
			}
			wdog.Objects[wdog.RootPath] = info.ModTime()
			continue
		}

		dir, err := os.ReadDir(wdog.RootPath)
		if err != nil {
			log.Panicln(err)
		}
		objects := sniffDirObjects(&WatchObjects{}, wdog.RootPath, dir, 0, ignoredFileList(wdog.RootPath))
		changedFiles := []string{}
		for k, ts := range *objects {
			prev := wdog.Objects[k]
			if prev.Before(ts) {
				changedFiles = append(changedFiles, k)
			}
		}
		wdog.Objects = *objects

		if len(changedFiles) > 0 {
			bark(changedFiles)
		}

		time.Sleep(time.Second * 2)
	}
}

func sniffDirObjects(objects *WatchObjects, rootPath string, dir []fs.DirEntry, i int, ignoredFiles []string) *WatchObjects {
	if len(dir) < 1 || i >= len(dir) {
		return objects
	}

	item, err := dir[i].Info()
	if item == nil || err != nil {
		log.Println("Difficulty when patrolling due to:", err)
		return objects
	}

	itemFullPath := fmt.Sprintf("%s/%s", rootPath, item.Name())
	if shouldIgnore(ignoredFiles, itemFullPath) {
		return objects
	}
	if item.IsDir() {
		nestedDir, err := os.ReadDir(itemFullPath)
		if err != nil {
			log.Printf("Difficulty when patrolling dir '%s' error: %s", itemFullPath, err)
			return objects
		}
		objects = sniffDirObjects(objects, itemFullPath, nestedDir, 0, ignoredFileList(itemFullPath))
	} else {
		(*objects)[itemFullPath] = item.ModTime()
	}

	return sniffDirObjects(objects, rootPath, dir, i+1, ignoredFiles)
}

func ignoredFileList(dirPath string) []string {
	filenames := []string{}
	gitignorePath := fmt.Sprintf("%s/.gitignore", dirPath)
	f, err := os.ReadFile(gitignorePath)
	if err != nil {
		return filenames
	}
	filenames = strings.Split(string(f), "\n")
	return filenames
}

func shouldIgnore(ignoredFiles []string, filename string) bool {
	isIncluded := false
	for i := 0; i < len(ignoredFiles); i++ {
		p := ignoredFiles[i]
		ignoredDir, ignoredFile := path.Split(p)
		targetDir, targetFile := path.Split(filename)

		if matched, _ := path.Match(ignoredDir, targetDir); matched {
			isIncluded = true
			break
		}
		if matched, _ := path.Match(ignoredFile, targetFile); matched {
			isIncluded = true
			break
		}
	}
	return isIncluded
}
