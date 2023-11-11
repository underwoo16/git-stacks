package commands

import (
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/utils"
)

func Stack(args []string) {
	stackName := stackNameFromArgs(args)
	msg := fmt.Sprintf("Creating stack '%s'...", stackName)
	fmt.Println(msg)

	parentBranchRef := git.GetCurrentRef()

	parentRefSha := git.GetCurrentSha()

	tempFilePath := fmt.Sprintf(".git/temp-%s", stackName)
	hashObject := fmt.Sprintf("%s\n%s", parentBranchRef, parentRefSha)
	utils.WriteToFile(tempFilePath, hashObject)

	objectSha := git.CreateHashObject(tempFilePath)
	utils.RemoveFile(tempFilePath)

	newRef := fmt.Sprintf("refs/stacks/%s", stackName)
	git.UpdateRef(newRef, objectSha)

	git.CreateAndCheckoutBranch(stackName)

	msg = fmt.Sprintf("Done! Switched to new stack '%s'", stackName)
	fmt.Println(msg)
}

func stackNameFromArgs(args []string) string {
	if len(args) == 0 {
		fmt.Println("No stack name provided")
		os.Exit(1)
	}

	return args[0]
}
