package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

type DownCommand struct {
	GitService   git.GitService
	StackService *stacks.StackService
}

func (dc *DownCommand) Run() {
	currentNode := dc.StackService.GetCurrentStackNode()

	if currentNode == nil || currentNode.Parent == nil {
		fmt.Printf("%s\n", colors.CurrentStack(dc.GitService.GetCurrentBranch()))
		fmt.Printf("Already at bottom of stack.\n")
		return
	}

	parentBranch := currentNode.ParentBranch

	fmt.Printf("%s <- %s\n", parentBranch, colors.OtherStack(currentNode.Name))

	dc.GitService.CheckoutBranch(parentBranch)
	fmt.Printf("Switched to %s.\n", colors.CurrentStack(parentBranch))
}
