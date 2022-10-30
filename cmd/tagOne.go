/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tagOneCmd represents the tagOne command
var tagOneCmd = &cobra.Command{
	Use:   "tag-one",
	Short: "Tag AMI to a single account",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test")

		// accountName, err := cmd.Flags().GetString("account-name")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// roleName, err := cmd.Flags().GetString("role-name")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// amiId, err := cmd.Flags().GetString("ami-id")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// tagKey, err := cmd.Flags().GetString("tag-key")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// tagValue, err := cmd.Flags().GetString("tag-value")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// region, err := cmd.Flags().GetString("region")
		// if err != nil {
		// 	region := "us-west-2"
		// }

		// if !args[].IsSet("account-name") {
		// 	fmt.Println("Missing --account-name -n")
		// 	return nil
		// }

		// if !c.IsSet("ami-id") {
		// 	fmt.Println("Missing --ami-id -i")
		// 	return nil
		// }

		// if !c.IsSet("tag-key") {
		// 	fmt.Println("Missing --tag-key -k")
		// 	return nil
		// }

		///
	// 	accounts := GetAccounts()
	// 	for _, account := range accounts {

	// 		if *account.Name != accountName {
	// 			continue
	// 		}
	// 		creds, sess := GetCredentials(*account.Id, roleName, region)
	// 		svc := ec2.New(sess, &aws.Config{Credentials: creds})
	// 		resp := DescribeAMI(amiId, svc)

	// 		if resp.Images == nil {
	// 			fmt.Printf("Failed\t%s\t%s\n", *account.Id, *account.Name)
	// 			continue
	// 		}

	// 		TagAMI(amiId, tagKey, tagValue, svc)

	// 		fmt.Printf("Success\t%s\t%s\n", *account.Id, *account.Name)

	// 	}
	},
}

func init() {
	rootCmd.AddCommand(tagOneCmd)
	tagOneCmd.PersistentFlags().String("account-name", "--account-name", "The account name (Not ARN)")
	tagOneCmd.PersistentFlags().String("ami-id", "--ami-id", "The ami id to tag")
	tagOneCmd.PersistentFlags().String("tag-key", "--key", "The account name (Not ARN)")
	tagOneCmd.PersistentFlags().String("tag-value", "--value", "The account name (Not ARN)")
	tagOneCmd.PersistentFlags().String("role-name", "--role", "The account name (Not ARN)")
}
