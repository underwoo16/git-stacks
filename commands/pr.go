package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func Pr(args []string) {
	fmt.Println("pr")

	currentStack := stacks.GetCurrentStackNode()
	pullRequests := git.GetPullRequests()
	if len(args) < 1 {
		git.Rebase(currentStack.ParentBranch, currentStack.Name)
		submitPullRequestForStack(currentStack, pullRequests)
	}

	if args[0] == "all" {
		submitAllPullRequests(pullRequests)
	}
}

func submitAllPullRequests(pullRequests []git.PullRequest) {
	trunk := stacks.GetGraph()
	Resync(trunk)

	prQueue := queue.New()
	prQueue.Push(trunk)

	for !prQueue.IsEmpty() {
		stack := prQueue.Pop().(*stacks.StackNode)

		submitPullRequestForStack(stack, pullRequests)

		for _, child := range stack.Children {
			prQueue.Push(child)
		}
	}
}

func submitPullRequestForStack(stack *stacks.StackNode, pullRequests []git.PullRequest) {
	if stack == nil {
		fmt.Println("No stack found")
		os.Exit(1)
	}

	parent := stack.Parent
	if parent == nil {
		fmt.Println("No parent branch found")
		os.Exit(1)
	}

	pullRequest := pullRequestFor(stack.Name, parent.Name, pullRequests)
	if pullRequest != nil {
		fmt.Printf("Pull request already exists for %s\n", stack.Name)
		fmt.Printf("View it here: %s\n", colors.Blue(pullRequest.Url))
		return
	}

	git.ForcePushBranch(stack.Name)

	git.CreatePullRequest(parent.Name, stack.Name)
}

func pullRequestFor(head string, base string, pulls []git.PullRequest) *git.PullRequest {
	for _, pull := range pulls {
		if pull.HeadRefName == head && pull.BaseRefName == base {
			return &pull
		}
	}
	return nil

}
