package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func Pr(args []string) {
	gitService := git.NewGitService()
	currentStack := stacks.GetCurrentStackNode()
	pullRequests := git.GetPullRequests()
	if len(args) < 1 {
		gitService.Rebase(currentStack.ParentBranch, currentStack.Name)
		submitPullRequestForStack(currentStack, pullRequests)

		gitService.CheckoutBranch(currentStack.Name)
		return
	}

	if args[0] == "all" {
		trunk := stacks.GetGraph()
		Resync(trunk)
		stacks.CacheGraphToDisk(trunk)

		submitAllPullRequests(trunk, pullRequests)

		gitService.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func submitAllPullRequests(trunk *stacks.StackNode, pullRequests []git.PullRequest) {
	prQueue := queue.New()
	prQueue.Push(trunk)

	for !prQueue.IsEmpty() {
		stack := prQueue.Pop().(*stacks.StackNode)

		for _, child := range stack.Children {
			prQueue.Push(child)
		}

		submitPullRequestForStack(stack, pullRequests)
	}
}

func submitPullRequestForStack(stack *stacks.StackNode, pullRequests []git.PullRequest) {
	if stack == nil {
		return
	}

	parent := stack.Parent
	if parent == nil {
		return
	}

	pullRequest := pullRequestFor(stack.Name, parent.Name, pullRequests)
	if pullRequest != nil {
		fmt.Printf("Pull request already exists for %s\n", stack.Name)
		fmt.Printf("View it here: %s\n", colors.Blue(pullRequest.Url))
		return
	}

	gitService := git.NewGitService()
	gitService.ForcePushBranch(stack.Name)

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
