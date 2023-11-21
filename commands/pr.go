package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

func Pr(args []string) {
	fmt.Println("pr")

	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	switch args[0] {
	case "submit":
		submit()
	case "update":
		update()
	default:
		fmt.Println("Unknown command provided")
	}
}

func submit() {
	fmt.Println("submit")

	// TODO: add progress logging
	currentStack := stacks.GetCurrentStackNode()
	stacks.ResyncChildren([]*stacks.StackNode{currentStack}, currentStack.ParentRefSha)
	submitPullRequests(currentStack)
}

func submitPullRequests(stackNode *stacks.StackNode) {
	if stackNode == nil {
		return
	}

	fmt.Printf("Pushing branches...\n")

	// TODO: Don't push trunk
	git.ForcePushBranch(stackNode.ParentBranch)
	git.ForcePushBranch(stackNode.Name)

	// TODO: If PR exists, don't submit
	fmt.Printf("Submitting PR for %s <- %s\n", stackNode.ParentBranch, stackNode.Name)

	git.CreatePullRequest(stackNode.ParentBranch, stackNode.Name)

	for _, child := range stackNode.Children {
		submitPullRequests(child)
	}
}

func update() {
	fmt.Println("update")
}
