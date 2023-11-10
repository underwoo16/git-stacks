package commands

import (
	"fmt"
	"log"
	"os/exec"
)

func Stack(args []string) {
	stackName := stackNameFromArgs(args)
	fmt.Println("Creating stack", stackName, "...")

	// ${parent branch ref} - git symbolic-ref HEAD
	cmd := exec.Command("git", "symbolic-ref", "HEAD")
	ref, symRefErr := cmd.Output()

	if symRefErr != nil {
		log.Fatal(symRefErr)
	}
	refName := string(ref)
	fmt.Println("current ref:", refName)

	// ${parent ref sha}
	cmd = exec.Command("git", "rev-parse", "HEAD")
	sha, revParseErr := cmd.Output()

	if revParseErr != nil {
		log.Fatal(revParseErr)
	}
	refSha := string(sha)

	fmt.Println("current ref sha:", refSha)
}

func stackNameFromArgs(args []string) string {
	var nameArg string
	if len(args) > 0 {
		nameArg = args[0]
	}

	return nameArg
}
