package git

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/underwoo16/git-stacks/utils"
)

type PullRequest struct {
	BaseRefName string `json:"baseRefName"`
	HeadRefName string `json:"headRefName"`
	Id          string `json:"id"`
	Number      int    `json:"number"`
	Url         string `json:"url"`
}

// TODO: Add func to test if gh is installed

func GetPullRequests() []PullRequest {
	out, err := exec.Command("gh", "pr", "list", "--author", "@me", "--json", "number,baseRefName,headRefName,url").Output()
	utils.CheckError(err)

	pullRequests := []PullRequest{}
	err = json.Unmarshal(out, &pullRequests)
	utils.CheckError(err)
	return pullRequests
}

// TODO: Default title, body, and submit
func CreatePullRequest(baseBranch string, headBranch string) {
	cmd := exec.Command("gh", "pr", "create", "-B", baseBranch, "-H", headBranch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	utils.CheckError(err)
}
