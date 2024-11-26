package main

import "fmt"

type Engine struct {
	Accounts []*Account
}

func (e *Engine) CreateNewAccount(username, password string) *Account {
	newAcc := &Account{
		Username:   username,
		Password:   password,
		Karma:      0,
		Subreddits: make([]*Subreddit, 0),
	}

	e.Accounts = append(e.Accounts, newAcc)

	fmt.Println("Account created: ", username+"\n")

	return newAcc
}

func (e *Engine) GetSubreddits() []*Subreddit {
	subs := make([]*Subreddit, 0)
	for _, acc := range e.Accounts {
		subs = append(subs, acc.Subreddits...)
	}

	return subs
}

func (e *Engine) ListAccounts() {
	for _, acc := range e.Accounts {
		// format nicely
		fmt.Printf("%s\n", acc.Username)
	}
}
