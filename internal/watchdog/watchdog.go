package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type Watchdog struct {
	RootPath string
	Objects  []WatchObject
}

func recursDir(objects *[]WatchObject, rootPath string, dir []fs.DirEntry, i int) *[]WatchObject {
	if len(dir) < 1 {
		return objects
	}

	item, err := dir[i].Info()
	if item == nil || err != nil {
		log.Println("Error occurred:", err)
		return objects
	}

	itemFullPath := fmt.Sprintf("%s/%s", rootPath, item.Name())
	if item.IsDir() {
		nestedDir, _ := os.ReadDir(itemFullPath)
		objects = recursDir(objects, itemFullPath, nestedDir, 0)
	} else {
		*objects = append(*objects, WatchObject{
			Name:        itemFullPath,
			LastModTime: item.ModTime(),
		})
	}

	if i+1 >= len(dir) {
		return objects
	}
	return recursDir(objects, rootPath, dir, i+1)
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

	objects := recursDir(&[]WatchObject{}, rootPath, dir, 0)
	for _, obj := range *objects {
		fmt.Printf("%s\n", obj.Name)
	}

	wdog := Watchdog{
		RootPath: rootPath,
		Objects:  *objects,
	}
	return &wdog
}

func (wdog *Watchdog) StartWatching(bark func(files []string)) error {
	return nil
}
