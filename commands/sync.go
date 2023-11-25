package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

type SyncCommand struct {
	MetadataService metadata.MetadataService
	StackService    stacks.StackService
	GitService      git.GitService
}

func (s *SyncCommand) Run() {
	fmt.Printf("Syncing stacks...\n")

	currentBranch := s.GitService.GetCurrentBranch()
	trunk := s.StackService.GetGraph()
	s.Resync(trunk)
	s.StackService.CacheGraphToDisk(trunk)
	s.GitService.CheckoutBranch(currentBranch)
}

// TODO: move this method into StackService
func (s *SyncCommand) Resync(trunk *stacks.StackNode) {
	syncQueue := queue.New()
	syncQueue.Push(trunk)

	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		s.SyncStack(stack, syncQueue)
	}
}

// TODO: move this method into StackService
func (s *SyncCommand) SyncStack(stack *stacks.StackNode, syncQueue *queue.Queue) {
	if !s.StackService.NeedsSync(stack) {
		return
	}

	fmt.Printf("Rebasing %s onto %s\n", colors.CurrentStack(stack.Name), colors.OtherStack(stack.ParentBranch))

	err := s.GitService.Rebase(stack.ParentBranch, stack.Name)
	if err != nil {
		fmt.Printf("'%s' rebase failed\n", stack.Name)
		fmt.Printf("Resolve conflicts and run %s\n", colors.Yellow("git-stacks continue"))
		fmt.Printf("Alternatively, run %s to abort the rebase\n", colors.Yellow("git-stacks rebase --abort"))

		var branches []string
		for !syncQueue.IsEmpty() {
			branches = append(branches, syncQueue.Pop().(*stacks.StackNode).Name)
		}
		s.MetadataService.StoreContinueInfo(stack.Name, branches)
		os.Exit(1)
	}

	newParentSha := s.GitService.RevParse(stack.ParentBranch)
	newSha := s.GitService.RevParse(stack.Name)
	stack.RefSha = newSha
	stack.ParentRefSha = newParentSha
}
