/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "greets a person",
	Long:  `this command greets the person whose name you provide`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello, %s\n", name)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name of the person to greet")
}
