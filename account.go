package main

import (
	"fmt"
	"math/rand"
)

type DirectMessage struct {
	From *Account
	To   *Account
	Text string
}

type Account struct {
	Username string
	Password string
	Karma    int

	// direct messages
	DirectMessages []*DirectMessage

	// joined subreddits
	Subreddits []*Subreddit
}

func (a *Account) JoinSubreddit(s *Subreddit) {
	s.Accounts = append(s.Accounts, a)
	a.Subreddits = append(a.Subreddits, s)

	fmt.Printf("User %s joined subreddit %s\n", a.Username, s.Name)
}

func (a *Account) LeaveSubreddit(s *Subreddit) {
	// owner cannot leave
	if s.CreatedBy.Username == a.Username {
		fmt.Printf("User %s is the owner of subreddit %s so they cannot leave\n", a.Username, s.Name)
		return
	}

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

	fmt.Printf("User %s left subreddit %s\n", a.Username, s.Name)
}

func (a *Account) CreatePost(subreddit *Subreddit, text string) {
	// check if user is in subreddit
	for _, sub := range a.Subreddits {
		if sub.Name == subreddit.Name {
			fmt.Printf("User cannot create post in subreddit they are not in\n")
			break
		}
	}

	newPost := &Post{
		Author:    a,
		Text:      text,
		Upvotes:   0,
		Downvotes: 0,
		Comments:  make([]*Comment, 0),
		Subreddit: subreddit,
	}

	subreddit.Posts = append(subreddit.Posts, newPost)

	fmt.Printf("User %s created a post in subreddit %s\n", a.Username, subreddit.Name)
}

func (a *Account) UpvotePost(post *Post) {
	post.Upvotes++
	post.Author.Karma += 5
	fmt.Printf("User %s upvoted post in subreddit %s\n", a.Username, post.Subreddit.Name)
}

func (a *Account) DownvotePost(post *Post) {
	post.Downvotes++
	post.Author.Karma--
	fmt.Printf("User %s downvoted post in subreddit %s\n", a.Username, post.Subreddit.Name)
}

func (a *Account) CommentOnPost(post *Post, text string) {
	newComment := &Comment{
		Author:   a,
		Text:     text,
		Comments: make([]*Comment, 0),
	}

	post.Comments = append(post.Comments, newComment)

	fmt.Printf("User %s commented on post in subreddit %s\n", a.Username, post.Subreddit.Name)
}

func (a *Account) CommentOnComment(comment *Comment, text string) {
	newComment := &Comment{
		Author:   a,
		Text:     text,
		Comments: make([]*Comment, 0),
	}

	comment.Comments = append(comment.Comments, newComment)

	fmt.Printf("User %s commented on comment in subreddit\n", a.Username)
}

func (a *Account) CreateNewSubreddit(name, description string) *Subreddit {
	newSub := &Subreddit{
		Name:        name,
		Description: description,
		CreatedBy:   a,
	}

	a.Subreddits = append(a.Subreddits, newSub)

	fmt.Printf("User %s created subreddit %s\n", a.Username, name)

	return newSub
}

func (a *Account) GetFeed() []*Post {
	// choose random number of posts between 1 and 30
	// from the subreddits the user is in

	subreddits := a.Subreddits

	if len(subreddits) == 0 {
		fmt.Printf("User %s is not in any subreddits\n", a.Username)
		return []*Post{}
	}

	// get random number of subreddits
	numSubs := rand.Intn(len(subreddits)) + 1

	// get random subreddits
	subs := make([]*Subreddit, 0)
	for i := 0; i < numSubs; i++ {
		index := rand.Intn(len(subreddits))
		subs = append(subs, subreddits[index])
	}

	posts := make([]*Post, 0)
	for _, sub := range subs {
		posts = append(posts, sub.Posts...)
	}

	fmt.Printf("User %s got feed\n", a.Username)
	return posts
}

func (a *Account) GetUserKarma() int {
	fmt.Printf("User %s got karma %d\n", a.Username, a.Karma)
	return a.Karma
}

func (a *Account) SendDirectMessage(to *Account, text string) {
	newMessage := &DirectMessage{
		From: a,
		To:   to,
		Text: text,
	}
	a.DirectMessages = append(a.DirectMessages, newMessage)

	fmt.Printf("User %s sent direct message to user %s\n", a.Username, to.Username)
}
