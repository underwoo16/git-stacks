package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	trunk := stacks.GetGraphFromCache()
	fmt.Printf("%v\n", trunk)
}

// TODO: store columns in a map
func bfs(node *stacks.StackNode, col int, arr []*stacks.StackNode) []*stacks.StackNode {
	arr = append(arr, node)

	childCount := len(node.Children)
	for i := childCount - 1; i >= 0; i-- {
		child := node.Children[i]
		arr = bfs(child, col+i, arr)
	}

	return arr
}

// func dfs(node stacks.StackNode) {
// 	queue := []stacks.StackNode{node}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		fmt.Printf("Visiting: %s\n", current.Name)

// 		for _, child := range current.Children {
// 			queue = append(queue, child)
// 		}
// 	}
// }
