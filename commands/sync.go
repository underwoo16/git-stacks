package commands

import (
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Sync() {
	// TODO: Handle merge failure (e.g. conflicts) and continue
	// TODO: sync from current node instead of trunk?
	trunk := stacks.GetGraph()
	stacks.ResyncChildren([]*stacks.StackNode{trunk}, trunk.ParentRefSha)
	git.CheckoutBranch(trunk.Name)
	stacks.CacheGraphToDisk(trunk)
}
