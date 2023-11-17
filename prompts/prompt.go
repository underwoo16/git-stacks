package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/underwoo16/git-stacks/utils"
)

func PromptUser(items []string, label string, templates *promptui.SelectTemplates) string {
	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.ToLower(item)
		input = strings.ToLower(input)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	utils.CheckError(err)

	return items[i]
}
