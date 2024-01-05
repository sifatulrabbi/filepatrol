package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type File struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"time"`
}

type DirWatcher struct {
	Path  string `json:"path"`
	Files []File `json:"files"`
}

func NewDirWatcher(path string) DirWatcher {
	return DirWatcher{
		Path:  path,
		Files: []File{},
	}
}

var fileToIgnore = []string{
	"*.env.*", // needs supports for pattern matching.
	".secret",
}

func (w *DirWatcher) Watch(callback func()) error {
	// first make sure the directory exists
	dir, err := os.ReadDir(w.Path)
	if os.IsNotExist(err) {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			// get all the files of the dir and their last modified time.
			for i := 0; i < len(dir); i++ {
				entry := dir[i]
				fileInfo, err := entry.Info()
				if err != nil {
					// this means the file is changed, need to trigger the file change event
					fmt.Println(err)
					callback()
					continue
				}

				file := File{
					Name:         fileInfo.Name(),
					LastModified: fileInfo.ModTime(),
				}
				if w.isChanged(file) {
					callback()
				}
				w.addToFilesList(file)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	wg.Wait()
	return nil
}

func (w *DirWatcher) isChanged(f File) bool {
	for _, file := range w.Files {
		if file.Name != f.Name {
			continue
		}

		if f.LastModified.UnixMilli() != file.LastModified.UnixMilli() {
			return true
		}
	}

	return false
}

func (w *DirWatcher) addToFilesList(f File) {
	for i := 0; i < len(w.Files); i++ {
		if w.Files[i].Name == f.Name {
			w.Files[i] = f
			return
		}
	}
	w.Files = append(w.Files, f)
}

// this would determine whether the watcher should ignore the file or not.
func (w *DirWatcher) shouldIgnore(f File) {
	// TODO: need regex for patter matching.
}
