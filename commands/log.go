package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

func Log() {
	// TODO: use cache if exists
	trunk := stacks.BuildStackGraphFromScratch()

	arr, m := bfs(trunk, 0, []*stacks.StackNode{}, map[string]int{})

	// TODO: print out in reverse order
	// 		taking into account column and parent column
	for depth := len(arr) - 1; depth >= 0; depth-- {
		node := arr[depth]
		fmt.Printf("%s %d\n", node.Name, m[node.Name])
	}
}

// TODO: store columns in a map
func bfs(node *stacks.StackNode, col int, arr []*stacks.StackNode, m map[string]int) ([]*stacks.StackNode, map[string]int) {
	arr = append(arr, node)
	m[node.Name] = col

	childCount := len(node.Children)
	for i := childCount - 1; i >= 0; i-- {
		child := node.Children[i]
		arr, m = bfs(child, col+i, arr, m)
	}

	return arr, m
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
