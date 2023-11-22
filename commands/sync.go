package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func Sync() {
	// TODO: Handle merge failure (e.g. conflicts) and continue
	fmt.Printf("Syncing stacks...\n")
	currentBranch := git.GetCurrentBranch()
	trunk := stacks.GetGraph()
	resync(trunk)
	stacks.CacheGraphToDisk(trunk)
	git.CheckoutBranch(currentBranch)
}

func resync(trunk *stacks.StackNode) {
	syncQueue := queue.New()
	syncQueue.Push(trunk)

	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		if !stacks.NeedsSync(stack) {
			continue
		}

		fmt.Printf("Syncing %s onto %s\n", stack.Name, stack.ParentBranch)

		err := git.Rebase(stack.ParentBranch, stack.Name)
		if err != nil {
			fmt.Printf("'%s' rebase failed\n", stack.Name)
			fmt.Printf("Resolve conflicts and run %s\n", colors.Yellow("git-stacks continue"))
			fmt.Printf("Alternatively, run %s to abort the rebase\n", colors.Yellow("git-stacks rebase --abort"))

			// TODO Store continue info in a file
			// should be the branch we are currently rebasing
			// and the syncQueue
		}

		newParentSha := git.RevParse(stack.ParentBranch)
		newSha := git.RevParse(stack.Name)
		stack.RefSha = newSha
		stack.ParentRefSha = newParentSha
	}
}
