package main

import (
	"fmt"

	"github.com/turnage/graw/reddit"
)

//CreateBot creates a bot with credentials
func CreateBot() (reddit.Bot, error) {
	bot, err := reddit.NewBotFromAgentFile("agent", 0)

	if err != nil {
		println("Error en creacion de bot")
		return nil, err
	}

	return bot, err
}

//FetchNRedditPosts lists posts from a valid subreddit
func FetchNRedditPosts(bot reddit.Bot, n int, subreddit string) ([]*reddit.Post, error) {
	//pls help
	harvest, err := bot.Listing("/r/"+subreddit, "")

	if err != nil {
		fmt.Println("Error fetching from subreddit")
		return nil, err
	}

	PostsArray := harvest.Posts[:n]
	return PostsArray, nil
}

//func List

func main() {

	redditbot, err := CreateBot()
	if err != nil {
		panic(err)
	}
	PostsArray, err := FetchNRedditPosts(redditbot, 50, "programming")

	for _, post := range PostsArray {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}
