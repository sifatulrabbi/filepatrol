package watchdog

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

type WatchObjects map[string]time.Time

type Watchdog struct {
	RootPath string
	Objects  WatchObjects
}

func sniffDirObjects(objects *WatchObjects, rootPath string, dir []fs.DirEntry, i int) *WatchObjects {
	if len(dir) < 1 {
		return objects
	}

	item, err := dir[i].Info()
	if item == nil || err != nil {
		log.Println("Difficulty when patrolling due to:", err)
		return objects
	}

	itemFullPath := fmt.Sprintf("%s/%s", rootPath, item.Name())
	if item.IsDir() {
		nestedDir, _ := os.ReadDir(itemFullPath)
		objects = sniffDirObjects(objects, itemFullPath, nestedDir, 0)
	} else {
		(*objects)[itemFullPath] = item.ModTime()
	}

	if i+1 >= len(dir) {
		return objects
	}
	return sniffDirObjects(objects, rootPath, dir, i+1)
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

	objects := sniffDirObjects(&WatchObjects{}, rootPath, dir, 0)
	wdog := Watchdog{
		RootPath: rootPath,
		Objects:  *objects,
	}
	return &wdog
}

func (wdog *Watchdog) StartWatching(bark func(files []string)) {
	fmt.Println("Watch dog has started patrolling files...")

	for {
		changedFiles := []string{}
		dir, err := os.ReadDir(wdog.RootPath)
		if err != nil {
			if os.IsNotExist(err) {
				panic("the root path is invalid please provide a valid path for patrolling!")
			} else {
				panic(err)
			}
		}

		objects := sniffDirObjects(&WatchObjects{}, wdog.RootPath, dir, 0)
		for obj, ts := range *objects {
			prev := wdog.Objects[obj]
			if prev.UnixMilli() < ts.UnixMilli() {
				changedFiles = append(changedFiles, obj)
			}
		}
		wdog.Objects = *objects

		if len(changedFiles) > 0 {
			bark(changedFiles)
		}

		time.Sleep(time.Second * 2)
	}
}
