package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "cus",
	Short: "Command line utility",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cus version 0.0.1")
	},
}

var hiCMD = &cobra.Command{
	Use:   "hi [name]",
	Short: "hi command to say Hi",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hi %s!\n", args[0])
	},
}

func main() {
	rootCMD.AddCommand(hiCMD)
	err := rootCMD.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
