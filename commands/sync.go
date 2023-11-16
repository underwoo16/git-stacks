package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Sync() {
	// TODO: Handle merge failure (e.g. conflicts) and continue
	currentStack := stacks.GetGraph()
	stacks.ResyncChildren([]*stacks.StackNode{currentStack}, currentStack.ParentRefSha)
	git.CheckoutBranch(currentStack.Name)
}
