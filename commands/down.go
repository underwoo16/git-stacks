package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
)

func (c *commandRunner) Down() {
	currentNode := c.stackService.GetCurrentStackNode()

	if currentNode == nil || currentNode.Parent == nil {
		fmt.Printf("%s\n", colors.CurrentStack(c.gitService.GetCurrentBranch()))
		fmt.Printf("Already at bottom of stack.\n")
		return
	}

	parentBranch := currentNode.ParentBranch

	fmt.Printf("%s <- %s\n", parentBranch, colors.OtherStack(currentNode.Name))

	c.gitService.CheckoutBranch(parentBranch)
	fmt.Printf("Switched to %s.\n", colors.CurrentStack(parentBranch))
}
