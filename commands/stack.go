package commands

import (
	"fmt"
	"log"
	"os/exec"
)

func Stack(args []string) {
	stackName := stackNameFromArgs(args)
	fmt.Println("Creating stack", stackName, "...")

	// ${parent branch name} - git symbolic-ref HEAD
	parentBranchRef, err := exec.Command("git", "symbolic-ref", "HEAD").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current ref:", string(parentBranchRef))

	// ${parent branch revision} - git rev-parse ref
	// this isn't working?
	cmd := exec.Command("git", "rev-parse", string(parentBranchRef))
	fmt.Println(cmd.String())

	sha, revParseErr := cmd.Output()
	if revParseErr != nil {
		log.Fatal(revParseErr)
	}

	fmt.Println("current ref sha:", string(sha))
}

func stackNameFromArgs(args []string) string {
	var nameArg string
	if len(args) > 0 {
		nameArg = args[0]
	}

	return nameArg
}
