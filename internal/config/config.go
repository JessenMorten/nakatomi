package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/JessenMorten/nakatomi/internal/logging"
)

type Config struct {
	LogLevel            logging.LogLevel
	Environment         string
	GetChannelsEndpoint string
	GetListingsEndpoint string
}

type jsonConfig struct {
	LogLevel    string
	Environment string
}

func LoadConfigurationFromFile() (*Config, error) {
	// Create json config
	config := jsonConfig{}

	// Read config file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	// Deserialize json
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// Parse log level
	logLevel, err := parseLogLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	// Create config
	realConfig := Config{}
	realConfig.LogLevel = logLevel
	shh(&realConfig)
	return &realConfig, nil
}

func parseLogLevel(logLevel string) (logging.LogLevel, error) {
	upper := strings.ToUpper(logLevel)

	if upper == "ERROR" {
		return logging.Error, nil
	} else if upper == "WARINNG" {
		return logging.Warning, nil
	} else if upper == "INFORMATION" {
		return logging.Information, nil
	} else if upper == "DEBUG" {
		return logging.Debug, nil
	} else {
		return -1, fmt.Errorf("failed to parse log level: '%v'", logLevel)
	}
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
