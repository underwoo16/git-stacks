package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Sync() {
	// TODO: Handle merge failure (e.g. conflicts) and continue
	// TODO: sync from current node instead of trunk?
	fmt.Printf("Syncing...\n")
	trunk := stacks.GetGraph()
	fmt.Printf("Syncing %s...\n", trunk.Name)
	fmt.Printf("Children: %s\n", trunk.Children)
	stacks.ResyncChildren([]*stacks.StackNode{trunk}, trunk.ParentRefSha)
	git.CheckoutBranch(trunk.Name)
	stacks.CacheGraphToDisk(trunk)
}
