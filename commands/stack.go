package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Stack(args []string) {
	parentBranch := git.GetCurrentBranch()
	parentRefSha := git.GetCurrentSha()

	if !stacks.ConfigExists() {
		fmt.Println("No stacks config found. Initializing...")
		stacks.UpdateConfig(stacks.Config{Trunk: parentBranch})
	}

	stackName := stackNameFromArgs(args)
	if git.BranchExists(stackName) {
		fmt.Printf("Branch '%s' already exists\n", stackName)
		os.Exit(1)
	}

	fmt.Printf("Creating stack '%s'...\n", stackName)

	stacks.CreateStack(stackName, parentBranch, parentRefSha)

	fmt.Printf("Done! Switched to new stack '%s'\n", stackName)
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
