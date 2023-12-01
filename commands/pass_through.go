package commands

func (c *commandRunner) PassThrough(args []string) {
	c.gitService.PassThrough(args)
}
