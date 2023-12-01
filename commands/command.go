package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/stacks"
)

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
	switch args[0] {
	case "continue":
		c.Continue()
	case "stack":
		c.Stack(args[1:])
	case "show":
		c.Show()
	case "down":
		c.Down()
	case "up":
		c.Up()
	case "sync":
		c.Sync()
	case "pr":
		c.PullRequest(args[1:])
	case "push-stack":
		c.Push(args[1:])
	default:
		c.PassThrough(args)
	}
}
