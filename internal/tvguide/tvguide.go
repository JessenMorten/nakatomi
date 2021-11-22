package tvguide

import (
	"time"

	"github.com/JessenMorten/nakatomi/internal/config"
	"github.com/JessenMorten/nakatomi/internal/logging"
)

type TvListing struct {
	Title        string
	ChannelTitle string
	StartTime    time.Time
}

type TvGuide interface {
	Search(title string) ([]TvListing, error)
}

func NewHttpTvGuide(logger logging.Logger, config config.Config) TvGuide {
	return httpTvGuide{
		logger: logger,
		config: config,
	}
}
