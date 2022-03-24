package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/brandongallagher199/1337x-Bot-Go/config"
	"github.com/brandongallagher199/1337x-Bot-Go/internal/listener"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	discordConfig, err := config.ReadDiscordConf()
	if err != nil {
		log.Fatal(err)
	}
	listener.Create(discordConfig)
}
