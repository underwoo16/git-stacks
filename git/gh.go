package git

import (
	"os"
	"os/exec"

	"github.com/underwoo16/git-stacks/utils"
)

// TODO: Default title, body, and submit
func CreatePullRequest(baseBranch string, headBranch string) {
	cmd := exec.Command("gh", "pr", "create", "-B", baseBranch, "-H", headBranch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	utils.CheckError(err)
}
