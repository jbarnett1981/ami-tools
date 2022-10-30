/*
Copyright © 2022 Julian Barnett jbarnett@housecanary.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ops-helper",
	Short: "ops-helper script to query all accounts in HC and tag AMIs",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("region", "--region", "AWS Region")
}


