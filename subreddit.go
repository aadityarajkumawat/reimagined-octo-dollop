package main

type Comment struct {
	Author   *Account
	Text     string
	Comments []*Comment
}

type Post struct {
	Author    *Account
	Text      string
	Upvotes   int
	Downvotes int
	Comments  []*Comment
	Subreddit *Subreddit
}

type Subreddit struct {
	// Each subreddit will have a unique name
	Name        string
	Description string

	CreatedBy *Account

	Posts []*Post

	// members of this subreddit
	Accounts []*Account
}
