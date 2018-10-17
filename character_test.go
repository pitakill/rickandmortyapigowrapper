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
