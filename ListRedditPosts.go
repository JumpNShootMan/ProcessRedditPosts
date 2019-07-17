package listnredditposts

import "github.com/turnage/graw/reddit"

//CreateBot creates a bot with credentials
func CreateBot() (reddit.Bot, error) {
	bot, err := reddit.NewBotFromAgentFile("agent", 0)

	if err != nil {
		println("Error en creacion de bot")
		return nil, err
	}

	return bot, err
}

//ListNRedditPosts lists posts from a valid subreddit
func ListNRedditPosts(bot reddit.Bot, n int, subreddit string) (reddit.Bot, error) {
	//pls help

	return nil, nil
}
