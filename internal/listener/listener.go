package listener

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/brandongallagher199/1337x-Bot-Go/config"
	"github.com/brandongallagher199/1337x-Bot-Go/internal/mgnetmeutils"
	"github.com/brandongallagher199/1337x-Bot-Go/internal/torrentserviceutils"

	"github.com/bwmarrin/discordgo"
)

var botConfig *config.Conf

func Create(config *config.Conf) {
	botConfig = config

	discord, err := discordgo.New(fmt.Sprintf("Bot %s", botConfig.Discord.Token))

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	discord.AddHandler(messageCreate)
	discord.Identify.Intents = discordgo.IntentGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	discord.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online",
		Activities: []*discordgo.Activity{
			&discordgo.Activity{
				Name: "you use .torrent",
				Type: discordgo.ActivityTypeWatching,
			},
		},
	})
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	if len(m.Content) > 0 {
		if string(m.Content[0]) == botConfig.Discord.Prefix {
			//Trasform potential query into array
			splitMessage := strings.Fields(m.Content)
			command := splitMessage[0]
			switch command {
			case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, botConfig.Discord.Command):
				torrentCmd(s, m, splitMessage)
			case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, "help"):
				helpCmd(s, m)
			case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, "github"):
				githubCmd(s, m)
			}
		}

	}
}

func torrentCmd(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) <= 1 {
		s.ChannelMessageSend(m.ChannelID, "No search query was given. Please run the command again with something to search.")
		return
	}

	s.ChannelTyping(m.ChannelID)
	queryString := strings.Join(args[1:], " ")
	torrentLinks := torrentserviceutils.QueryTorrentService(queryString)

	if len(torrentLinks) == 0 || torrentLinks == nil {
		author := &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")}
		newField := &discordgo.MessageEmbedField{Name: "Not Found", Value: "Torrent not found on 1337x, please refine your search.", Inline: false}

		embed := &discordgo.MessageEmbed{
			Type:   discordgo.EmbedTypeLink,
			Author: author,
			Color:  16711680,
			Fields: []*discordgo.MessageEmbedField{newField},
		}

		complexMessage := &discordgo.MessageSend{Embed: embed}

		_, err := s.ChannelMessageSendComplex(m.ChannelID, complexMessage)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	shortened := mgnetmeutils.GetMagnetLinks(torrentLinks[:])
	fieldArray := make([]*discordgo.MessageEmbedField, 0)

	for idx, torrent := range shortened {
		name := fmt.Sprintf("%d. %s ", idx+1, torrent.Title)
		value := fmt.Sprintf("**[magnet](%s)** | Seeds: %d | Size: %s", torrent.Magnet, torrent.Seeds, torrent.Size)
		newField := &discordgo.MessageEmbedField{Name: name, Value: value, Inline: false}
		fieldArray = append(fieldArray, newField)
	}

	author := &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")}
	footer := &discordgo.MessageEmbedFooter{Text: "For more results, check 1337x."}
	embed := &discordgo.MessageEmbed{Type: discordgo.EmbedTypeLink, Author: author, Footer: footer, Color: 15102219, Fields: fieldArray[:]}
	complexMessage := &discordgo.MessageSend{Embed: embed}

	_, err := s.ChannelMessageSendComplex(m.ChannelID, complexMessage)
	if err != nil {
		fmt.Println(err)
	}
}

func helpCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID,
		"```.torrent <torrent name> // .torrent The Witcher 3 Wild Hunt```")
}

func githubCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "https://github.com/brandongallagher1999/1337x-Bot-Go")
}
