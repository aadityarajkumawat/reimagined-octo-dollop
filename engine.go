package main

import "fmt"

type Engine struct {
	Accounts   []*Account
	Subreddits []*Subreddit
}

func (e *Engine) CreateNewAccount(username, password string) {
	newAcc := &Account{
		Username:   username,
		Password:   password,
		Karma:      0,
		Subreddits: make([]*Subreddit, 0),
	}

	e.Accounts = append(e.Accounts, newAcc)

	fmt.Println("Account created: ", username)
}

func (e *Engine) CreateNewSubreddit(by *Account, name, description string) {
	newSub := &Subreddit{
		Name:        name,
		Description: description,
		CreatedBy:   by,
	}

	e.Subreddits = append(e.Subreddits, newSub)

	fmt.Println("Subreddit created: ", name)
}
