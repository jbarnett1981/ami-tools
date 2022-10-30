/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tagAllAccountsCmd represents the tagAllAccounts command
var tagAllAccountsCmd = &cobra.Command{
	Use:   "tag-all-accounts",
	Short: "Tags AMI to all accounts in HC Organization",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tag-all-accounts called")
	},
}

func init() {
	rootCmd.AddCommand(tagAllAccountsCmd)
	tagAllAccountsCmd.PersistentFlags().String("ami-id", "--ami-id", "The ami id to tag")
	tagAllAccountsCmd.PersistentFlags().String("tag-key", "--key", "The account name (Not ARN)")
	tagAllAccountsCmd.PersistentFlags().String("tag-value", "--value", "The account name (Not ARN)")
	tagAllAccountsCmd.PersistentFlags().String("role-name", "--role", "The account name (Not ARN)")
}
