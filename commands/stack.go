package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/underwoo16/git-stacks/utils"
)

func Stack(args []string) {
	stackName := stackNameFromArgs(args)
	fmt.Printf("Creating stack %s...", stackName)

	// ${parent branch ref}
	out, err := exec.Command("git", "symbolic-ref", "HEAD").Output()
	utils.CheckError(err)
	refName := string(out)

	// ${parent ref sha}
	out, err = exec.Command("git", "rev-parse", "HEAD").Output()
	utils.CheckError(err)
	refSha := string(out)

	tempFilePath := fmt.Sprintf(".git/temp-%s", stackName)
	hashObject := fmt.Sprintf("%s\n%s", strings.TrimSpace(refName), strings.TrimSpace(refSha))
	utils.WriteToFile(tempFilePath, hashObject)

	out, err = exec.Command("git", "hash-object", "-w", tempFilePath).Output()
	utils.CheckError(err)
	objSha := strings.TrimSpace(string(out))

	newRef := fmt.Sprintf("refs/stacks/%s", stackName)
	_, err = exec.Command("git", "update-ref", newRef, objSha).Output()
	utils.CheckError(err)

	_, err = exec.Command("git", "checkout", "-b", stackName).Output()
	utils.CheckError(err)

	err = os.Remove(tempFilePath)
	utils.CheckError(err)
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
