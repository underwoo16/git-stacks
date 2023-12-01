package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/metadata"
)

func (c *commandRunner) Stack(args []string) {
	parentBranch := c.gitService.GetCurrentBranch()
	parentRefSha := c.gitService.GetCurrentSha()

	if !c.metadataService.ConfigExists() {
		fmt.Println("No stacks found. Initializing...")
		c.metadataService.UpdateConfig(metadata.Config{Trunk: parentBranch})
	}

	stackName := stackNameFromArgs(args)
	if c.gitService.BranchExists(stackName) {
		fmt.Printf("Branch '%s' already exists\n", stackName)
		os.Exit(1)
	}

	fmt.Printf("Creating stack '%s'...\n", stackName)

	c.stackService.CreateStack(stackName, parentBranch, parentRefSha)

	fmt.Printf("Done! Switched to new stack '%s'\n", colors.CurrentStack(stackName))
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
