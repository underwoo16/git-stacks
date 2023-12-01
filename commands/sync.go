package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
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
	s.StackService.Resync(trunk)
	s.GitService.CheckoutBranch(currentBranch)
}
