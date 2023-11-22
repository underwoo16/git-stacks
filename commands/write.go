package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

// TODO: fix this method
func Write(args []string) {
	// TODO: Check if on a stack
	// TODO: Check if any changes to commit

	fmt.Println("Writing to stack...")
	currentStack := stacks.GetCurrentStackNode()

	// TODO: check if -m passed to avoid vim
	// TODO: we don't need to amend, lets just add new commits
	git.Commit()

	// stacks.ResyncChildren([]*stacks.StackNode{currentStack}, parentRefSha)
	git.CheckoutBranch(currentStack.Name)
}
