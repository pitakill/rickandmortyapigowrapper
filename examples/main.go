package main

import (
	"fmt"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

func main() {
	// Episode
	//printCharactersFromEpisode(30)
	//printCharactersFromMultipleEpisodes([]int{30, 24, 1, 2})
	//printNextPageEpisodes()

	// Location
	//printResidentsFromLocation(1)
	//printResidentsFromMultipleLocations([]int{3, 21})
	printNextPageLocations()
}

func print(characters *rnm.MultipleCharacters, episode int, typeSection string) {
	for i, character := range *characters {
		name := character.Name

		if i == 0 {
			fmt.Printf("The characters from the %s %d are: %s", typeSection, episode, name)
			continue
		}

		if i == len(*characters)-1 {
			fmt.Printf(" and %s.\n", name)
			break
		}

		fmt.Printf(", %s", name)
	}
}
