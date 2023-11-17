package commands

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type branch struct {
	Name          string
	CommitMessage string
	Sha           string
}

func Test() {
	branches := []branch{
		{Name: "first", Sha: "abcde", CommitMessage: "start work"},
		{Name: "develop", Sha: "x8734", CommitMessage: "Added new interface"},
		{Name: "feature", Sha: "g02g02", CommitMessage: "Created UI"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Name | green }}",
		Inactive: "  {{ .Name | yellow }}",
		Selected: "* {{ .Name | green | green }}",
		Details: `
--------- {{ .Name }} ----------
{{ .CommitMessage | faint }}
{{ .Sha | faint }}`,
	}

	searcher := func(input string, index int) bool {
		branch := branches[index]
		name := strings.ToLower(branch.Name)
		input = strings.ToLower(input)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select child branch",
		Items:     branches,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Switched to %s\n", branches[i].Name)
}
