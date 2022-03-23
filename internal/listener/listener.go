package listener

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/brandongallagher199/1337x-Bot-Go/config"

	"github.com/bwmarrin/discordgo"
)

var prefix string = "."
var command string = "torrent"

func Create() {

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory.", err)
		return
	}

	configFilePath := filepath.Join(curDir, "config", "config.yml")

	config, err := config.ReadDiscordConf(configFilePath)
	if err != nil {
		fmt.Println("error loading discord config,", err)
		return
	}

	discord, err := discordgo.New("Bot " + config.Discord.Token)

	prefix = config.Discord.Prefix
	command = config.Discord.Command

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

	if string(m.Content[0]) == prefix {
		splitMessage := strings.Fields(m.Content)
		if len(splitMessage) > 1 && splitMessage[0] == prefix+command {
			if splitMessage[1] == "--help" {
				_, err := s.ChannelMessageSend(m.ChannelID, "Usage: "+prefix+command+" <query>")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				queryString := strings.Join(splitMessage[1:], " ")

				/*here we would place the call for the api, once it returns we would pass the response to a shorener function
				once shortening is done we would return the new message and that would be it.
				*/
				_, err := s.ChannelMessageSend(m.ChannelID, "Hello "+m.Author.Username+", your query was: "+queryString)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
