package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/stacks"
)

type Command interface {
	Run()
}

type CommandRunner interface {
	Run(args []string)
}

type commandRunner struct {
	gitService      git.GitService
	gitHubService   git.GitHubService
	stackService    stacks.StackService
	metadataService metadata.MetadataService
}

func NewCommandRunner(gitService git.GitService, gitHubService git.GitHubService, stackService stacks.StackService, metadataService metadata.MetadataService) *commandRunner {
	return &commandRunner{
		gitService:      gitService,
		gitHubService:   gitHubService,
		stackService:    stackService,
		metadataService: metadataService,
	}
}

func (c *commandRunner) Run(args []string) {
	var cmd Command
	switch args[0] {
	case "continue":
		cmd = NewContinueCommand(c.gitService, c.metadataService, c.stackService)
	case "stack":
		cmd = NewStackCommand(args[1:], c.gitService, c.metadataService, c.stackService)
	case "show":
		cmd = NewShowCommand(c.gitService, c.stackService)
	case "down":
		cmd = NewDownCommand(c.gitService, c.stackService)
	case "up":
		cmd = NewUpCommand(c.gitService, c.stackService)
	case "sync":
		cmd = NewSyncCommand(c.gitService, c.metadataService, c.stackService)
	case "pr":
		cmd = NewPrCommand(args[1:], c.gitService, c.gitHubService, c.metadataService, c.stackService)
	case "push-stack":
		cmd = NewPushCommand(args[1:], c.gitService, c.metadataService, c.stackService)
	default:
		cmd = NewPassThroughCommand(args, c.gitService)
	}

	cmd.Run()
}
