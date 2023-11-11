package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/commands"
)

func main() {
	allArgs := os.Args[1:]

	if len(allArgs) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	commandArgs := allArgs[1:]
	switch allArgs[0] {
	case "stack":
		commands.Stack(commandArgs)
	case "log":
		commands.Log()
	default:
		fmt.Println("Unknown command")
	}
}
