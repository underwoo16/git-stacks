package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Down() {
	currentNode := stacks.GetCurrentStackNode()
	parentNode := currentNode.Parent

	if parentNode == nil {
		fmt.Printf("%s\n", colors.Purple(currentNode.Name))
		fmt.Printf("Already at bottom of stack.\n")
		return
	}

	fmt.Printf("%s\n", colors.Purple(currentNode.Name))
	fmt.Printf("\u2B91  %s\n", parentNode.Name)

	git.CheckoutBranch(parentNode.Name)
	fmt.Printf("Checked out %s.\n", colors.Blue(parentNode.Name))
}
