package git

import (
	"os/exec"
	"strings"

	"github.com/underwoo16/git-stacks/utils"
)

func GetCurrentRef() string {
	out, err := exec.Command("git", "symbolic-ref", "HEAD").Output()
	utils.CheckError(err)
	refName := strings.TrimSpace(string(out))
	return refName
}

func GetCurrentSha() string {
	return RevParse("HEAD")
}

func RevParse(rev string) string {
	out, err := exec.Command("git", "rev-parse", rev).Output()
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
