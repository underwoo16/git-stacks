package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	trunk := stacks.BuildStackGraphFromScratch()
	fmt.Printf("%v\n", trunk)
	arr := bfs(trunk, 0, []*stacks.StackNode{})
	for depth := len(arr) - 1; depth >= 0; depth-- {
		fmt.Printf("%s\n", arr[depth].Name)
	}
}

// TODO: store columns in a map
func bfs(node *stacks.StackNode, col int, arr []*stacks.StackNode) []*stacks.StackNode {
	arr = append(arr, node)
	fmt.Printf("%s col %d depth %d\n", node.Name, col, len(arr))

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
