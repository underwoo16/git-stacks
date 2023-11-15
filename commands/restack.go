package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Restack() {
	// TODO: Handle merge failure (e.g. conflicts) and continue
	currentStack := stacks.GetGraph()
	stacks.RestackChildren([]*stacks.StackNode{currentStack}, currentStack.ParentRefSha)
	git.CheckoutBranch(currentStack.Name)
}
