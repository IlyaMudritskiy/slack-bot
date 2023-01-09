package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
	"time"
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

func CountAge(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	var year = request.Param("year")
	yob, err := strconv.Atoi(year)
	var year_now = time.Now().Year()
	if err != nil {
		fmt.Println("Error")
	}
	var age = year_now - yob
	var r = fmt.Sprintf("Age is %d", age)
	response.Reply(r)
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

	var examples = []string{"My year of birth is 1990", "My year of birth is 2000"}

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		Examples: examples,
		Handler: CountAge,
	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var bot_err = bot.Listen(ctx)

	if bot_err != nil {
		log.Fatal(bot_err)
	}
}
