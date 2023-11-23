package stacks

import "github.com/underwoo16/git-stacks/git"

func GetGraph() *StackNode {
	if CacheExists() {
		return GetGraphFromCache()
	}

	graph := GetGraphFromRefs()
	CacheGraphToDisk(graph)
	return graph
}

func GetGraphFromRefs() *StackNode {
	config := GetConfig()
	trunk := StackNode{
		Name:     config.Trunk,
		RefSha:   git.RevParse(config.Trunk),
		Children: []*StackNode{},
	}

	allStacks := getStacks()
	allStacks = append(allStacks, &trunk)
	return BuildGraphIterative(&trunk, allStacks)
}

func BuildGraphIterative(trunk *StackNode, allStacks []*StackNode) *StackNode {
	stackMap := make(map[string]*StackNode)
	for _, stack := range allStacks {
		stackMap[stack.Name] = stack
	}

	for _, stack := range allStacks {
		parent := stackMap[stack.ParentBranch]
		if parent != nil {
			stack.Parent = parent
			parent.Children = append(parent.Children, stack)
		}
	}

	return trunk
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

		parentNode := stackMap[branch.ParentBranchName]
		if parentNode != nil {
			node.Parent = parentNode
		}
	}

	trunkName := GetConfig().Trunk
	trunk := stackMap[trunkName]

	return trunk
}

func CacheGraphToDisk(trunk *StackNode) {
	trunk = FindTrunk(trunk)
	branches := bfs(trunk, []Branch{})
	UpdateCache(Cache{Branches: branches})
}

func FindTrunk(node *StackNode) *StackNode {
	if node.Parent == nil {
		return node
	}

	return FindTrunk(node.Parent)
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
