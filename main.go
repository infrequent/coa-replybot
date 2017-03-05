package main

import (
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	BotID string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// If the message is "Hello" reply with ":wave: Welcome to the Government Service! How can I help you today?"
	if m.Content == "Hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, ":wave: Welcome to the Government Service! How can I help you today?")
	}

	// If the message is "I'd like to find information about paying rates, please!" reply with "Sure, here is the link you're after: http://govt.org.au/rates"
	if m.Content == "I'd like to find information about paying rates, please!" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Sure, here is the link you're after: http://govt.org.au/rates")
	}
}
