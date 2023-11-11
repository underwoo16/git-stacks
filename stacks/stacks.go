package stacks

import (
	"strings"

	"github.com/underwoo16/git-stacks/git"
)

type Stack struct {
	Name            string
	ParentBranchRef string
	ParentRefSha    string
}

func ReadStack(ref string) Stack {
	out := git.Show(ref)
	items := strings.Fields(out)

	return Stack{Name: ref, ParentBranchRef: items[0], ParentRefSha: items[1]}
}
