package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

type PrCommand struct {
	GitService      git.GitService
	StackService    *stacks.StackService
	GitHubService   *git.GitHubService
	MetadataService *stacks.MetadataService
}

func (pc *PrCommand) Run(args []string) {
	pullRequests := pc.GitHubService.GetPullRequests()
	currentStack := pc.StackService.GetCurrentStackNode()
	if len(args) < 1 {
		pc.GitService.Rebase(currentStack.ParentBranch, currentStack.Name)
		pc.submitPullRequestForStack(currentStack, pullRequests)

		pc.GitService.CheckoutBranch(currentStack.Name)
		return
	}

	if args[0] == "all" {
		trunk := pc.StackService.GetGraph()

		syncCommand := &SyncCommand{
			MetadataService: pc.MetadataService,
			StackService:    pc.StackService,
			GitService:      pc.GitService,
		}

		syncCommand.Resync(trunk)
		pc.StackService.CacheGraphToDisk(trunk)

		pc.submitAllPullRequests(trunk, pullRequests)

		pc.GitService.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func (pc *PrCommand) submitAllPullRequests(trunk *stacks.StackNode, pullRequests []git.PullRequest) {
	prQueue := queue.New()
	prQueue.Push(trunk)

	for !prQueue.IsEmpty() {
		stack := prQueue.Pop().(*stacks.StackNode)

		for _, child := range stack.Children {
			prQueue.Push(child)
		}

		pc.submitPullRequestForStack(stack, pullRequests)
	}
}

func (pc *PrCommand) submitPullRequestForStack(stack *stacks.StackNode, pullRequests []git.PullRequest) {
	if stack == nil {
		return
	}

	parent := stack.Parent
	if parent == nil {
		return
	}

	pullRequest := pc.pullRequestFor(stack.Name, parent.Name, pullRequests)
	if pullRequest != nil {
		fmt.Printf("Pull request already exists for %s\n", stack.Name)
		fmt.Printf("View it here: %s\n", colors.Blue(pullRequest.Url))
		return
	}

	pc.GitService.ForcePushBranch(stack.Name)

	pc.GitHubService.CreatePullRequest(parent.Name, stack.Name)
}

func (pc *PrCommand) pullRequestFor(head string, base string, pulls []git.PullRequest) *git.PullRequest {
	for _, pull := range pulls {
		if pull.HeadRefName == head && pull.BaseRefName == base {
			return &pull
		}
	}
	return nil
}
