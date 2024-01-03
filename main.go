package main

import (
	"fmt"
	"log"
)

// "github.com/fsnotify/fsnotify"

func main() {
	watcher := NewDirWatcher("./tmp")
	if err := watcher.Watch(func() { fmt.Println("Dir changed") }); err != nil {
		log.Fatalln(err)
	}
}
