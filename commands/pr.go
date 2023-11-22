package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func Pr(args []string) {
	fmt.Println("pr")

	// TODO: consider existing pull requests before submitting new ones

	currentStack := stacks.GetCurrentStackNode()
	if len(args) < 1 {
		submitPullRequestForStack(currentStack)
	}

	if args[0] == "all" {
		// TODO: submit pull requests for all stacks
		submitAllPullRequests(currentStack)
	}
}

func submitAllPullRequests(stack *stacks.StackNode) {
	// TODO: resync all stacks

	prQueue := queue.New()
	prQueue.Push(stack)

	for !prQueue.IsEmpty() {
		stack := prQueue.Pop().(*stacks.StackNode)

		submitPullRequestForStack(stack)

		for _, child := range stack.Children {
			prQueue.Push(child)
		}
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

	// TODO: make this optional in case we already synced the stack
	git.Rebase(parent.Name, stack.Name)

	git.ForcePushBranch(stack.Name)

	git.CreatePullRequest(parent.Name, stack.Name)
}
