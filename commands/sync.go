package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

type SyncCommand struct {
	MetadataService *stacks.MetadataService
	StackService    *stacks.StackService
	GitService      git.GitService
}

func (sc *SyncCommand) Run() {
	fmt.Printf("Syncing stacks...\n")

	currentBranch := sc.GitService.GetCurrentBranch()
	trunk := sc.StackService.GetGraph()
	sc.Resync(trunk)
	sc.StackService.CacheGraphToDisk(trunk)
	sc.GitService.CheckoutBranch(currentBranch)
}

// TODO: move this method into StackService
func (sc *SyncCommand) Resync(trunk *stacks.StackNode) {
	syncQueue := queue.New()
	syncQueue.Push(trunk)

	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		sc.SyncStack(stack, syncQueue)
	}
}

// TODO: move this method into StackService
func (sc *SyncCommand) SyncStack(stack *stacks.StackNode, syncQueue *queue.Queue) {
	if !sc.StackService.NeedsSync(stack) {
		return
	}

	fmt.Printf("Rebasing %s onto %s\n", colors.CurrentStack(stack.Name), colors.OtherStack(stack.ParentBranch))

	err := sc.GitService.Rebase(stack.ParentBranch, stack.Name)
	if err != nil {
		fmt.Printf("'%s' rebase failed\n", stack.Name)
		fmt.Printf("Resolve conflicts and run %s\n", colors.Yellow("git-stacks continue"))
		fmt.Printf("Alternatively, run %s to abort the rebase\n", colors.Yellow("git-stacks rebase --abort"))

		sc.MetadataService.StoreContinueInfo(stack.Name, syncQueue)
		os.Exit(1)
	}

	newParentSha := sc.GitService.RevParse(stack.ParentBranch)
	newSha := sc.GitService.RevParse(stack.Name)
	stack.RefSha = newSha
	stack.ParentRefSha = newParentSha
}
