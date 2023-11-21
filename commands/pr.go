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

	// TODO: can store pull requests in map for easy lookup
	pullRequests := git.GetPullRequests()
	submitPullRequests(currentStack, pullRequests)

	// TODO: should we cache PR info?
}

func submitPullRequests(stackNode *stacks.StackNode, existingPullRequests []git.PullRequest) {
	if stackNode == nil {
		return
	}

	fmt.Printf("Pushing branches...\n")

	// TODO: Don't push trunk
	// TODO: Don't push branches that are already pushed - can store in map
	git.ForcePushBranch(stackNode.ParentBranch)
	git.ForcePushBranch(stackNode.Name)

	pullRequest := pullRequestFor(stackNode.ParentBranch, stackNode.Name, existingPullRequests)
	if pullRequest != nil {
		fmt.Printf("PR already exists for %s <- %s\n", stackNode.ParentBranch, stackNode.Name)
		fmt.Printf("%s\n", pullRequest.Url)
	} else {
		fmt.Printf("Submitting PR for %s <- %s\n", stackNode.ParentBranch, stackNode.Name)
		git.CreatePullRequest(stackNode.ParentBranch, stackNode.Name)
	}

	fmt.Println()

	for _, child := range stackNode.Children {
		submitPullRequests(child, existingPullRequests)
	}
}

// TODO: can store pull requests in map for easy lookup
func pullRequestFor(base string, head string, pulls []git.PullRequest) *git.PullRequest {
	for _, pull := range pulls {
		if pull.BaseRefName == base && pull.HeadRefName == head {
			return &pull
		}
	}
	return nil
}

func update() {
	fmt.Println("update")
}
