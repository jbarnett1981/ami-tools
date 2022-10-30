package funcs

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
)

//
func loadInput(input *organizations.ListAccountsInput, svc organizations.Organizations, l *log.Logger, accounts []*organizations.Account) []*organizations.Account {

	result, err := svc.ListAccounts(input)
	if err != nil {
	}

	accts := []*organizations.Account{}
	for _, v := range result.Accounts {
		if *v.Status == "ACTIVE" {
			accts = append(accts, v)
		}
	}

	// if there is a NextToken
	if result.NextToken != nil {
		input.SetNextToken(*result.NextToken)
		return loadInput(input, svc, l, accts)
	}
	return append(accounts, accts...)

}

// GetAccounts todo
func GetAccounts() []*organizations.Account {
	

	svc := organizations.New(session.New())
	input := &organizations.ListAccountsInput{}
	l := log.New(os.Stderr, "", 0)
	accounts := make([]*organizations.Account, 20)

	return loadInput(input, *svc, l, accounts)

}