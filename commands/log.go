package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/stacks"
)

var vertical = "│"
var bend = "┘"
var horizontal = "─"
var spacer = "  "
var circle = "◯"
var branch = "├"

func Log() {
	// TODO: use cache if exists
	trunk := stacks.BuildStackGraphFromScratch()

	arr, m := bfs(trunk, 0, []*stacks.StackNode{}, map[string]int{})

	for depth := len(arr) - 1; depth >= 0; depth-- {
		node := arr[depth]
		col := m[node.Name]
		if col > 0 {
			fmt.Printf("%s", vertical)
			for i := 0; i < col; i++ {
				fmt.Printf("%s", spacer)
			}
		}
		fmt.Printf("%s %s\n", circle, node.Name)
		if depth > 0 {
			fmt.Printf("%s", vertical)
			for i := 0; i < col; i++ {
				fmt.Printf("%s", spacer)
				fmt.Printf("%s", vertical)
			}
			fmt.Printf("\n")

			if col > 0 && isLowestChild(node, depth, arr) {
				fmt.Printf("%s", branch)
				for i := 0; i < col; i++ {
					fmt.Printf("%s%s", horizontal, horizontal)
				}
				fmt.Printf("%s\n", bend)
			}
		}
	}
}

func isLowestChild(child *stacks.StackNode, depth int, arr []*stacks.StackNode) bool {
	parent := child.Parent
	if parent == nil {
		return false
	}

	if len(parent.Children) == 1 {
		return false
	}

	for _, sibling := range parent.Children {
		if sibling == child {
			continue
		}

		// find index of sibling in arr
		for i, node := range arr {
			if node == sibling {
				if i < depth {
					return false
				}
			}
		}
	}

	return true

}

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
