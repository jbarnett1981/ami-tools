/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentVersion = "UNKNOWN"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
		Short: "Prints the current version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}