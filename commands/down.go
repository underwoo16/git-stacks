package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Down() {
	currentNode := stacks.GetCurrentStackNode()

	if currentNode == nil || currentNode.Parent == nil {
		fmt.Printf("%s\n", colors.CurrentStack(git.GetCurrentBranch()))
		fmt.Printf("Already at bottom of stack.\n")
		return
	}

	parentBranch := currentNode.ParentBranch

	fmt.Printf("%s <- %s\n", parentBranch, colors.OtherStack(currentNode.Name))

	git.CheckoutBranch(parentBranch)
	fmt.Printf("Switched to %s.\n", colors.CurrentStack(parentBranch))
}
