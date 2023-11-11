package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	currentStack := stacks.GetCurrentStack()
	currentStackName := stacks.GetNameFromRef(currentStack.Name)

	// print out graph
	stackLinkedList := stacks.GetStackList()
	current := &stackLinkedList

	for current != nil {
		if current.Name == currentStackName {
			fmt.Printf("* ")
		}

		fmt.Printf("%s", current.Name)

		if current.Parent != nil {
			fmt.Printf(" <- ")
		} else {
			fmt.Printf("\n")
		}
		current = current.Parent
	}
}
