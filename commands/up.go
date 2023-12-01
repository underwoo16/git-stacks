package commands

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/prompts"
)

func (c *commandRunner) Up() {
	currentNode := c.stackService.GetCurrentStackNode()
	if currentNode == nil {
		log.Fatal("Not on a known stack.")
	}

	children := currentNode.Children

	if len(children) == 0 {
		log.Fatalf("No stacks on top of %s\n", colors.CurrentStack(currentNode.Name))
	}

	if len(children) == 1 {
		c.switchToFrom(children[0].Name, currentNode.Name)
		return
	}

	branches := []string{}
	for _, child := range children {
		branches = append(branches, child.Name)
	}

	r := prompts.PromptUser(branches, "Select child branch", branchPromptTemplate())
	c.switchToFrom(r, currentNode.Name)
}

func (c *commandRunner) switchToFrom(to string, from string) {
	fmt.Printf("%s -> %s\n", colors.OtherStack(from), to)

	c.gitService.CheckoutBranch(to)
	fmt.Printf("Switched to %s.\n", colors.CurrentStack(to))
}

func branchPromptTemplate() *promptui.SelectTemplates {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ . | green }}",
		Inactive: "{{ . | yellow }}",
		Selected: "* {{ . | green }}",
	}

	return templates
}
