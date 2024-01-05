package main

import (
	"fmt"

	"github.com/sifatulrabb/filepatrol"
)

func main() {
	watcher := filepatrol.NewWatchdog("tmp")
	watcher.StartWatching(changeHandler)
}

func changeHandler(files []string) {
	fmt.Println("files changed:", files)
}
