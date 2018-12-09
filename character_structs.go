package rickandmorty

import (
	"strconv"
)

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	URL     string   `json:"url"`
	Created string   `json:"created"`
}

type AllCharacters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

type MultipleCharacters []Character

func (c *Character) GetLocation() (*Location, error) {
	if c.Location.URL == "" {
		return &Location{}, nil
	}

	locationIDString := getLastElementSplitedBy(c.Location.URL, "/")
	locationID, err := strconv.Atoi(locationIDString)
	if err != nil {
		return &Location{}, err
	}

	return GetLocation(locationID)
}

func (c *Character) GetEpisodes() (*MultipleEpisodes, error) {
	var idsEpisodes []int

	for _, episodeURL := range c.Episode {
		episodeIDString := getLastElementSplitedBy(episodeURL, "/")
		episodeID, err := strconv.Atoi(episodeIDString)
		if err != nil {
			return &MultipleEpisodes{}, err
		}
		idsEpisodes = append(idsEpisodes, episodeID)
	}

	return GetEpisodesArray(idsEpisodes)
}

func (c *Character) GetOrigin() (*Location, error) {
	if c.Origin.URL == "" {
		return &Location{}, nil
	}

	locationIDString := getLastElementSplitedBy(c.Origin.URL, "/")
	locationID, err := strconv.Atoi(locationIDString)
	if err != nil {
		return &Location{}, err
	}

	return GetLocation(locationID)
}
