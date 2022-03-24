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

	if string(m.Content[0]) == botConfig.Discord.Prefix {
		splitMessage := strings.Fields(m.Content)
		if len(splitMessage) > 1 && splitMessage[0] == fmt.Sprintf("%s%s", botConfig.Discord.Prefix, botConfig.Discord.Command) {
			if splitMessage[1] == "--help" {
				_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Usage: %s%s <query>", botConfig.Discord.Prefix, botConfig.Discord.Command))
				if err != nil {
					fmt.Println(err)
				}
			} else {
				s.ChannelTyping(m.ChannelID)
				queryString := strings.Join(splitMessage[1:], " ")
				torrentLinks := torrentserviceutils.QueryTorrentService(queryString)
				shortened := mgnetmeutils.GetMagnetLinks(torrentLinks[:])
				fieldArray := make([]*discordgo.MessageEmbedField, 0)
				var counter int = 1
				for i := range shortened {
					name := fmt.Sprintf("%d. %s| ", counter, shortened[i].Title)
					value := fmt.Sprintf("%s | Seeds: %d | Size: %s", shortened[i].Magnet, shortened[i].Seeds, shortened[i].Size)
					newField := &discordgo.MessageEmbedField{Name: name, Value: value, Inline: false}
					fieldArray = append(fieldArray, newField)
					counter++
				}
				author := &discordgo.MessageEmbedAuthor{Name: "@" + m.Author.Username}
				embed := &discordgo.MessageEmbed{Type: discordgo.EmbedTypeLink, Author: author, Fields: fieldArray[:]}
				complexMessage := &discordgo.MessageSend{Embed: embed}
				_, err := s.ChannelMessageSendComplex(m.ChannelID, complexMessage)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
