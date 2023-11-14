package stacks

import "github.com/underwoo16/git-stacks/git"

func BuildStackGraphFromScratch() *StackNode {
	allStacks := getStacks()

	config := GetConfig()
	trunk := StackNode{
		Name:     config.Trunk,
		RefSha:   git.RevParse(config.Trunk),
		Children: []*StackNode{},
	}

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

func CacheGraphToDisk(trunk *StackNode) {
	branches := bfs(trunk, []Branch{})
	UpdateCache(Cache{Branches: branches})

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
