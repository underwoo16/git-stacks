package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Stack(args []string) {
	parentBranchRef := git.GetCurrentRef()
	parentRefSha := git.GetCurrentSha()

	if !stacks.ConfigExists() {
		fmt.Println("No stacks config found. Creating...")
		parentBranch := stacks.GetNameFromRef(parentBranchRef)
		stacks.UpdateConfig(stacks.Config{Trunk: parentBranch})
	}

	stackName := stackNameFromArgs(args)
	fmt.Printf("Creating stack '%s'...\n", stackName)

	stacks.InsertStack(stackName, parentBranchRef, parentRefSha)

	fmt.Printf("Done! Switched to new stack '%s'\n", stackName)
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
