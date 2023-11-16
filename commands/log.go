package commands

import (
	"fmt"
	"strings"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

var vertical = "│"
var horizontal = "─"
var spacer = "  "

// TODO: use different symbols for current stack and other stacks
var circle = "◯"
var dot = "◉"

var endBranch = "┘"
var horizBranch = "├"
var vertBranch = "┴"

func Log() {
	currentBranch := git.GetCurrentBranch()
	trunk := stacks.GetGraph()

	depthStack, colMap := bfs(trunk, 0, []*stacks.StackNode{}, map[string]int{})

	sb := strings.Builder{}
	for depth := len(depthStack) - 1; depth >= 0; depth-- {
		node := depthStack[depth]
		col := colMap[node.Name]

		writeRow(&sb, col)
		writeStackLabel(&sb, node, currentBranch)

		if depth > 0 {
			writeColumns(&sb, col)

			if col > 0 && isLowestChild(node, depth, depthStack) {
				writeConnectingBranches(&sb, col)
			}
		}
	}

	fmt.Print(sb.String())
}

func writeRow(sb *strings.Builder, col int) {
	if col > 0 {
		sb.WriteString(vertical)
		for i := 0; i < col; i++ {
			if i > 0 {
				sb.WriteString(vertical)
			}
			sb.WriteString(spacer)
		}
	}
}

func writeStackLabel(sb *strings.Builder, node *stacks.StackNode, currentBranch string) {
	nodePrefix := circle
	nodeSuffix := ""
	if node.Name == currentBranch {
		nodePrefix = colors.CurrentStack(dot)
		nodeSuffix = "*"
	}

	if stacks.NeedsSync(node) {
		nodeSuffix += " (needs sync)"
	}

	stackLabel := fmt.Sprintf("%s %s\n", node.Name, nodeSuffix)
	if node.Name == currentBranch {
		stackLabel = colors.CurrentStack(stackLabel)
	} else {
		stackLabel = colors.OtherStack(stackLabel)
	}

	sb.WriteString(nodePrefix + " " + stackLabel)
}

func writeColumns(sb *strings.Builder, col int) {
	for i := 0; i < 3; i++ {
		sb.WriteString(vertical)
		for i := 0; i < col; i++ {
			sb.WriteString(spacer)
			sb.WriteString(vertical)
		}
		sb.WriteString("\n")
	}
}

// "├ ─ ─ ┴ ─ ─ ┘"
func writeConnectingBranches(sb *strings.Builder, col int) {
	sb.WriteString(horizBranch)
	for i := 0; i < col; i++ {
		sb.WriteString(horizontal + horizontal)
		if i < col-1 {
			sb.WriteString(vertBranch)
		} else {
			sb.WriteString(endBranch + "\n")
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
