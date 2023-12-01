package commands

import (
	"fmt"
	"log"

	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

func (c *commandRunner) Continue() {
	if !c.metadataService.ContinueInfoExists() {
		log.Fatal("No continue info found")
	}

	fmt.Println("Continuing sync...")
	continueInfo := c.metadataService.GetContinueInfo()

	trunk := c.stackService.GetGraph()
	stackMap := make(map[string]*stacks.StackNode)
	populateMap(trunk, stackMap)

	fmt.Println("Continuing rebase...")
	c.gitService.RebaseContinue()

	continueBanch := continueInfo.ContinueBranch
	stack := stackMap[continueBanch]
	stack.RefSha = c.gitService.RevParse(continueBanch)
	stack.ParentRefSha = c.gitService.RevParse(stack.ParentBranch)

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

		c.stackService.SyncStack(stack, syncQueue)
	}

	fmt.Println("Sync complete")
	c.stackService.CacheGraphToDisk(trunk)

	c.metadataService.RemoveContinueInfo()
}

func populateMap(stack *stacks.StackNode, stackMap map[string]*stacks.StackNode) {
	stackMap[stack.Name] = stack
	for _, child := range stack.Children {
		populateMap(child, stackMap)
	}
}
