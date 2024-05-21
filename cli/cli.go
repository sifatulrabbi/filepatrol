package cli

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/sifatulrabb/filepatrol/statichttp"
)

// returns (execType, rootPath, command)
func ParseUserInput() (string, string, string) {
	rootPath := flag.String("path", ".", "Path to the root directory. Default is '.'")
	cmd := flag.String("cmd", "", "Enter the command you want to execute on file change.")
	execType := flag.String("exec", "command", "Enter the executioner type. Default is 'command'. Use 'filepatrol.http' for static server")

	flag.Parse()
	fmt.Printf("execType = %s, rootPath = %s, cmd = %s\n", *execType, *rootPath, *cmd)

	if *execType != statichttp.STATIC_SRERVER_TYPE && *cmd == "" {
		log.Fatalln("Please provide '--cmd'")
	}

	return *execType, *rootPath, *cmd
}

func CommandExecutor(command string) {
	const (
		ColorRed   = "\033[31m"
		ColorGreen = "\033[32m"
		ColorBlue  = "\033[34m"
		ColorReset = "\033[0m"
	)

	fmt.Printf("%s", ColorBlue)
	fmt.Println("changes detected")
	fmt.Printf("executing:`%s`\n", command)
	fmt.Printf("%s", ColorReset)

	cmd := exec.Command("bash", "-c", command)
	if output, err := cmd.Output(); err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
