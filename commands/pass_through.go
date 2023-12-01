package commands

import "github.com/underwoo16/git-stacks/git"

type passThroughCommand struct {
	args       []string
	GitService git.GitService
}

func NewPassThroughCommand(args []string, gitService git.GitService) *passThroughCommand {
	return &passThroughCommand{
		args:       args,
		GitService: gitService,
	}
}

func (p *passThroughCommand) Run() {
	p.GitService.PassThrough(p.args)
}
