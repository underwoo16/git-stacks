package main

import (
	"testing"

	"github.com/underwoo16/git-stacks/stacks"
)

func BenchmarkBuildGraphRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trunk := stacks.StackNode{
			Name:     "master",
			RefSha:   "1234567890",
			Children: []*stacks.StackNode{},
		}

		allStacks := []*stacks.StackNode{
			{
				Name:         "feature/1",
				RefSha:       "1234567890",
				ParentBranch: "master",
				Children:     []*stacks.StackNode{},
			},
			{
				Name:         "feature/2",
				RefSha:       "1234567890",
				ParentBranch: "master",
				Children:     []*stacks.StackNode{},
			},
			{
				Name:         "feature/3",
				RefSha:       "1234567890",
				ParentBranch: "master",
				Children:     []*stacks.StackNode{},
			},
			{
				Name:         "feature/4",
				RefSha:       "1234567890",
				ParentBranch: "feature/3",
				Children:     []*stacks.StackNode{},
			},
			{
				Name:         "feature/5",
				RefSha:       "1234567890",
				ParentBranch: "feature/3",
				Children:     []*stacks.StackNode{},
			},
			{
				Name:         "feature/6",
				RefSha:       "1234567890",
				ParentBranch: "feature/3",
				Children:     []*stacks.StackNode{},
			},
		}

		stacks.BuildGraphIterative(&trunk, allStacks)
	}
}
