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

func TestFromEpisodeGetCharacter(t *testing.T) {
	episode, err := GetEpisode(20)
	if err != nil {
		t.Error(err)
	}

	characters, err := episode.GetCharacters()
	if err != nil {
		t.Error(err)
	}

	// Characters from episode 20 are:
	ids := []int{1, 2, 3, 4, 5, 26, 139, 202, 273, 341}
	charactersFromSlice, err := GetCharactersArray(ids)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(characters, charactersFromSlice)

	if !comparation {
		t.Error("The response from episode.GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(charactersFromSlice)
	}
}

func TestFromMultipleEpisodesGetCharacters(t *testing.T) {
	episodes, err := GetEpisodesArray([]int{24, 26})
	if err != nil {
		t.Error(err)
	}

	characters, err := episodes.GetCharacters()
	if err != nil {
		t.Error(err)
	}

	// Chracters from episode 24 and 36
	ids24 := []int{1, 2, 3, 4, 9, 70, 107, 167, 171, 189, 240, 265, 272, 276, 329}
	ids26 := []int{1, 2, 3, 4, 5, 23, 47, 115, 137, 142, 180, 204, 296, 297, 319, 320, 365, 369, 467, 468, 469}

	charactersFromEpisode24, err := GetCharactersArray(ids24)
	if err != nil {
		t.Error(err)
	}

	charactersFromEpisode26, err := GetCharactersArray(ids26)
	if err != nil {
		t.Error(err)
	}

	charactersFromBothEpisodes := []MultipleCharacters{
		*charactersFromEpisode24,
		*charactersFromEpisode26,
	}

	comparation := cmp.Equal(characters, charactersFromBothEpisodes)

	if !comparation {
		t.Error("The response from episodes.GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(charactersFromBothEpisodes)
	}
}

func TestFromAllEpisodesGetNextPage(t *testing.T) {
	options := map[string]interface{}{"page": 1}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	episodesNextPage, err := episodes.GetNextPage()
	if err != nil {
		t.Error(err)
	}

	optionsPage2 := map[string]interface{}{"page": 2}

	episodesPage2, err := GetEpisodes(optionsPage2)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(episodesNextPage, episodesPage2)

	if !comparation {
		t.Error("The response from episodes.GetNextPage was:")
		t.Error(episodesNextPage)
		t.Error("The data against is being run this test is:")
		t.Error(episodesPage2)
	}
}

func TestFromAllEpisodesGetPreviousPage(t *testing.T) {
	options := map[string]interface{}{"page": 2}

	episodes, err := GetEpisodes(options)
	if err != nil {
		t.Error(err)
	}

	episodesPreviousPage, err := episodes.GetPreviousPage()
	if err != nil {
		t.Error(err)
	}

	optionsPage1 := map[string]interface{}{"page": 1}

	episodesPage1, err := GetEpisodes(optionsPage1)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(episodesPreviousPage, episodesPage1)

	if !comparation {
		t.Error("The response from episodes.GetPreviousPage was:")
		t.Error(episodesPreviousPage)
		t.Error("The data against is being run this test is:")
		t.Error(episodesPage1)
	}
}
