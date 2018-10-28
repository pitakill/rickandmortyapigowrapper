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
