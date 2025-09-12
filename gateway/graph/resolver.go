package graph

import (
	account "aurora-graph/account/client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AccountClient *account.Client
}

func NewGraphQLServer(accountUrl string) (*Resolver, error) {
	accClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		AccountClient: accClient,
	}, nil
}
