package commands

import (
	"fmt"
	"math"
	"strings"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/stacks"
)

var vertical = "│"
var horizontal = "─"
var spacer = "  "

var circle = "◌"
var dot = "●"

var endBranch = "┘"
var horizBranch = "├"
var vertBranch = "┴"

type ShowCommand struct {
	GitService   git.GitService
	StackService stacks.StackService
}

func (s *ShowCommand) Run() {
	currentBranch := s.GitService.GetCurrentBranch()
	trunk := s.StackService.GetGraph()

	depthStack, colMap := bfs(trunk, 0, []*stacks.StackNode{}, map[string]int{})

	sb := strings.Builder{}
	for depth := len(depthStack) - 1; depth >= 0; depth-- {
		node := depthStack[depth]
		col := colMap[node.Name]
		logBetween := s.GitService.LogBetween(node.ParentBranch, node.Name)

		s.writeRow(&sb, col)
		s.writeStackLabel(&sb, node, currentBranch)

		if depth > 0 {
			s.writeColumns(&sb, col, logBetween)

			if col > 0 && s.isLowestChild(node, depth, depthStack) {
				s.writeConnectingBranches(&sb, col)
			}
		}
	}

	fmt.Print(sb.String())
}

func (s *ShowCommand) writeRow(sb *strings.Builder, col int) {
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

func (s *ShowCommand) writeStackLabel(sb *strings.Builder, node *stacks.StackNode, currentBranch string) {
	nodePrefix := circle
	nodeSuffix := ""
	if node.Name == currentBranch {
		nodePrefix = colors.CurrentStack(dot)
		nodeSuffix = "*"
	}

	if s.StackService.NeedsSync(node) {
		nodeSuffix += " (needs rebase)"
	}

	stackLabel := fmt.Sprintf("%s %s\n", node.Name, nodeSuffix)
	if node.Name == currentBranch {
		stackLabel = colors.CurrentStack(stackLabel)
	} else {
		stackLabel = colors.OtherStack(stackLabel)
	}

	sb.WriteString(nodePrefix + " " + stackLabel)
}

// TODO: clean this up
//
//	maybe move splitting the logs to array in git package
func (s *ShowCommand) writeColumns(sb *strings.Builder, col int, log string) {
	logs := strings.FieldsFunc(log, func(r rune) bool {
		return r == '\n'
	})

	numVerticals := int(math.Max(3, float64(len(logs)+2)))
	for i := 0; i < numVerticals; i++ {
		sb.WriteString(vertical)
		for i := 0; i < col; i++ {
			sb.WriteString(spacer)
			sb.WriteString(vertical)
		}

		logIdx := i - 1
		if logIdx >= 0 && logIdx < len(logs) {
			sb.WriteString(" " + colors.Gray(logs[logIdx]))
		}
		sb.WriteString("\n")
	}
}

// "├ ─ ─ ┴ ─ ─ ┘"
func (s *ShowCommand) writeConnectingBranches(sb *strings.Builder, col int) {
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

func (s *ShowCommand) isLowestChild(child *stacks.StackNode, depth int, arr []*stacks.StackNode) bool {
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
