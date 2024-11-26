package main

import "fmt"

type Account struct {
	Username string
	Password string
	Karma    int

	// joined subreddits
	Subreddits []*Subreddit
}

func (a *Account) JoinSubreddit(s *Subreddit) {
	s.Accounts = append(s.Accounts, a)
	a.Subreddits = append(a.Subreddits, s)

	fmt.Printf("User %s joined subreddit %s", a.Username, s.Name)
}

func (a *Account) LeaveSubreddit(s *Subreddit) {
	for i, acc := range s.Accounts {
		if acc.Username == a.Username {
			s.Accounts = append(s.Accounts[:i], s.Accounts[i+1:]...)
		}
	}

	for i, sub := range a.Subreddits {
		if sub.Name == s.Name {
			a.Subreddits = append(a.Subreddits[:i], a.Subreddits[i+1:]...)
		}
	}

	fmt.Printf("User %s left subreddit %s", a.Username, s.Name)
}

func (a *Account) CreatePost(subreddit *Subreddit, text string) {
	newPost := &Post{
		Author:    a,
		Text:      text,
		Upvotes:   0,
		Downvotes: 0,
		Comments:  make([]*Comment, 0),
		Subreddit: subreddit,
	}

	subreddit.Posts = append(subreddit.Posts, newPost)

	fmt.Printf("User %s created a post in subreddit %s", a.Username, subreddit.Name)
}

func (a *Account) UpvotePost(post *Post) {
	post.Upvotes++

	fmt.Printf("User %s upvoted post in subreddit %s", a.Username, post.Subreddit.Name)
}

func (a *Account) DownvotePost(post *Post) {
	post.Downvotes++

	fmt.Printf("User %s downvoted post in subreddit %s", a.Username, post.Subreddit.Name)
}

func (a *Account) CommentOnPost(post *Post, text string) {
	newComment := &Comment{
		Author:   a,
		Text:     text,
		Comments: make([]*Comment, 0),
	}

	post.Comments = append(post.Comments, newComment)

	fmt.Printf("User %s commented on post in subreddit %s", a.Username, post.Subreddit.Name)
}

func (a *Account) CommentOnComment(comment *Comment, text string) {
	newComment := &Comment{
		Author:   a,
		Text:     text,
		Comments: make([]*Comment, 0),
	}

	comment.Comments = append(comment.Comments, newComment)

	fmt.Printf("User %s commented on comment in subreddit %s", a.Username, comment.Author.Subreddits[0].Name)
}
