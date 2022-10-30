/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"ops-helper-v2/pkg/funcs"

	"github.com/spf13/cobra"
)

// getAllAccountsCmd represents the getAllAccounts command
var getAllAccountsCmd = &cobra.Command{
	Use:   "get-all-accounts",
	Short: "Get all accounts from HC Organization",
	Run: func(cmd *cobra.Command, args []string) {
		accounts := funcs.GetAccounts()
		for _, account := range accounts {
			fmt.Printf("Success\t%s\t%s\n", *account.Id, *account.Name)
		}
		return
	},
}



func init() {
	rootCmd.AddCommand(getAllAccountsCmd)
}
