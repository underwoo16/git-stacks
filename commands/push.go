package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

// TODO: check if branch is ahead before pushing
// TODO: check if branch is behind before pushing

type PushCommand struct {
	GitService      git.GitService
	StackService    *stacks.StackService
	MetadataService *stacks.MetadataService
}

func (pc *PushCommand) Run(args []string) {
	currentStack := pc.StackService.GetCurrentStackNode()
	if len(args) < 1 {
		if pc.StackService.NeedsSync(currentStack) {
			fmt.Printf("Rebasing %s onto %s\n", colors.CurrentStack(currentStack.Name), colors.OtherStack(currentStack.ParentBranch))
			pc.GitService.Rebase(currentStack.ParentBranch, currentStack.Name)

			// TODO: consolidate this logic - it is repeated in several places
			newParentSha := pc.GitService.RevParse(currentStack.ParentBranch)
			newSha := pc.GitService.RevParse(currentStack.Name)
			currentStack.RefSha = newSha
			currentStack.ParentRefSha = newParentSha
			pc.StackService.CacheGraphToDisk(currentStack)

			fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
			pc.GitService.ForcePushBranch(currentStack.Name)
			return
		}

		fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
		pc.GitService.PushBranch(currentStack.Name)
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

		pc.pushAllStacks(trunk)

		pc.GitService.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func (pc *PushCommand) pushAllStacks(trunk *stacks.StackNode) {
	pushQueue := queue.New()
	pushQueue.Push(trunk)

	for !pushQueue.IsEmpty() {
		stack := pushQueue.Pop().(*stacks.StackNode)

		for _, child := range stack.Children {
			pushQueue.Push(child)
		}

		if stack.Parent == nil {
			continue
		}

		fmt.Printf("Pushing %s\n", colors.CurrentStack(stack.Name))
		pc.GitService.ForcePushBranch(stack.Name)
	}
}
