package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Restack() {
	// TODO: use cache if exists
	// TODO: get subtree starting from current branch
	trunk := stacks.BuildStackGraphFromScratch()
	restackChildren([]*stacks.StackNode{trunk}, trunk.RefSha)
}

func restackChildren(children []*stacks.StackNode, parentSha string) {
	for _, child := range children {
		if child.ParentRefSha != parentSha {
			fmt.Printf("%s onto %s\n", child.Name, child.ParentBranch)
		} else {
			fmt.Printf("%s already up to date with %s\n", child.Name, child.ParentBranch)
		}

		// TODO: update refSha if rebased
		restackChildren(child.Children, child.RefSha)
	}
}
