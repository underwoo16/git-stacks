package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Stack(args []string) {
	gitService := git.NewGitService()
	metadataService := stacks.NewMetadataService(gitService)
	parentBranch := gitService.GetCurrentBranch()
	parentRefSha := gitService.GetCurrentSha()

	if !metadataService.ConfigExists() {
		fmt.Println("No stacks found. Initializing...")
		metadataService.UpdateConfig(stacks.Config{Trunk: parentBranch})
	}

	stackName := stackNameFromArgs(args)
	if gitService.BranchExists(stackName) {
		fmt.Printf("Branch '%s' already exists\n", stackName)
		os.Exit(1)
	}

	fmt.Printf("Creating stack '%s'...\n", stackName)

	stacks.CreateStack(stackName, parentBranch, parentRefSha)

	fmt.Printf("Done! Switched to new stack '%s'\n", colors.CurrentStack(stackName))
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
