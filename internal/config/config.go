package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Environment         string
	GetChannelsEndpoint string
	GetListingsEndpoint string
}

func LoadConfigurationFromFile() (*Config, error) {
	config := Config{
		Environment: "Development",
	}
	shh(&config)

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func shh(config *Config) {
	base := []byte{
		104, 116, 116, 112, 115, 58,
		47, 47, 116, 118, 116, 105,
		100, 45, 97, 112, 105, 46,
		97, 112, 105, 46, 116, 118,
		50, 46, 100, 107, 47, 97,
		112, 105, 47, 116, 118, 116,
		105, 100, 47, 118, 49,
	}

	channels := []byte{
		47, 115, 99, 104, 101, 100,
		117, 108, 101, 115, 47, 99,
		104, 97, 110, 110, 101, 108, 115,
	}

	listings := []byte{
		47, 101, 112, 103, 47, 100, 97,
		121, 118, 105, 101, 119, 115, 47,
	}

	config.GetChannelsEndpoint = string(append(base, channels...))
	config.GetListingsEndpoint = string(append(base, listings...))
}
