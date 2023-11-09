package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/underwoo16/flapjacks/utils"
)

func Stack(args []string) {
	fmt.Println(args)
	stackExists := utils.FileExists(".git/refs/stack")

	if !stackExists {
		utils.CreateFile(".git/refs/stack")

		stackName := stackNameFromArgs(args)
		fmt.Println("Creating new stack: ", stackName)

		utils.WriteToFile(".git/refs/stack", stackName)
	}

	// get existing stack name - first line in .git/refs/stack

	// get existing branch name - line in .git/refs/stack with * prefix

	// 3. Create new branch in stack - incrementing number

	// 4. add reference to branch in stack file

	// 5. checkout new branch

}

func stackNameFromArgs(args []string) string {
	var nameArg string
	if len(args) > 0 {
		nameArg = args[0]
		fmt.Println("Stack name arg: ", nameArg)
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter stack name: ")
		input, _ := reader.ReadString('\n')
		fmt.Println("You entered:", input)
		nameArg = input
	}

	return nameArg
}
