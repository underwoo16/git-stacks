package stacks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func BuildStackGraphFromScratch() *StackNode {
	allStacks := getStacks()

	config := GetConfig()
	trunk := StackNode{
		Name:     config.Trunk,
		RefSha:   git.RevParse(config.Trunk),
		Children: []*StackNode{},
	}

	fmt.Printf("Building graph from trunk: %s\n", trunk.Name)
	BuildGraphRecursive(&trunk, allStacks)

	return &trunk
}

func BuildGraphRecursive(trunk *StackNode, allStacks []*StackNode) {
	for _, stack := range allStacks {
		if stack.ParentBranch == trunk.Name {
			stack.Parent = trunk
			trunk.Children = append(trunk.Children, stack)
		}
	}

	if len(trunk.Children) == 0 {
		return
	}

	for _, child := range trunk.Children {
		BuildGraphRecursive(child, allStacks)
	}
}

func CacheGraphToDisk(trunk *StackNode) {
	branches := bfs(trunk, []Branch{})
	UpdateCache(Cache{Branches: branches})

}

func GetGraphFromCache() *StackNode {
	branches := GetCache().Branches
	stackMap := make(map[string]*StackNode)

	for _, branch := range branches {
		node := StackNode{
			Name:         branch.Name,
			RefSha:       branch.BranchRevision,
			ParentBranch: branch.ParentBranchName,
			ParentRefSha: branch.ParentBranchRevision,
			Children:     []*StackNode{},
		}
		stackMap[branch.Name] = &node
	}

	for _, branch := range branches {
		node := stackMap[branch.Name]
		for _, child := range branch.Children {
			node.Children = append(node.Children, stackMap[child])
		}
	}

	trunkName := GetConfig().Trunk
	trunk := stackMap[trunkName]

	return trunk
}

func bfs(node *StackNode, arr []Branch) []Branch {
	branch := Branch{
		Name:                 node.Name,
		BranchRevision:       node.RefSha,
		ParentBranchName:     node.ParentBranch,
		ParentBranchRevision: node.ParentRefSha,
		Children:             []string{},
	}
	for i := 0; i < len(node.Children); i++ {
		child := node.Children[i]
		branch.Children = append(branch.Children, child.Name)
		arr = bfs(child, arr)
	}

	arr = append(arr, branch)
	return arr
}

func InsertStack(name string, parentBranch string, parentRefSha string) {
	// perform git operations to create and switch to stack
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

	currentStack := convertHeadToStack(currentBranch)
	return readStack(currentStack)
}
