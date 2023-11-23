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
	// TODO: Handle merge failure (e.g. conflicts) and continue
	fmt.Printf("Syncing stacks...\n")
	currentBranch := git.GetCurrentBranch()
	trunk := stacks.GetGraph()
	Resync(trunk)
	stacks.CacheGraphToDisk(trunk)
	git.CheckoutBranch(currentBranch)
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

	fmt.Printf("Syncing %s onto %s\n", stack.Name, stack.ParentBranch)

	err := git.Rebase(stack.ParentBranch, stack.Name)
	if err != nil {
		fmt.Printf("'%s' rebase failed\n", stack.Name)
		fmt.Printf("Resolve conflicts and run %s\n", colors.Yellow("git-stacks continue"))
		fmt.Printf("Alternatively, run %s to abort the rebase\n", colors.Yellow("git-stacks rebase --abort"))

		stacks.StoreContinueInfo(stack.Name, syncQueue)
		os.Exit(1)
	}

	newParentSha := git.RevParse(stack.ParentBranch)
	newSha := git.RevParse(stack.Name)
	stack.RefSha = newSha
	stack.ParentRefSha = newParentSha
}
