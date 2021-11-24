package tvguide

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/JessenMorten/nakatomi/internal/config"
	"github.com/JessenMorten/nakatomi/internal/logging"
)

type httpTvGuide struct {
	config config.Config
	logger logging.Logger
}

type getChannelsResponse struct {
	Channels []channel `json:"channels"`
}

type getListingsResponse struct {
	Id       string    `json:"id"`
	Listings []listing `json:"programs"`
}

type listing struct {
	Id            string `json:"id"`
	StartTimeUnix int64  `json:"start"`
	StopTimeUnix  int64  `json:"stop"`
	Title         string `json:"title"`
}

type channel struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	IconUrl    string `json:"icon"`
	LogoUrl    string `json:"logo"`
	SvgLogoUrl string `json:"svgLogo"`
	Sort       int    `json:"sort"`
	Language   string `json:"language"`
}

func (h httpTvGuide) Search(title string) ([]TvListing, error) {
	// Get all channels
	h.logger.Information("Searching for listings containing '%v'", title)
	url := h.config.GetChannelsEndpoint
	channelsResponse := getChannelsResponse{}
	err := h.getFromJson(url, &channelsResponse)

	if err != nil {
		return nil, err
	}

	h.logger.Debug("Found %v channel(s)", len(channelsResponse.Channels))

	// Find listings for today
	date := time.Now()
	results := []TvListing{}
	for _, channel := range channelsResponse.Channels {
		url = h.config.GetListingsEndpoint + date.Format("2006-01-02") + "?ch=" + channel.Id
		listingsResponse := []getListingsResponse{}
		err := h.getFromJson(url, &listingsResponse)
		if err != nil {
			return nil, err
		}

		// Check if title exists
		for _, p := range listingsResponse {
			for _, listing := range p.Listings {
				upperTitle := strings.ToUpper(listing.Title)
				isMatch := strings.Contains(upperTitle, strings.ToUpper(title))

				if isMatch {
					results = append(results, TvListing{
						Title:        listing.Title,
						ChannelTitle: channel.Title,
						StartTime:    time.Unix(listing.StartTimeUnix, 0),
					})
				}
			}
		}
	}

	return results, nil
}

func (h httpTvGuide) getFromJson(url string, v interface{}) error {
	// Create request
	h.logger.Debug("GET %v", url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/json")

	// Send request
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	// Return serialized result
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}
