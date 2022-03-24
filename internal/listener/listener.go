package listener

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/brandongallagher199/1337x-Bot-Go/config"
	"github.com/brandongallagher199/1337x-Bot-Go/internal/mgnetmeutils"

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
				queryString := strings.Join(splitMessage[1:], " ")
				fmt.Println("Query: " + queryString)
				/*here we would place the call for the api, once it returns we would pass the response to a shorener function
				once shortening is done we would return the new message and that would be it.
				*/
				magnetlinks := [2]string{"magnet:?xt=urn:btih:E707E17C8CAF2E4DA0DA99F4E4FC72DA931D42CE&dn=Flaky.2022.1080p.WEB-DL.AAC2.0.H.264-CMRG&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=udp%3A%2F%2F47.ip-51-68-199.eu%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2920%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce", "magnet:?xt=urn:btih:EE02F21D63BF65084A64D712A8B78D1FE7A4F604&dn=Walk.With.Me.2022.1080p.WEB-DL.AAC2.0.H.264-CMRG&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=udp%3A%2F%2F47.ip-51-68-199.eu%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2920%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce"}
				shortened := mgnetmeutils.GetMagnetLinks(magnetlinks[:])
				fieldArray := make([]*discordgo.MessageEmbedField, 0)
				var counter int = 1
				for i := range shortened {
					name := fmt.Sprintf("%d. You need this", counter)
					newField := &discordgo.MessageEmbedField{Name: name, Value: shortened[i], Inline: false}
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
