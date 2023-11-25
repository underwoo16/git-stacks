package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

// TODO: use rerere
func Sync() {
	fmt.Printf("Syncing stacks...\n")

	gitService := git.NewGitService()
	currentBranch := gitService.GetCurrentBranch()
	trunk := stacks.GetGraph()
	Resync(trunk)
	stacks.CacheGraphToDisk(trunk)
	gitService.CheckoutBranch(currentBranch)
}

func Resync(trunk *stacks.StackNode) {
	syncQueue := queue.New()
	syncQueue.Push(trunk)

	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		SyncStack(stack, syncQueue)
	}
}

func SyncStack(stack *stacks.StackNode, syncQueue *queue.Queue) {
	if !stacks.NeedsSync(stack) {
		return
	}

	fmt.Printf("Rebasing %s onto %s\n", colors.CurrentStack(stack.Name), colors.OtherStack(stack.ParentBranch))

	gitService := git.NewGitService()
	err := gitService.Rebase(stack.ParentBranch, stack.Name)
	if err != nil {
		fmt.Printf("'%s' rebase failed\n", stack.Name)
		fmt.Printf("Resolve conflicts and run %s\n", colors.Yellow("git-stacks continue"))
		fmt.Printf("Alternatively, run %s to abort the rebase\n", colors.Yellow("git-stacks rebase --abort"))

		stacks.StoreContinueInfo(stack.Name, syncQueue)
		os.Exit(1)
	}

	// TODO: don't create new git service - use class reference once this is a class
	newParentSha := gitService.RevParse(stack.ParentBranch)
	newSha := gitService.RevParse(stack.Name)
	stack.RefSha = newSha
	stack.ParentRefSha = newParentSha
}
