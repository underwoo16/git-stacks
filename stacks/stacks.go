package stacks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/utils"
)

type StackNode struct {
	Name         string
	Parent       *StackNode
	Children     []*StackNode
	RefSha       string
	ParentBranch string
	ParentRefSha string
}

func readStack(ref string) *StackNode {
	out := git.Show(ref)
	items := strings.Fields(out)
	name := GetNameFromRef(ref)
	return &StackNode{
		Name:         name,
		RefSha:       git.RevParse(name),
		ParentBranch: GetNameFromRef(items[0]),
		ParentRefSha: items[1],
		Children:     []*StackNode{},
	}
}

func GetNameFromRef(ref string) string {
	ref = strings.Replace(ref, "refs/heads/", "", -1)
	return strings.Replace(ref, "refs/stacks/", "", -1)
}

func convertHeadToStack(ref string) string {
	return strings.Replace(ref, "refs/heads", "refs/stacks", -1)
}

func getStacks() []*StackNode {
	var existingStacks []*StackNode
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

func UpdateStack(stack *StackNode) {
	tempFilePath := fmt.Sprintf(".git/temp-%s", stack.Name)

	hashObject := fmt.Sprintf("%s\n%s", stack.ParentBranch, stack.ParentRefSha)
	utils.WriteToFile(tempFilePath, hashObject)

	objectSha := git.CreateHashObject(tempFilePath)
	utils.RemoveFile(tempFilePath)

	newRef := fmt.Sprintf("refs/stacks/%s", stack.Name)
	git.UpdateRef(newRef, objectSha)
}

func CreateStack(name string, parentBranch string, parentRefSha string) {
	tempFilePath := fmt.Sprintf(".git/temp-%s", name)

	hashObject := fmt.Sprintf("%s\n%s", parentBranch, parentRefSha)
	utils.WriteToFile(tempFilePath, hashObject)

	objectSha := git.CreateHashObject(tempFilePath)
	utils.RemoveFile(tempFilePath)

	newRef := fmt.Sprintf("refs/stacks/%s", name)
	git.UpdateRef(newRef, objectSha)
	git.CreateAndCheckoutBranch(name)

	// TODO:
	// Update stack graph
	// Update cache from stack graph
}

func StackExists(ref string) bool {
	name := GetNameFromRef(ref)
	// TODO: get .git directory dynamically to avoid hardcoding
	return utils.FileExists(fmt.Sprintf(".git/refs/stacks/%s", name))
}

func GetCurrentStackNode() *StackNode {
	currentBranch := git.GetCurrentRef()
	if !StackExists(currentBranch) {
		return nil
	}

	trunk := GetGraph()
	return findStack(trunk, GetNameFromRef(currentBranch))
}

func findStack(node *StackNode, branch string) *StackNode {
	if node.Name == branch {
		return node
	}

	for _, child := range node.Children {
		if found := findStack(child, branch); found != nil {
			return found
		}
	}

	return nil
}

func NeedsSync(stack *StackNode) bool {
	if stack.Parent == nil {
		return false
	}

	return stack.ParentRefSha != stack.RefSha
}

func RestackChildren(children []*StackNode, parentSha string) {
	for _, child := range children {
		if child.Parent == nil {
			RestackChildren(child.Children, child.RefSha)
			continue
		}

		childName := colors.CurrentStack(child.Name)
		parentName := colors.OtherStack(child.Parent.Name)
		if NeedsSync(child) {
			fmt.Printf("%s restacking onto %s\n", childName, parentName)
			git.Rebase(child.ParentBranch, child.Name)
			newSha := git.RevParse(child.Name)
			child.RefSha = newSha
			child.ParentRefSha = parentSha

			// TODO: is it better to update the cache here or at the end using full graph?
			UpdateStack(child)
		} else {
			fmt.Printf("%s up to date with %s\n", childName, parentName)
		}

		RestackChildren(child.Children, child.RefSha)
	}
}
