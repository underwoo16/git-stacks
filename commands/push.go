package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/stacks"
)

// TODO: check if branch is ahead before pushing
// TODO: check if branch is behind before pushing

type PushCommand struct {
	args            []string
	GitService      git.GitService
	MetadataService metadata.MetadataService
	StackService    stacks.StackService
}

func NewPushCommand(args []string, gitService git.GitService, metadataService metadata.MetadataService, stackService stacks.StackService) *PushCommand {
	return &PushCommand{
		args:            args,
		GitService:      gitService,
		MetadataService: metadataService,
		StackService:    stackService,
	}
}

func (p *PushCommand) Run() {
	currentStack := p.StackService.GetCurrentStackNode()
	if len(p.args) < 1 {

		p.StackService.SyncStack(currentStack, queue.New())

		fmt.Printf("Pushing %s\n", colors.CurrentStack(currentStack.Name))
		p.GitService.ForcePushBranch(currentStack.Name)
		return
	}

	if p.args[0] == "all" {
		trunk := p.StackService.GetGraph()
		p.StackService.Resync(trunk)

		p.pushAllStacks(trunk)

		p.GitService.CheckoutBranch(currentStack.Name)
		return
	}

	// TODO: add help message
	fmt.Println("Invalid arguments")
}

func (p *PushCommand) pushAllStacks(trunk *stacks.StackNode) {
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
		p.GitService.ForcePushBranch(stack.Name)
	}
}
