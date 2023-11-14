package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Restack() {
	// TODO: use cache if exists
	// TODO: get subtree starting from current branch
	trunk := stacks.BuildStackGraphFromScratch()
	restackChildren([]*stacks.StackNode{trunk}, trunk.ParentRefSha)

	// TODO: is there a way to rebase other branches without the side effect of switching to them?
	git.CheckoutBranch(trunk.Name)
}

func restackChildren(children []*stacks.StackNode, parentSha string) {
	for _, child := range children {
		if child.Parent == nil {
			restackChildren(child.Children, child.RefSha)
			continue
		}

		childName := colors.CurrentStack(child.Name)
		parentName := colors.OtherStack(child.Parent.Name)
		if child.ParentRefSha != parentSha {
			fmt.Printf("%s restacking onto %s\n", childName, parentName)
			git.Rebase(child.ParentBranch, child.Name)
			newSha := git.RevParse(child.Name)
			child.RefSha = newSha
		} else {
			fmt.Printf("%s up to date with %s\n", childName, parentName)
		}

		restackChildren(child.Children, child.RefSha)
	}
}
