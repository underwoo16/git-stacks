package commands

import (
	"fmt"
	"log"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

type ContinueCommand struct {
	GitService      git.GitService
	MetadataService *stacks.MetadataService
	StackService    *stacks.StackService
}

func (c *ContinueCommand) Run() {
	if !c.MetadataService.ContinueInfoExists() {
		log.Fatal("No continue info found")
	}

	fmt.Println("Continuing sync...")
	continueInfo := c.MetadataService.GetContinueInfo()

	trunk := c.StackService.GetGraph()
	stackMap := make(map[string]*stacks.StackNode)
	populateMap(trunk, stackMap)

	fmt.Println("Continuing rebase...")
	c.GitService.RebaseContinue()

	continueBanch := continueInfo.ContinueBranch
	stack := stackMap[continueBanch]
	stack.RefSha = c.GitService.RevParse(continueBanch)
	stack.ParentRefSha = c.GitService.RevParse(stack.ParentBranch)

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

	syncCommand := SyncCommand{
		GitService:      c.GitService,
		MetadataService: c.MetadataService,
		StackService:    c.StackService,
	}
	for !syncQueue.IsEmpty() {
		stack := syncQueue.Pop().(*stacks.StackNode)
		for _, child := range stack.Children {
			syncQueue.Push(child)
		}

		syncCommand.SyncStack(stack, syncQueue)
	}

	fmt.Println("Sync complete")
	c.StackService.CacheGraphToDisk(trunk)

	c.MetadataService.RemoveContinueInfo()
}

func populateMap(stack *stacks.StackNode, stackMap map[string]*stacks.StackNode) {
	stackMap[stack.Name] = stack
	for _, child := range stack.Children {
		populateMap(child, stackMap)
	}
}
