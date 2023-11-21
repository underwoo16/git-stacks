package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/underwoo16/git-stacks/colors"
	"github.com/underwoo16/git-stacks/utils"
)

func GetCurrentBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	utils.CheckError(err)
	branchName := strings.TrimSpace(string(out))
	return branchName
}

func GetCurrentRef() string {
	out, err := exec.Command("git", "symbolic-ref", "HEAD").Output()
	utils.CheckError(err)
	refName := strings.TrimSpace(string(out))
	return refName
}

func GetCurrentSha() string {
	return RevParse("HEAD")
}

func RevParse(ref string) string {
	out, err := exec.Command("git", "rev-parse", ref).Output()
	utils.CheckError(err)
	refSha := strings.TrimSpace(string(out))
	return refSha
}

func CreateHashObject(filepath string) string {
	out, err := exec.Command("git", "hash-object", "-w", filepath).Output()
	utils.CheckError(err)
	objectSha := strings.TrimSpace(string(out))
	return objectSha
}

func UpdateRef(ref string, sha string) {
	_, err := exec.Command("git", "update-ref", ref, sha).Output()
	utils.CheckError(err)
}

func CreateAndCheckoutBranch(branch string) {
	_, err := exec.Command("git", "checkout", "-b", branch).Output()
	utils.CheckError(err)
}

func CheckoutBranch(branch string) {
	_, err := exec.Command("git", "checkout", branch).Output()
	utils.CheckError(err)
}

func Show(thing string) string {
	out, err := exec.Command("git", "show", thing).Output()
	utils.CheckError(err)
	result := strings.TrimSpace(string(out))
	return result
}

func Rebase(trunk string, branch string) {
	_, err := exec.Command("git", "rebase", trunk, branch).Output()
	utils.CheckError(err)
}

func Commit() {
	cmd := exec.Command("git", "commit")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func CommitAmend() {
	cmd := exec.Command("git", "commit", "--amend")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func BranchExists(branch string) bool {
	out, err := exec.Command("git", "branch").Output()
	utils.CheckError(err)
	branches := string(out)
	return strings.Contains(branches, branch)
}

func PassThrough(args []string) {
	fmt.Printf(colors.Gray("Running: \""))

	cmdStr := fmt.Sprintf("git %s", strings.Join(args, " "))
	fmt.Printf(colors.Yellow(cmdStr))

	fmt.Printf(colors.Gray("\"\n"))

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils.CheckError(err)
}

func LogBetween(from string, to string) string {
	out, err := exec.Command("git", "log", "--oneline", "--no-decorate", fmt.Sprintf("%s..%s", from, to)).Output()
	utils.CheckError(err)
	return string(out)
}

func PushBranch(branch string) {
	_, err := exec.Command("git", "push", "-u", "origin", branch).Output()
	utils.CheckError(err)
}

func ForcePushBranch(branch string) {
	_, err := exec.Command("git", "push", "-f", "-u", "origin", branch).Output()
	utils.CheckError(err)
}
