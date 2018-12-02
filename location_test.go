package rickandmorty

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLocationsFirstPage(t *testing.T) {
	options := map[string]interface{}{"page": 1}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, locations, opt)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationsFifthtPage(t *testing.T) {
	options := map[string]interface{}{"page": 5}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_fifth-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	comparation := cmp.Equal(pagedResults, locations)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationsWithParamNil(t *testing.T) {
	locations, err := GetLocations(nil)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, locations, opt)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationsWithFilterParams(t *testing.T) {
	options := map[string]interface{}{
		"name": "earth",
	}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_filter_name-earth.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, locations, opt)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationsWithFilterParamsCombined(t *testing.T) {
	options := map[string]interface{}{
		"name": "earth",
		"type": "planet",
	}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_filter_name-earth_type-planet.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, locations, opt)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationsWithRandomFilters(t *testing.T) {
	options := map[string]interface{}{
		"these":   "params",
		"must":    "be",
		"ignored": "by",
		"the":     "function",
		"even":    []string{"with", "this"},
	}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_first-page.json")
	if err != nil {
		t.Error(err)
	}

	pagedResults := new(AllLocations)

	json.Unmarshal(data, &pagedResults)

	opt := sliceEmptyNullReturnTrue()

	comparation := cmp.Equal(pagedResults, locations, opt)

	if !comparation {
		t.Error("The response from GetLocations was:")
		t.Error(locations)
		t.Error("The data against is being run this test is:")
		t.Error(pagedResults)
	}
}

func TestGetLocationOne(t *testing.T) {
	location, err := GetLocation(1)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/location_1.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Location)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, location)

	if !comparation {
		t.Error("The response from GetLocation was:")
		t.Error(location)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetLocationSixtySix(t *testing.T) {
	location, err := GetLocation(66)
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/location_66.json")
	if err != nil {
		t.Error(err)
	}

	result := new(Location)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, location)

	if !comparation {
		t.Error("The response from GetLocation was:")
		t.Error(location)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestGetLocationsArray(t *testing.T) {
	location, err := GetLocationsArray([]int{1, 20})
	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/locations_1-20.json")
	if err != nil {
		t.Error(err)
	}

	result := new(MultipleLocations)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, location)

	if !comparation {
		t.Error("The response from GetLocationsArray was:")
		t.Error(location)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestFromLocationGetResidents(t *testing.T) {
	location, err := GetLocation(21)
	if err != nil {
		t.Error(err)
	}

	residents, err := location.GetResidents()
	if err != nil {
		t.Error(err)
	}

	// Characters from location 21 are:
	ids := []int{7, 436}
	charactersFromSlice, err := GetCharactersArray(ids)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(residents, charactersFromSlice)

	if !comparation {
		t.Error("The response from location.GetResidents was:")
		t.Error(residents)
		t.Error("The data against is being run this test is:")
		t.Error(charactersFromSlice)
	}
}

func TestFromMultipleLocationsGetResidents(t *testing.T) {
	locations, err := GetLocationsArray([]int{1, 21})
	if err != nil {
		t.Error(err)
	}

	residents, err := locations.GetResidents()
	if err != nil {
		t.Error(err)
	}

	// Chracters from episode 1 and 21
	ids1 := []int{38, 45, 71, 82, 83, 92, 112, 114, 116, 117, 120, 127, 155, 169, 175, 179, 186, 201, 216, 239, 271, 302, 303, 338, 343, 356, 394}
	ids21 := []int{7, 436}

	residentsFromLocation1, err := GetCharactersArray(ids1)
	if err != nil {
		t.Error(err)
	}

	residentsFromLocation21, err := GetCharactersArray(ids21)
	if err != nil {
		t.Error(err)
	}

	residentsFromBothLocations := []MultipleCharacters{
		*residentsFromLocation1,
		*residentsFromLocation21,
	}

	comparation := cmp.Equal(residents, residentsFromBothLocations)

	if !comparation {
		t.Error("The response from locations.GetResidents was:")
		t.Error(residents)
		t.Error("The data against is being run this test is:")
		t.Error(residentsFromBothLocations)
	}
}

func TestFromAllLocationsGetNextPage(t *testing.T) {
	options := map[string]interface{}{"page": 1}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	locationsNextPage, err := locations.GetNextPage()
	if err != nil {
		t.Error(err)
	}

	optionsPage2 := map[string]interface{}{"page": 2}

	locationsPage2, err := GetLocations(optionsPage2)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(locationsNextPage, locationsPage2)

	if !comparation {
		t.Error("The response from locations.GetNextPage was:")
		t.Error(locationsNextPage)
		t.Error("The data against is being run this test is:")
		t.Error(locationsPage2)
	}
}

func TestFromAllLocationsGetPreviousPage(t *testing.T) {
	options := map[string]interface{}{"page": 2}

	locations, err := GetLocations(options)
	if err != nil {
		t.Error(err)
	}

	locationsPreviousPage, err := locations.GetPreviousPage()
	if err != nil {
		t.Error(err)
	}

	optionsPage1 := map[string]interface{}{"page": 1}

	locationsPage1, err := GetLocations(optionsPage1)
	if err != nil {
		t.Error(err)
	}

	comparation := cmp.Equal(locationsPreviousPage, locationsPage1)

	if !comparation {
		t.Error("The response from locations.GetPreviousPage was:")
		t.Error(locationsPreviousPage)
		t.Error("The data against is being run this test is:")
		t.Error(locationsPage1)
	}
}
