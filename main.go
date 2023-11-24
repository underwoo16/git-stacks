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
	case "continue":
		commands.Continue()
	case "stack":
		commands.Stack(args[1:])
	case "show":
		commands.Show()
	case "down":
		commands.Down()
	case "up":
		commands.Up()
	case "sync":
		commands.Sync()
	case "pr":
		commands.Pr(args[1:])
	case "push-stack":
		commands.Push(args[1:])
	default:
		commands.PassThrough(args)
	}
}
