package main

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/commands"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/stacks"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	gitService := git.NewGitService()
	gitHubService := git.NewGitHubService()
	metadataService := metadata.NewMetadataService(gitService)
	stackService := stacks.NewStackService(gitService, metadataService)

	switch args[0] {
	case "continue":
		cmd := commands.ContinueCommand{
			GitService:      gitService,
			MetadataService: metadataService,
			StackService:    stackService,
		}
		cmd.Run()
	case "stack":
		cmd := commands.StackCommand{
			GitService:      gitService,
			MetadataService: metadataService,
			StackService:    stackService,
		}
		cmd.Run(args[1:])
	case "show":
		cmd := commands.ShowCommand{
			GitService:   gitService,
			StackService: stackService,
		}
		cmd.Run()
	case "down":
		cmd := commands.DownCommand{
			GitService:   gitService,
			StackService: stackService,
		}
		cmd.Run()
	case "up":
		cmd := commands.UpCommand{
			GitService:   gitService,
			StackService: stackService,
		}
		cmd.Run()
	case "sync":
		cmd := commands.SyncCommand{
			GitService:      gitService,
			MetadataService: metadataService,
			StackService:    stackService,
		}
		cmd.Run()
	case "pr":
		cmd := commands.PrCommand{
			GitService:      gitService,
			MetadataService: metadataService,
			StackService:    stackService,
			GitHubService:   gitHubService,
		}
		cmd.Run(args[1:])
	case "push-stack":
		cmd := commands.PushCommand{
			GitService:      gitService,
			MetadataService: metadataService,
			StackService:    stackService,
		}
		cmd.Run(args[1:])
	default:
		cmd := commands.PassThroughCommand{
			GitService: gitService,
		}
		cmd.Run(args)
	}
}
