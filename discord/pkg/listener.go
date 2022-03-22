package listener

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func create() {
	discord, err := discordgo.New("Bot " + "authentication token")

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(messageCreate)

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	var prefix string = "."

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
