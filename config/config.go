package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Discord struct {
		Token            string
		Prefix           string
		Command          string
		MaxLinksPerQuery int32 `yaml:"maxLinksPerQuery"`
	}
}

func ReadDiscordConf() (*Conf, error) {
	base, err := os.Getwd()
	base = path.Join(base, "config", "config.yml")
	buf, err := ioutil.ReadFile(base)
	if err != nil {
		log.Fatal(err)
	}

	discordConfig := &Conf{}
	err = yaml.Unmarshal(buf, discordConfig)
	if err != nil {
		return nil, fmt.Errorf("in file %s: %v", "config.yml", err)
	}

	return discordConfig, nil
}
