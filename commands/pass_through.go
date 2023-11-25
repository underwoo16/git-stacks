package commands

import (
	"github.com/underwoo16/git-stacks/git"
)

func PassThrough(args []string) {
	gitService := git.NewGitService()
	gitService.PassThrough(args)
}
