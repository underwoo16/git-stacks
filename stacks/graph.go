package stacks

func (s *StackService) GetGraph() *StackNode {
	if s.metadataService.CacheExists() {
		return s.GetGraphFromCache()
	}

	graph := s.GetGraphFromRefs()
	s.CacheGraphToDisk(graph)
	return graph
}

func (s *StackService) GetGraphFromRefs() *StackNode {
	config := s.metadataService.GetConfig()

	trunk := StackNode{
		Name:     config.Trunk,
		RefSha:   s.gitService.RevParse(config.Trunk),
		Children: []*StackNode{},
	}

	allStacks := s.getStacks()
	allStacks = append(allStacks, &trunk)
	return s.BuildGraphIterative(&trunk, allStacks)
}

func (s *StackService) BuildGraphIterative(trunk *StackNode, allStacks []*StackNode) *StackNode {
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

func (s *StackService) GetGraphFromCache() *StackNode {
	branches := s.metadataService.GetCache().Branches
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

	trunkName := s.metadataService.GetConfig().Trunk
	trunk := stackMap[trunkName]

	return trunk
}

func (s *StackService) CacheGraphToDisk(trunk *StackNode) {
	trunk = s.FindTrunk(trunk)
	branches := bfs(trunk, []Branch{})

	s.metadataService.UpdateCache(Cache{Branches: branches})
}

func (s *StackService) FindTrunk(node *StackNode) *StackNode {
	if node.Parent == nil {
		return node
	}

	return s.FindTrunk(node.Parent)
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
