package main

import (
	"bot/discord"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var counterUrl string
var discordBotToken string
var unused string

func init() {
	// Get the URL from the environment variable
	counterUrl = os.Getenv("COUNTER_API_URL")
	if counterUrl == "" {
		fmt.Println("Error: COUNTER_API_URL environment variable not set")
		os.Exit(1)
	}
	discordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
	if discordBotToken == "" {
		fmt.Println("Error: DISCORD_BOT_TOKEN environment variable not set")
		os.Exit(1)
	}
}

func main() {
	discord.Start(discordBotToken, counterUrl)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
