package rickandmorty

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetEpisodesFirstPage(t *testing.T) {
	options := map[string]interface{}{"page": 1}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, episodes)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodesSecondPage(t *testing.T) {
	options := map[string]interface{}{"page": 2}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_second-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, episodes)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodesWithParamNil(t *testing.T) {
	episodes, err := GetEpisodes(nil)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, episodes, opt)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodesWithFilterParams(t *testing.T) {
	options := map[string]interface{}{
		"name": "earth",
	}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_filter_name-earth.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, episodes, opt)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodesWithFilterParamsCombined(t *testing.T) {
	options := map[string]interface{}{
		"name": "earth",
		"type": "planet",
	}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_filter_name-earth_type-planet.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, episodes, opt)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodesWithRandomFilters(t *testing.T) {
	options := map[string]interface{}{
		"these":   "params",
		"must":    "be",
		"ignored": "by",
		"the":     "function",
		"even":    []string{"with", "this"},
	}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllEpisodes)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, episodes, opt)

	if !comparation {
		t.Error("The response from GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetEpisodeOne(t *testing.T) {
	episode, err := GetEpisode(1)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episode_1.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Episode)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, episode)

	if !comparation {
		t.Error("The response from GetEpisode was:")
		t.Error(episode)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetEpisodeSixtySix(t *testing.T) {
	episode, err := GetEpisode(31)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episode_31.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Episode)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, episode)

	if !comparation {
		t.Error("The response from GetEpisode was:")
		t.Error(episode)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetEpisodesArray(t *testing.T) {
	episode, err := GetEpisodesArray([]int{1, 20})
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/episodes_1-20.json")
	if err != nil {
		t.Error(err)
	}

	result := new(MultipleEpisodes)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, episode)

	if !comparation {
		t.Error("The response from GetEpisodesArray was:")
		t.Error(episode)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}
