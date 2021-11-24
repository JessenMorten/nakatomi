package tvguide

import (
	"math/rand"
	"time"

	"github.com/JessenMorten/nakatomi/internal/logging"
)

type dummyTvGuide struct {
	logger logging.Logger
}

var (
	random *rand.Rand
)

func (h dummyTvGuide) Search(title string) ([]TvListing, error) {
	if random == nil {
		h.logger.Debug("Creating new random")
		source := rand.NewSource(time.Now().Unix())
		random = rand.New(source)
	}

	success := random.Int()%2 == 0

	if success {
		return []TvListing{createDummyListing()}, nil
	}

	return []TvListing{}, nil
}

func createDummyListing() TvListing {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+3, 0, 0, 0, time.Local)

	return TvListing{
		Title:        "Die Hard",
		ChannelTitle: "McClane TV",
		StartTime:    startTime,
	}
}
