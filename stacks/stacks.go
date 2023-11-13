package stacks

import (
	"fmt"
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
	Value    Stack
	Parent   *StackNode
	Children []*StackNode
	Name     string
}

func readStack(ref string) Stack {
	out := git.Show(ref)
	items := strings.Fields(out)
	return Stack{Name: GetNameFromRef(ref), ParentBranchRef: items[0], ParentRefSha: items[1]}
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

		// TODO: this seems brittle
		ref := path[5:]
		stack := readStack(ref)
		existingStacks = append(existingStacks, stack)

		return nil
	})
	utils.CheckError(err)

	return existingStacks
}

func BuildStackGraphFromScratch() StackNode {
	allStacks := getStacks()

	config := GetConfig()
	trunk := StackNode{
		Name: config.Trunk,
	}
	BuildGraphRecursive(&trunk, allStacks)

	return trunk
}

func BuildGraphRecursive(trunk *StackNode, allStacks []Stack) {
	for _, stack := range allStacks {
		if stack.ParentBranchRef == trunk.Name {
			childNode := StackNode{
				Value:    stack,
				Name:     stack.Name,
				Parent:   trunk,
				Children: []*StackNode{},
			}
			trunk.Children = append(trunk.Children, &childNode)
		}
	}

	if len(trunk.Children) == 0 {
		return
	}

	for _, child := range trunk.Children {
		BuildGraphRecursive(child, allStacks)
	}
}

func InsertStack(name string, parentBranchRef string, parentRefSha string) {
	// perform git operations to create and switch to stack
	tempFilePath := fmt.Sprintf(".git/temp-%s", name)

	parentBranch := GetNameFromRef(parentBranchRef)
	hashObject := fmt.Sprintf("%s\n%s", parentBranch, parentRefSha)
	utils.WriteToFile(tempFilePath, hashObject)

	objectSha := git.CreateHashObject(tempFilePath)
	utils.RemoveFile(tempFilePath)
	newRef := fmt.Sprintf("refs/stacks/%s", name)
	git.UpdateRef(newRef, objectSha)

	git.CreateAndCheckoutBranch(name)

	// Update stack graph
	stackToInsert := Stack{
		Name:            name,
		ParentBranchRef: parentBranchRef,
		ParentRefSha:    parentRefSha,
	}
	fmt.Println(stackToInsert)

	// Update stack graph
	// Update cache from stack graph
}
