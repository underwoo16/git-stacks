package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/commands"
	"github.com/underwoo16/git-stacks/git"
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
	case "write":
		commands.Write(args[1:])
	case "test":
		// fmt.Println(git.LogBetween("first", "second"))
		err := git.Rebase("first", "second")
		if err != nil {
			fmt.Println("Need to resolve conflicts and then run 'git-stacks continue'")
		}
	default:
		commands.PassThrough(args)
	}
}
