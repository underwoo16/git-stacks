package commands

import (
	"fmt"
)

func (c *commandRunner) Sync() {
	fmt.Printf("Syncing stacks...\n")

	currentBranch := c.gitService.GetCurrentBranch()
	trunk := c.stackService.GetGraph()
	c.stackService.Resync(trunk)
	c.gitService.CheckoutBranch(currentBranch)
}
