package main

import (
	"github.com/JessenMorten/nakatomi/internal/config"
	"github.com/JessenMorten/nakatomi/internal/logging"
	"github.com/JessenMorten/nakatomi/internal/tvguide"
)

func main() {
	// Load configuration
	config, err := config.LoadConfigurationFromFile()
	if err != nil {
		panic(err)
	}

	// Setup dependencies
	logger := logging.NewConsoleLogger(config.LogLevel, "main")
	var g tvguide.TvGuide
	if config.Environment == "Production" {
		g = tvguide.NewHttpTvGuide(
			logging.NewConsoleLogger(config.LogLevel, "http-tvguide"),
			*config,
		)
	} else {
		g = tvguide.NewDummyTvGuide(
			logging.NewConsoleLogger(config.LogLevel, "dummy-tvguide"),
		)
	}

	// Start application
	logger.Information("Nakatomi is running in '%v' environment", config.Environment)

	results, err := g.Search("Die Hard")

	if err != nil {
		logger.Error("Search failed: %v", err)
	}

	logger.Information("Found %v result(s)", len(results))

	for _, result := range results {
		logger.Information("%v at %v on %v", result.Title, result.StartTime, result.ChannelTitle)
	}

	logger.Information("Nakatomi stopped")
}
