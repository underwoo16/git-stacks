package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

// TODO: check if branch is ahead before pushing
// TODO: check if branch is behind before pushing

func (c *commandRunner) Push(args []string) {
	currentStack := c.stackService.GetCurrentStackNode()
	if len(args) < 1 {

		c.stackService.SyncStack(currentStack, queue.New())

		fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
		c.gitService.ForcePushBranch(currentStack.Name)
		return
	}

	if args[0] == "all" {
		trunk := c.stackService.GetGraph()
		c.stackService.Resync(trunk)

		c.pushAllStacks(trunk)

		c.gitService.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func (c *commandRunner) pushAllStacks(trunk *stacks.StackNode) {
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
		c.gitService.ForcePushBranch(stack.Name)
	}
}
