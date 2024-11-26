package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/google/uuid"
)

func getRandomUser(engine *Engine) *Account {
	// get random index
	index := rand.Intn(len(engine.Accounts))
	return engine.Accounts[index]
}

type AccountActor struct {
	Account *Account
}

type JoinSubredditMsg struct {
	Subreddit *Subreddit
}

type LeaveSubredditMsg struct {
	Subreddit *Subreddit
}

type CreateSubredditMsg struct {
	Name        string
	Description string
}

type CreatePostMsg struct {
	Subreddit *Subreddit
	Text      string
}

type UpvotePostMsg struct {
	Post *Post
}

type DownvotePostMsg struct {
	Post *Post
}

type CommentOnPostMsg struct {
	Post *Post
	Text string
}

type CommentOnCommentMsg struct {
	Comment *Comment
	Text    string
}

type GetFeedMsg struct{}

type GetKarmaMsg struct{}

type SendDirectMessageMsg struct {
	To   *Account
	Text string
}

func (state *AccountActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *JoinSubredditMsg:
		state.Account.JoinSubreddit(msg.Subreddit)
	case *LeaveSubredditMsg:
		state.Account.LeaveSubreddit(msg.Subreddit)
	case *CreateSubredditMsg:
		state.Account.CreateNewSubreddit(msg.Name, msg.Description)
	case *CreatePostMsg:
		state.Account.CreatePost(msg.Subreddit, msg.Text)
	case *UpvotePostMsg:
		state.Account.UpvotePost(msg.Post)
	case *DownvotePostMsg:
		state.Account.DownvotePost(msg.Post)
	case *CommentOnPostMsg:
		state.Account.CommentOnPost(msg.Post, msg.Text)
	case *CommentOnCommentMsg:
		state.Account.CommentOnComment(msg.Comment, msg.Text)
	case *GetFeedMsg:
		state.Account.GetFeed()
	case *GetKarmaMsg:
		state.Account.GetUserKarma()
	case *SendDirectMessageMsg:
		state.Account.SendDirectMessage(msg.To, msg.Text)
	}
}

const (
	MAX_USERS = 300
)

func main() {
	system := actor.NewActorSystem()
	engine := &Engine{
		Accounts: make([]*Account, 0),
	}

	users := 0

	// Create 100 user actors
	for users < MAX_USERS {
		actionType := rand.Intn(2) + 1

		account := &Account{}

		if len(engine.Accounts) == 0 {
			actionType = 1
		} else {
			account = getRandomUser(engine)
		}

		if actionType == 1 {
			totalUsers := len(engine.Accounts)
			username := fmt.Sprintf("User%d", totalUsers+1)
			account = engine.CreateNewAccount(username, "pass")
			users++
		}

		props := actor.PropsFromProducer(func() actor.Actor {
			return &AccountActor{Account: account}
		})

		pid := system.Root.Spawn(props)

		// Create a random number between 1 and 8
		actionIndex := rand.Intn(10) + 1

		if actionIndex == 1 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				system.Root.Send(pid, &JoinSubredditMsg{Subreddit: randomSubreddit})
			} else {
				// fmt.Println("No subreddits to join")
			}
		} else if actionIndex == 2 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				system.Root.Send(pid, &LeaveSubredditMsg{Subreddit: randomSubreddit})
			} else {
				// fmt.Println("No subreddits to leave")
			}
		} else if actionIndex == 3 {
			system.Root.Send(pid, &CreateSubredditMsg{Name: "Subreddit | " + uuid.New().String(), Description: "test"})
		} else if actionIndex == 4 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				system.Root.Send(pid, &CreatePostMsg{Subreddit: randomSubreddit, Text: "test"})
			} else {
				// fmt.Println("No subreddits to leave")
			}
		} else if actionIndex == 5 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				posts := randomSubreddit.Posts
				if len(posts) != 0 {
					randomPost := posts[rand.Intn(len(posts))]
					system.Root.Send(pid, &UpvotePostMsg{Post: randomPost})
				} else {
					// fmt.Println("No posts to upvote")
				}
			} else {
				// fmt.Println("No subreddits to upvote")
			}
		} else if actionIndex == 6 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				posts := randomSubreddit.Posts
				if len(posts) != 0 {
					randomPost := posts[rand.Intn(len(posts))]
					system.Root.Send(pid, &DownvotePostMsg{Post: randomPost})
				} else {
					// fmt.Println("No posts to downvote")
				}
			} else {
				// fmt.Println("No subreddits to downvote")
			}
		} else if actionIndex == 7 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				posts := randomSubreddit.Posts
				if len(posts) != 0 {
					randomPost := posts[rand.Intn(len(posts))]
					system.Root.Send(pid, &CommentOnPostMsg{Post: randomPost, Text: "test"})
				} else {
					// fmt.Println("No posts to comment on")
				}
			} else {
				// fmt.Println("No subreddits to comment on")
			}
		} else if actionIndex == 8 {
			subs := engine.GetSubreddits()
			if len(subs) != 0 {
				randomSubreddit := subs[rand.Intn(len(subs))]
				posts := randomSubreddit.Posts
				if len(posts) != 0 {
					randomPost := posts[rand.Intn(len(posts))]
					comments := randomPost.Comments
					if len(comments) != 0 {
						randomComment := comments[rand.Intn(len(comments))]
						system.Root.Send(pid, &CommentOnCommentMsg{Comment: randomComment, Text: "test"})
					} else {
						// fmt.Println("No comments to comment on")
					}
				} else {
					// fmt.Println("No posts to comment on")
				}
			} else {
				// fmt.Println("No subreddits to comment on")
			}
		} else if actionIndex == 9 {
			system.Root.Send(pid, &GetFeedMsg{})
		} else if actionIndex == 10 {
			to := getRandomUser(engine)
			system.Root.Send(pid, &SendDirectMessageMsg{To: to, Text: "test"})
		}

		system.Root.Send(pid, &GetKarmaMsg{})
		time.Sleep(50 * time.Millisecond) // Simulate staggered interactions

	}
}
