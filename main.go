package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/commands"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	switch args[0] {
	case "stack":
		commands.Stack(args[1:])
	case "log":
		commands.Log()
	case "down":
		commands.Down()
	case "up":
		commands.Up()
	case "restack":
		commands.Restack()
	case "write":
		commands.Write(args[1:])
	default:
		commands.PassThrough(args)
	}
}
