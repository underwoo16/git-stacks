package commands

import "github.com/underwoo16/git-stacks/git"

type PassThroughCommand struct {
	GitService git.GitService
}

func (p *PassThroughCommand) Run(args []string) {
	p.GitService.PassThrough(args)
}
