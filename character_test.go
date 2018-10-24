package rickandmorty

import (
	"encoding/json"
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
