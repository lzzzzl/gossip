package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Neo4jConfig struct {
	URL      string `toml:"url"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Config struct {
	Neo4j Neo4jConfig `toml:"neo4j"`
}

func (config *Config) Load() {
	var configData []byte
	cfg := os.Getenv("GOSSIP_CONFIG")
	if cfg != "" {
		configData = config.GetConfigDataFromFile(cfg)
	} else {
		panic("config file not found")
	}
	config.LoadFromToml(configData)
}

func (config *Config) GetConfigDataFromFile(cfg string) []byte {
	_, err := os.Stat(cfg)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(cfg)
	if err != nil {
		panic(err)
	}
	return data
}

func (config *Config) LoadFromToml(b []byte) {
	_, err := toml.Decode(string(b), config)
	if err != nil {
		panic(err)
	}
}
