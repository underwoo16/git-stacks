package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Restack() {
	// TODO: move this logic to stacks package
	// TODO: Handle merge failure (e.g. conflicts) and continue
	// TODO: get subtree starting from current branch
	trunk := stacks.GetGraph()
	stacks.RestackChildren([]*stacks.StackNode{trunk}, trunk.ParentRefSha)

	// TODO: is there a way to rebase other branches without the side effect of switching to them?
	git.CheckoutBranch(trunk.Name)
}
