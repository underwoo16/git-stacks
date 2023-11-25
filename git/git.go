package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/utils"
)

type GitService interface {
	GetCurrentBranch() string
	GetCurrentSha() string
	RevParse(ref string) string
	CreateHashObject(filepath string) string
	UpdateRef(ref string, sha string)
	CreateAndCheckoutBranch(branch string)
	CheckoutBranch(branch string)
	Show(thing string) string
	Rebase(trunk string, branch string) error
	RebaseContinue()
	Commit()
	CommitAmend()
	BranchExists(branch string) bool
	PassThrough(args []string)
	LogBetween(from string, to string) string
	PushBranch(branch string)
	ForcePushBranch(branch string)
	DirectoryPath() string
}

type gitService struct{}

func NewGitService() *gitService {
	return &gitService{}
}

func (g *gitService) GetCurrentBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	utils.CheckError(err)
	branchName := strings.TrimSpace(string(out))
	return branchName
}

func (g *gitService) GetCurrentSha() string {
	return g.RevParse("HEAD")
}

func (g *gitService) RevParse(ref string) string {
	out, err := exec.Command("git", "rev-parse", ref).Output()
	utils.CheckError(err)
	refSha := strings.TrimSpace(string(out))
	return refSha
}

func (g *gitService) CreateHashObject(filepath string) string {
	out, err := exec.Command("git", "hash-object", "-w", filepath).Output()
	utils.CheckError(err)
	objectSha := strings.TrimSpace(string(out))
	return objectSha
}

func (g *gitService) UpdateRef(ref string, sha string) {
	_, err := exec.Command("git", "update-ref", ref, sha).Output()
	utils.CheckError(err)
}

func (g *gitService) CreateAndCheckoutBranch(branch string) {
	_, err := exec.Command("git", "checkout", "-b", branch).Output()
	utils.CheckError(err)
}

func (g *gitService) CheckoutBranch(branch string) {
	_, err := exec.Command("git", "checkout", branch).Output()
	utils.CheckError(err)
}

func (g *gitService) Show(thing string) string {
	out, err := exec.Command("git", "show", thing).Output()
	utils.CheckError(err)
	result := strings.TrimSpace(string(out))
	return result
}

func (g *gitService) Rebase(trunk string, branch string) error {
	_, err := exec.Command("git", "rebase", trunk, branch).Output()
	return err
}

func (g *gitService) RebaseContinue() {
	cmd := exec.Command("git", "rebase", "--continue")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func (g *gitService) Commit() {
	cmd := exec.Command("git", "commit")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func (g *gitService) CommitAmend() {
	cmd := exec.Command("git", "commit", "--amend")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func (g *gitService) BranchExists(branch string) bool {
	out, err := exec.Command("git", "branch").Output()
	utils.CheckError(err)
	branches := string(out)
	return strings.Contains(branches, branch)
}

func (g *gitService) PassThrough(args []string) {
	fmt.Printf(colors.Gray("Running \""))

	cmdStr := fmt.Sprintf("git %s", strings.Join(args, " "))
	fmt.Printf(colors.Yellow(cmdStr))

	fmt.Printf(colors.Gray("\" via git\n\n"))

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func (g *gitService) LogBetween(from string, to string) string {
	out, err := exec.Command("git", "log", "--oneline", "--no-decorate", fmt.Sprintf("%s..%s", from, to)).Output()
	utils.CheckError(err)
	return string(out)
}

func (g *gitService) PushBranch(branch string) {
	_, err := exec.Command("git", "push", "-u", "origin", branch).Output()
	utils.CheckError(err)
}

func (g *gitService) ForcePushBranch(branch string) {
	_, err := exec.Command("git", "push", "-f", "-u", "origin", branch).Output()
	utils.CheckError(err)
}

func (g *gitService) DirectoryPath() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	utils.CheckError(err)

	return fmt.Sprintf("%s/.git", strings.TrimSpace(string(out)))
}
