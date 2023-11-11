package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	currentStack := stacks.GetCurrentStack()
	currentStackName := stacks.GetNameFromRef(currentStack.Name)

	stackLinkedList := stacks.GetStackList()
	current := &stackLinkedList

	for current != nil {
		if current.Name == currentStackName {
			fmt.Printf(colors.Blue("* "))
			fmt.Printf(colors.Blue(current.Name))
		} else {
			fmt.Printf(colors.Yellow(current.Name))
		}

		if current.Parent != nil {
			fmt.Printf(colors.White(" <- "))
		} else {
			fmt.Printf("\n")
		}
		current = current.Parent
	}
}
