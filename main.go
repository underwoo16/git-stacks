package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/flapjacks/commands"
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
	default:
		fmt.Println("Unknown command")
	}
}
