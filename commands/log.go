package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	fmt.Println("Showing stack...")

	// print out graph
	stackLinkedList := stacks.GetStackList()
	current := &stackLinkedList

	for current != nil {
		fmt.Printf("%s", current.Name)

		if current.Parent != nil {
			fmt.Printf(" <- ")
		} else {
			fmt.Printf("\n")
		}
		current = current.Parent
	}
}
