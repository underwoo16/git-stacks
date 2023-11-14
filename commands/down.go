package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Down() {
	currentNode := stacks.GetCurrentStackNode()

	if currentNode == nil {
		fmt.Printf("%s\n", colors.CurrentStack(git.GetCurrentBranch()))
		fmt.Printf("Already at bottom of stack.\n")
		return
	}

	parentBranch := currentNode.ParentBranch

	fmt.Printf("%s\n", colors.OtherStack(currentNode.Name))
	fmt.Printf("\u2B91  %s\n", parentBranch)

	git.CheckoutBranch(parentBranch)
	fmt.Printf("Checked out %s.\n", colors.CurrentStack(parentBranch))
}
