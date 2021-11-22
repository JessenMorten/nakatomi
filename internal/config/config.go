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
