package rickandmorty

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func GetEpisodes(options map[string]interface{}) (*AllEpisodes, error) {
	endpoint := endpointEpisode

	hasParams := false
	params := make(map[string]string)

	if options == nil {
		options = map[string]interface{}{
			"endpoint": endpoint,
		}
	}

	for k, v := range options {
		switch v.(type) {
		case int:
			if k == "page" {
				hasParams = true
				params[k] = strconv.FormatInt(int64(v.(int)), 10)
			}
			delete(options, k)
		case string:
			// Skip endpoint in options
			if k == "endpoint" {
				continue
			}
			// Valid parameters to be passed to the parameters map
			validParams := []string{"name", "status", "species", "type", "gender"}
			exists := containsString(validParams, k)
			if exists {
				hasParams = true
				params[k] = v.(string)
			}
			// Cleanup the options map
			delete(options, k)
		default:
			// Cleanup the options map
			delete(options, k)
			// Set the endpoint
			options["endpoint"] = endpoint
		}
	}

	if hasParams {
		options["endpoint"] = endpoint
		options["params"] = params
	}

	data, err := makePetition(options)
	if err != nil {
		return &AllEpisodes{}, err
	}

	episodes := new(AllEpisodes)

	if err := mapstructure.Decode(data, &episodes); err != nil {
		return &AllEpisodes{}, err
	}

	return episodes, nil
}

func GetEpisode(integer int) (*Episode, error) {
	endpoint := endpointEpisode

	options := map[string]interface{}{
		"endpoint": endpoint,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := makePetition(options)
	if err != nil {
		return &Episode{}, err
	}

	episode := new(Episode)

	if err := mapstructure.Decode(data, &episode); err != nil {
		return &Episode{}, err
	}

	return episode, nil
}

func GetEpisodesArray(integers []int) (*MultipleEpisodes, error) {
	endpoint := endpointEpisode

	options := map[string]interface{}{
		"endpoint": endpoint,
		"integers": integers,
	}

	data, err := makePetition(options)
	if err != nil {
		return &MultipleEpisodes{}, err
	}

	episodes := new(MultipleEpisodes)

	if err := mapstructure.Decode(data, &episodes); err != nil {
		return &MultipleEpisodes{}, err
	}

	return episodes, nil
}
