package main

import (
	"log"

	"github.com/brandongallagher199/1337x-Bot-Go/config"
	"github.com/brandongallagher199/1337x-Bot-Go/internal/listener"
)

func main() {
	discordConfig, err := config.ReadDiscordConf()
	if err != nil {
		log.Fatal(err)
	}
	listener.Create(discordConfig)
}
