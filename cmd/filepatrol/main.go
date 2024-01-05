package main

import (
	"github.com/sifatulrabb/filepatrol"
	"github.com/sifatulrabb/filepatrol/cli"
)

func main() {
	rootPath, command := cli.ParseUserInput()
	watcher := filepatrol.NewWatchdog(rootPath)
	watcher.StartWatching(changeHandler(command))
}

func changeHandler(command string) func(files []string) {
	return func(files []string) {
		cli.CommandExecutor(command)
	}
}
