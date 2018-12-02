package main

import (
	"fmt"
	"log"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

func printResidentsFromLocation(id int) {
	// Get the Location from the Wrapper
	location, err := rnm.GetLocation(id)
	if err != nil {
		log.Printf("Can't get the location %d, %v", id, err)
	}

	// Get the Residents from the episode instance with the Wrapper
	residents, err := location.GetResidents()
	if err != nil {
		log.Printf("Can't get the residents, %v", err)
	}

	print(residents, location.ID, "location")
}

func printResidentsFromMultipleLocations(ids []int) {
	locations, err := rnm.GetLocationsArray(ids)
	if err != nil {
		log.Printf("Can't get the locations %v", err)
	}

	residents, err := locations.GetResidents()
	if err != nil {
		log.Printf("Can't get the residents %v", err)
	}

	for index, residentsFromOneLocation := range residents {
		print(&residentsFromOneLocation, (*locations)[index].ID, "location")
	}
}

func printNextPageLocations() {
	options := map[string]interface{}{"page": 1}

	locations, err := rnm.GetLocations(options)
	if err != nil {
		log.Printf("Can't get the locations %v", err)
	}

	nextLocations, err := locations.GetNextPage()
	if err != nil {
		log.Printf("Can't get the locations %v", err)
	}

	for _, location := range nextLocations.Results {
		fmt.Printf("The location '%s' is the type '%s' and dimension '%s'\n", location.Name, location.Type, location.Dimension)
	}
}
