package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/stacks"
)

type StackCommand struct {
	args            []string
	GitService      git.GitService
	MetadataService metadata.MetadataService
	StackService    stacks.StackService
}

func NewStackCommand(args []string, gitService git.GitService, metadataService metadata.MetadataService, stackService stacks.StackService) *StackCommand {
	return &StackCommand{
		args:            args,
		GitService:      gitService,
		MetadataService: metadataService,
		StackService:    stackService,
	}
}

func (s *StackCommand) Run() {
	parentBranch := s.GitService.GetCurrentBranch()
	parentRefSha := s.GitService.GetCurrentSha()

	if !s.MetadataService.ConfigExists() {
		fmt.Println("No stacks found. Initializing...")
		s.MetadataService.UpdateConfig(metadata.Config{Trunk: parentBranch})
	}

	stackName := stackNameFromArgs(s.args)
	if s.GitService.BranchExists(stackName) {
		fmt.Printf("Branch '%s' already exists\n", stackName)
		os.Exit(1)
	}

	fmt.Printf("Creating stack '%s'...\n", stackName)

	s.StackService.CreateStack(stackName, parentBranch, parentRefSha)

	fmt.Printf("Done! Switched to new stack '%s'\n", colors.CurrentStack(stackName))
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
