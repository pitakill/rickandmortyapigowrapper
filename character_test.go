package rickandmorty

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCharactersFirstPage(t *testing.T) {
	options := map[string]interface{}{
		"page": 1,
	}

	characters, err := GetCharacters(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharactersFifthPage(t *testing.T) {
	options := map[string]interface{}{
		"page": 5,
	}

	characters, err := GetCharacters(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_fifth-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharactersWithParamNil(t *testing.T) {
	characters, err := GetCharacters(nil)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharactersWithFilterParams(t *testing.T) {
	options := map[string]interface{}{
		"name":   "rick",
		"status": "alive",
	}

	characters, err := GetCharacters(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_filter_name-rick_status-alive.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharactersWithFilterParamsAndRandomFilters(t *testing.T) {
	options := map[string]interface{}{
		"name":    "rick",
		"status":  "alive",
		"these":   "params",
		"must":    "be",
		"ignored": "by the function",
	}

	characters, err := GetCharacters(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_filter_name-rick_status-alive.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharactersWithdRandomFilters(t *testing.T) {
	options := map[string]interface{}{
		"these":   "params",
		"must":    "be",
		"ignored": "by",
		"the":     "function",
		"even":    []string{"with", "this"},
	}

	characters, err := GetCharacters(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllCharacters)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, characters)

	if !comparation {
		t.Error("The response from GetCharacters was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetCharacterOne(t *testing.T) {
	character, err := GetCharacter(1)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/character_1.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Character)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, character)

	if !comparation {
		t.Error("The response from GetCharacter was:")
		t.Error(character)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetCharacterSixtySix(t *testing.T) {
	character, err := GetCharacter(66)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/character_66.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Character)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, character)

	if !comparation {
		t.Error("The response from GetCharacter was:")
		t.Error(character)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetCharactersArray(t *testing.T) {
	characters, err := GetCharactersArray([]int{1, 183})
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/characters_1-183.json")
	if err != nil {
		t.Error(err)
	}

	result := new(MultipleCharacters)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, characters)
	if !comparation {
		t.Error("The response from GetCharactersArray was:")
		t.Error(characters)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestFromCharacterGetLocation(t *testing.T) {
	character, err := GetCharacter(61)
	if err != nil {
		t.Error(err)
	}

	location, err := character.GetLocation()
	if err != nil {
		t.Error(err)
	}

	// The location from the character 61 is 3
	location3, err := GetLocation(3)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(location, location3)

	if !comparation {
		t.Error("The response from character.GetLocation was:")
		t.Error(location)
		t.Error("The data against is being run this test is:")
		t.Error(location3)
	}
}

func TestFromCharacterGetLocationUnknown(t *testing.T) {
	character, err := GetCharacter(60)
	if err != nil {
		t.Error(err)
	}

	location, err := character.GetLocation()
	if err != nil {
		t.Error(err)
	}

	// The location from the character 60 is unknown
	// So the Method returns and Empty location but without error
	emptyLocation := &Location{}

	comparation := cmp.Equal(location, emptyLocation)

	if !comparation {
		t.Error("The response from character.GetLocation was:")
		t.Error(location)
		t.Error("The data against is being run this test is:")
		t.Error(emptyLocation)
	}
}

func TestFromCharacterGetEpisodes(t *testing.T) {
	character, err := GetCharacter(1)
	if err != nil {
		t.Error(err)
	}

	episodes, err := character.GetEpisodes()
	if err != nil {
		t.Error(err)
	}

	// The episodes of the Character 60 are
	idEpisodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	episodesFromCharacter1, err := GetEpisodesArray(idEpisodes)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(episodes, episodesFromCharacter1)

	if !comparation {
		t.Error("The response from character.GetEpisodes was:")
		t.Error(episodes)
		t.Error("The data against is being run this test is:")
		t.Error(episodesFromCharacter1)
	}
}

func TestFromCharacterGetOrigin(t *testing.T) {
	character, err := GetCharacter(1)
	if err != nil {
		t.Error(err)
	}

	origin, err := character.GetOrigin()
	if err != nil {
		t.Error(err)
	}

	// The origin of the character 1 is
	originFrom1, err := GetLocation(1)
	if err != nil {
		fmt.Println(err)
	}

	comparation := cmp.Equal(origin, originFrom1)

	if !comparation {
		t.Error("The response from character.GetOrigin was:")
		t.Error(origin)
		t.Error("The data against is being run this test is:")
		t.Error(originFrom1)
	}
}
