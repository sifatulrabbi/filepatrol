package main

import (
	"github.com/sifatulrabb/filepatrol"
	"github.com/sifatulrabb/filepatrol/cli"
	"github.com/sifatulrabb/filepatrol/statichttp"
)

const STATIC_SRERVER_TYPE = "filepatrol.http"

func main() {
	execType, rootPath, command := cli.ParseUserInput()
	watcher := filepatrol.NewWatchdog(rootPath)

	if execType == STATIC_SRERVER_TYPE {
		watcher.StartWatching(statichttp.RunStaticHttpServer(rootPath))
	} else {
		watcher.StartWatching(changeHandler(command))
	}
}

func changeHandler(command string) func(files []string) {
	return func(files []string) {
		cli.CommandExecutor(command)
	}
}
