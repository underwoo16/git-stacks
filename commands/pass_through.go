package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
)

func PassThrough(args []string) {
	fmt.Printf(colors.Gray("Passing command through to Git...\n"))
	git.PassThrough(args)
}
