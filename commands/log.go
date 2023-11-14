package commands

import (
	"fmt"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

var vertical = "│"
var horizontal = "─"
var spacer = "  "
var circle = "◯"
var dot = "◉"
var bend = "┘"
var horizBranch = "├"
var vertBranch = "┴"

func Log() {
	// TODO: use cache if exists
	currentBranch := git.GetCurrentBranch()
	trunk := stacks.BuildStackGraphFromScratch()

	depthStack, colMap := bfs(trunk, 0, []*stacks.StackNode{}, map[string]int{})

	for depth := len(depthStack) - 1; depth >= 0; depth-- {
		node := depthStack[depth]
		col := colMap[node.Name]
		if col > 0 {
			fmt.Printf("%s", vertical)
			for i := 0; i < col; i++ {
				if i > 0 {
					fmt.Printf("%s", vertical)
				}
				fmt.Printf("%s", spacer)
			}
		}

		nodePrefix := circle
		nodeSuffix := ""
		if node.Name == currentBranch {
			nodePrefix = colors.CurrentStack(dot)
			nodeSuffix = "*"
		}

		stackLabel := fmt.Sprintf("%s %s\n", node.Name, nodeSuffix)
		if node.Name == currentBranch {
			stackLabel = colors.CurrentStack(stackLabel)
		} else {
			stackLabel = colors.OtherStack(stackLabel)
		}

		fmt.Printf("%s ", nodePrefix)
		fmt.Printf(stackLabel)

		if depth > 0 {
			for i := 0; i < 3; i++ {
				fmt.Printf("%s", vertical)
				for i := 0; i < col; i++ {
					fmt.Printf("%s", spacer)
					fmt.Printf("%s", vertical)
				}
				fmt.Printf("\n")
			}

			// "├ ─ ─ ┴ ─ ─ ┘"
			if col > 0 && isLowestChild(node, depth, depthStack) {
				fmt.Printf("%s", horizBranch)
				for i := 0; i < col; i++ {
					fmt.Printf("%s%s", horizontal, horizontal)
					if i < col-1 {
						fmt.Printf("%s", vertBranch)
					} else {
						fmt.Printf("%s\n", bend)
					}

				}
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

func bfs(node *stacks.StackNode, col int, depthStack []*stacks.StackNode, colMap map[string]int) ([]*stacks.StackNode, map[string]int) {
	depthStack = append(depthStack, node)
	colMap[node.Name] = col

	childCount := len(node.Children)
	for i := childCount - 1; i >= 0; i-- {
		child := node.Children[i]
		depthStack, colMap = bfs(child, col+i, depthStack, colMap)
	}

	return depthStack, colMap
}
