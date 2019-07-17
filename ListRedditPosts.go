package main

import "github.com/turnage/graw/reddit"

//ListNRedditPosts(n int, subreddit string)
func ListNRedditPosts(n int, subreddit string) (reddit.Bot, error) {

	bot, err := reddit.NewBotFromAgentFile("agent", 0)

	if err != nil {
		return nil, err
		println("error en creacion de bot")
	}
}
