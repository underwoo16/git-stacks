package stacks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/utils"
)

type Stack struct {
	Name            string
	ParentBranchRef string
	ParentRefSha    string
}

type StackNode struct {
	Value  Stack
	Parent *StackNode
	Child  *StackNode
	Name   string
}

func GetCurrentStack() Stack {
	currentRef := git.GetCurrentRef()

	if refIsStack(currentRef) {
		currentStackRef := convertHeadToStack(currentRef)
		return readStack(currentStackRef)
	} else {
		return Stack{
			Name: currentRef,
		}
	}

}

func GetCurrentStackNode() StackNode {
	currentStack := GetCurrentStack()
	stackList := GetStackList()

	current := &stackList
	for current != nil {
		if current.Name == GetNameFromRef(currentStack.Name) {
			return *current
		}

		current = current.Parent
	}

	log.Fatal("Not currently on any known stack")
	return stackList
}

func GetStackList() StackNode {
	currentStacks := getStacks()

	m := make(map[string]StackNode)
	for _, stack := range currentStacks {
		node := StackNode{
			Value: stack,
			Name:  GetNameFromRef(stack.Name),
		}
		m[stack.Name] = node
	}

	tipStack := FindTip(currentStacks)
	tipNode := m[tipStack.Name]

	currentNode := &tipNode
	for currentNode != nil {
		currentStack := currentNode.Value

		parentStackNode, parentExists := m[convertHeadToStack(currentStack.ParentBranchRef)]
		if !parentExists {
			trunkNode := StackNode{
				Name: GetNameFromRef(currentStack.ParentBranchRef),
			}
			parentStackNode = trunkNode
		}

		parentStackNode.Child = currentNode

		currentNode.Parent = &parentStackNode
		if parentExists {
			currentNode = &parentStackNode
		} else {
			currentNode = nil
		}
	}

	return tipNode
}

func readStack(ref string) Stack {
	out := git.Show(ref)
	items := strings.Fields(out)

	return Stack{Name: ref, ParentBranchRef: items[0], ParentRefSha: items[1]}
}

func GetNameFromRef(ref string) string {
	ref = strings.Replace(ref, "refs/heads/", "", -1)
	return strings.Replace(ref, "refs/stacks/", "", -1)
}

func convertHeadToStack(ref string) string {
	return strings.Replace(ref, "refs/heads", "refs/stacks", -1)
}

func getStacks() []Stack {
	var existingStacks []Stack
	// TODO: get .git directory dynamically to avoid hardcoding
	err := filepath.Walk(".git/refs/stacks", func(path string, info os.FileInfo, err error) error {
		utils.CheckError(err)

		if info.IsDir() {
			return nil
		}

		ref := path[5:]
		stack := readStack(ref)
		existingStacks = append(existingStacks, stack)

		return nil
	})
	utils.CheckError(err)

	return existingStacks
}

func refIsStack(ref string) bool {
	index := strings.Index(ref, "refs/stacks/")
	return index == 0
}

func FindTip(stacks []Stack) Stack {
	if len(stacks) == 0 {
		fmt.Println("No stacks initialized.")
		os.Exit(1)
	}

	tip := stacks[0]
	for _, stack := range stacks {
		hasChild := false
		for _, child := range stacks {
			parentStackName := convertHeadToStack(child.ParentBranchRef)
			if parentStackName == stack.Name {
				hasChild = true
			}
		}

		if !hasChild {
			tip = stack
			break
		}
	}

	return tip
}
