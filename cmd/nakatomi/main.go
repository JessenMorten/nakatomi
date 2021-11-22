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
	logger := logging.NewConsoleLogger()
	tvguide := tvguide.NewHttpTvGuide(logger, *config)

	// Start application
	logger.Information("Nakatomi is running in '%v' environment", config.Environment)

	results, err := tvguide.Search("Die Hard")

	logger.Information("Found %v result(s)", len(results))

	for _, result := range results {
		logger.Information("%v at %v on %v", result.Title, result.StartTime, result.ChannelTitle)
	}

	logger.Information("Nakatomi stopped")
}
