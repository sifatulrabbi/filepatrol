package main

import (
	"github.com/sifatulrabbi/filepatrol/internals/cli"
	"github.com/sifatulrabbi/filepatrol/internals/statichttp"
	"github.com/sifatulrabbi/filepatrol/internals/watchdog"
)

func main() {
	execType, rootPath, command := cli.ParseUserInput()
	watcher := watchdog.NewWatchdog(rootPath)

	if execType == statichttp.STATIC_SRERVER_TYPE {
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
