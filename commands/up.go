package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Up() {
	currentNode := stacks.GetCurrentStackNode()
	childNode := currentNode.Child

	if childNode == nil {
		fmt.Printf("%s\n", currentNode.Name)
		fmt.Printf("Already at top of stack.\n")
		return
	}

	fmt.Printf("\u21B1 %s\n", colors.Purple(childNode.Name))
	fmt.Printf("%s\n", currentNode.Name)

	git.CheckoutBranch(childNode.Name)
	fmt.Printf("Checked out %s.\n", colors.Blue(childNode.Name))
}
