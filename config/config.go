package config

import (
	"fmt"
	"io/ioutil"

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

func ReadDiscordConf(filename string) (*Conf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	discordConfig := &Conf{}
	err = yaml.Unmarshal(buf, discordConfig)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return discordConfig, nil
}
