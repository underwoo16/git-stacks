package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/commands"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/stacks"
	"github.com/underwoo16/git-stacks/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	fileService := utils.NewFileService()
	gitService := git.NewGitService()
	gitHubService := git.NewGitHubService()
	metadataService := metadata.NewMetadataService(gitService, fileService)
	stackService := stacks.NewStackService(gitService, metadataService, fileService)

	cmdRunner := commands.NewCommandRunner(gitService, gitHubService, stackService, metadataService)
	cmdRunner.Run(args)
}
