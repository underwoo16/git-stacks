package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/underwoo16/git-stacks/stacks"
	"github.com/underwoo16/git-stacks/utils"
)

func Log() {
	fmt.Println("Showing stack...")

	// traverse refs/stacks
	currentStacks := getStacks()
	for _, stack := range currentStacks {
		fmt.Println(stack.Name)
	}

	// build linked list

	// mark current stack

	// print out graph
}

func getStacks() []stacks.Stack {
	var existingStacks []stacks.Stack
	err := filepath.Walk(".git/refs/stacks", func(path string, info os.FileInfo, err error) error {
		utils.CheckError(err)

		if info.IsDir() {
			return nil
		}

		ref := path[5:]
		stack := stacks.ReadStack(ref)
		existingStacks = append(existingStacks, stack)

		return nil
	})
	utils.CheckError(err)

	return existingStacks
}
