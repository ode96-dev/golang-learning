/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "generates base64 string",
	Long: `provided length, it generates base64 string
	For example:
	securerandom base64 -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		str, err := generateBase64String(length)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}

	},
}

func generateBase64String(byteLength int) (string, error) {
	//generate random bytes
	bytes := make([]byte, byteLength)

	//generate random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// convert to hex and truncate to the desired length
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String, nil

}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.Flags().IntP("length", "l", 4, "Length of base 64")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
