/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// stackCmd represents the stack command
var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Create a new stack, or switch to an existing stack.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stack called")

		stackIdentifier := args[0]
		if stackIdentifier != "" {
			// TODO: check if stack exists
			// 		if it does, switch to it
			fmt.Println("stack name: ", stackIdentifier)
			return
		}

		// TODO: generate sequential stack name
		//       ie "stack-1", "stack-2", etc incrementing from current stack name
		gitcmd := exec.Command("git", "checkout", "-b", "test/new-branch")

		var out strings.Builder
		gitcmd.Stdout = &out

		err := gitcmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(out.String())
	},
}

func init() {
	rootCmd.AddCommand(stackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
