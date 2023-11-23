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
func Push(args []string) {
	currentStack := stacks.GetCurrentStackNode()
	if len(args) < 1 {
		if stacks.NeedsSync(currentStack) {
			fmt.Printf("Rebasing %s onto %s\n", colors.CurrentStack(currentStack.Name), colors.OtherStack(currentStack.ParentBranch))
			git.Rebase(currentStack.ParentBranch, currentStack.Name)

			// TODO: consolidate this logic - it is repeated in several places
			newParentSha := git.RevParse(currentStack.ParentBranch)
			newSha := git.RevParse(currentStack.Name)
			currentStack.RefSha = newSha
			currentStack.ParentRefSha = newParentSha
			stacks.CacheGraphToDisk(currentStack)

			fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
			git.ForcePushBranch(currentStack.Name)
			return
		}

		fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
		git.PushBranch(currentStack.Name)
		return
	}

	if args[0] == "all" {
		trunk := stacks.GetGraph()
		Resync(trunk)
		stacks.CacheGraphToDisk(trunk)

		pushAllStacks(trunk)

		git.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func pushAllStacks(trunk *stacks.StackNode) {
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
		git.ForcePushBranch(stack.Name)
	}
}
