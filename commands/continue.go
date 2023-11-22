package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func Continue() {
	// TODO: check if continue info exists
	fmt.Println("Continuing sync...")
	continueInfo := stacks.GetContinueInfo()

	trunk := stacks.GetGraph()
	stackMap := make(map[string]*stacks.StackNode)
	populateMap(trunk, stackMap)

	fmt.Println("Continuing rebase...")
	git.RebaseContinue()

	continueBanch := continueInfo.ContinueBranch
	stack := stackMap[continueBanch]
	stack.RefSha = git.RevParse(continueBanch)
	stack.ParentRefSha = git.RevParse(stack.ParentBranch)

	fmt.Println("Syncing stacks...")
	branches := continueInfo.Branches

	if len(branches) == 0 {
		fmt.Println("Sync complete")
		return
	}

	syncQueue := queue.New()
	for _, branch := range branches {
		stack := stackMap[branch]
		syncQueue.Push(stack)
	}

	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		SyncStack(stack, syncQueue)
	}

	fmt.Println("Sync complete")
	stacks.CacheGraphToDisk(trunk)

	// TODO: Remove continue file
	// TODO: switch back to original branch
}

func populateMap(stack *stacks.StackNode, stackMap map[string]*stacks.StackNode) {
	stackMap[stack.Name] = stack
	for _, child := range stack.Children {
		populateMap(child, stackMap)
	}
}
