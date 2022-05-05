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
	if len(m.Content) > 0 {
		if string(m.Content[0]) == botConfig.Discord.Prefix {
			//Trasform potential query into array
			splitMessage := strings.Fields(m.Content)
			if len(splitMessage) > 1 { //If .torrent <torrent> is called
				switch splitMessage[0] {
					case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, botConfig.Discord.Command):
						s.ChannelTyping(m.ChannelID)
						queryString := strings.Join(splitMessage[1:], " ")
						torrentLinks := torrentserviceutils.QueryTorrentService(queryString)
						if len(torrentLinks) == 0 || torrentLinks == nil {

							author := &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL(""),
							newField := &discordgo.MessageEmbedField{Name: "Not Found", Value: "Torrent not found on 1337x, please refine your search.", Inline: false}
							fieldArray := make([]*discordgo.MessageEmbedField, 0)
							fieldArray = append(fieldArray, newField)
							embed := &discordgo.MessageEmbed{Type: discordgo.EmbedTypeLink, Author: author, Fields: fieldArray[:]}
							complexMessage := &discordgo.MessageSend{Embed: embed}
							_, err := s.ChannelMessageSendComplex(m.ChannelID, complexMessage)
							if err != nil {
								fmt.Println(err)
							}
							return
						}
						shortened := mgnetmeutils.GetMagnetLinks(torrentLinks[:])
						fieldArray := make([]*discordgo.MessageEmbedField, 0)
						var counter int = 1
						for i := range shortened {
							name := fmt.Sprintf("%d. %s ", counter, shortened[i].Title)
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
					case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, "help"):
						s.ChannelMessageSend(m.ChannelID, 
							"``` .torrent <torrent name> // .torrent The Witcher 3 Wild Hunt \n .invite // Invite link to the Discord Bot```")
					case fmt.Sprintf("%s%s", botConfig.Discord.Prefix, "github"):
						s.ChannelMessageSend(m.ChannelID, "https://github.com/brandongallagher1999/1337x-Bot-Go")
				}
			}
		}
	}

}
