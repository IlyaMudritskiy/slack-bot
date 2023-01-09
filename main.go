package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Load .env file
	var err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env variables
	var Slack_Bot_Token = os.Getenv("Slack_Bot_Token")
	var Bot_User_OAuth_Token = os.Getenv("Bot_User_OAuth_Token")

	// Create bot
	var bot = slacker.NewClient(Slack_Bot_Token, Bot_User_OAuth_Token)

	// Print command events in parallel
	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var bot_err = bot.Listen(ctx)

	if bot_err != nil {
		log.Fatal(bot_err)
	}
}
