package git

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type GitHubService interface {
	GetPullRequests() []PullRequest
	CreatePullRequest(baseBranch string, headBranch string)
}

type PullRequest struct {
	BaseRefName string `json:"baseRefName"`
	HeadRefName string `json:"headRefName"`
	Id          string `json:"id"`
	Number      int    `json:"number"`
	Url         string `json:"url"`
}

type gitHubService struct{}

// TODO: Add func to test if gh is installed

func NewGitHubService() *gitHubService {
	return &gitHubService{}
}

func (gh *gitHubService) GetPullRequests() []PullRequest {
	out, err := exec.Command("gh", "pr", "list", "--author", "@me", "--json", "number,baseRefName,headRefName,url").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pullRequests := []PullRequest{}
	err = json.Unmarshal(out, &pullRequests)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pullRequests
}

// TODO: Default title, body, and submit
func (gh *gitHubService) CreatePullRequest(baseBranch string, headBranch string) {
	cmd := exec.Command("gh", "pr", "create", "-B", baseBranch, "-H", headBranch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
