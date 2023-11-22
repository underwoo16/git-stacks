package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Pr(args []string) {
	fmt.Println("pr")

	currentStack := stacks.GetCurrentStackNode()
	if len(args) < 1 {
		submitPullRequestForStack(currentStack)
	}
}

func submitPullRequestForStack(stack *stacks.StackNode) {
	if stack == nil {
		fmt.Println("No stack found")
		os.Exit(1)
	}

	parent := stack.Parent
	if parent == nil {
		fmt.Println("No parent branch found")
		os.Exit(1)
	}

	git.Rebase(parent.Name, stack.Name)

	git.ForcePushBranch(stack.Name)

	git.CreatePullRequest(parent.Name, stack.Name)
}
