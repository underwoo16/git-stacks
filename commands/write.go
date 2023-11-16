package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Write(args []string) {
	// TODO: Check if on a stack
	// TODO: Check if any changes to commit

	fmt.Println("Writing to stack...")
	currentStack := stacks.GetCurrentStackNode()

	refSha := git.RevParse(currentStack.Name)
	parentRefSha := git.RevParse(currentStack.ParentBranch)

	// TODO: check if -m passed to avoid vim
	if refSha != parentRefSha {
		git.CommitAmend()
	} else {
		git.Commit()
	}

	stacks.ResyncChildren([]*stacks.StackNode{currentStack}, parentRefSha)
	git.CheckoutBranch(currentStack.Name)
}
