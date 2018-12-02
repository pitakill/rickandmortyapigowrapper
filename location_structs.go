package rickandmorty

import (
	"errors"
	"strconv"
)

type Location struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	Residents []string `json:"residents"`
	URL       string   `json:"url"`
	Created   string   `json:"created"`
}

type AllLocations struct {
	Info    Info       `json:"info"`
	Results []Location `json:"results"`
}

type MultipleLocations []Location

func (l *Location) GetResidents() (*MultipleCharacters, error) {
	var idsCharacters []int

	for _, characterURL := range l.Residents {
		characterIDString := getLastElementSplitedBy(characterURL, "/")
		characterID, err := strconv.Atoi(characterIDString)
		if err != nil {
			return &MultipleCharacters{}, err
		}

		idsCharacters = append(idsCharacters, characterID)
	}

	characters, err := GetCharactersArray(idsCharacters)
	if err != nil {
		return &MultipleCharacters{}, err
	}

	return characters, nil
}

func (l *MultipleLocations) GetResidents() ([]MultipleCharacters, error) {
	charactersFromAllLocations := []MultipleCharacters{}

	for _, location := range *l {
		residents, err := location.GetResidents()
		if err != nil {
			return []MultipleCharacters{}, err
		}

		charactersFromAllLocations = append(charactersFromAllLocations, *residents)
	}

	return charactersFromAllLocations, nil
}

func (al *AllLocations) GetNextPage() (*AllLocations, error) {
	if al.Info.Next == "" {
		return &AllLocations{}, errors.New("Nothing here")
	}

	pageString := getLastElementSplitedBy(al.Info.Next, "=")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return &AllLocations{}, err
	}

	options := map[string]interface{}{"page": page}

	return GetLocations(options)
}

func (al *AllLocations) GetPreviousPage() (*AllLocations, error) {
	if al.Info.Prev == "" {
		return &AllLocations{}, errors.New("Nothing here")
	}

	pageString := getLastElementSplitedBy(al.Info.Prev, "=")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return &AllLocations{}, err
	}

	options := map[string]interface{}{"page": page}

	return GetLocations(options)
}
