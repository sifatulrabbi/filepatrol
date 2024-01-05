package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// returns (rootPath, command)
func ParseUserInput() (string, string) {
	args := os.Args[1:]
	// sample command
	// filepatrol ./tmp echo "hello"
	if len(args) < 3 {
		panic(errors.New("your command if not properly formatted"))
	}

	rootPath := args[0]
	command := strings.Join(args[1:], " ")

	return rootPath, command
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
