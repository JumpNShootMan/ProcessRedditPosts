package main

import "github.com/turnage/graw/reddit"

//ListNRedditPosts(n int, subreddit string)

func ListNRedditPosts(n int, subreddit string) (reddit.Bot, error) {

	bot, err := reddit.NewBotFromAgentFile("agent", 0)

	if err != nil {
		println("Error en creacion de bot")
		return nil, err
	}

	return bot, err
}
