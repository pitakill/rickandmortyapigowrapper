package main

import (
	"fmt"
	"log"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

func printLocationFromCharacter(id int) {
	character, err := rnm.GetCharacter(id)
	if err != nil {
		log.Printf("Can't get the character %d", id)
	}

	location, err := character.GetLocation()
	if err != nil {
		log.Printf("Can't get the location %v", err)
	}

	fmt.Printf("The location of %s is %s\n", character.Name, location.Name)
}

func printEpisodesFromCharacter(id int) {
	character, err := rnm.GetCharacter(id)
	if err != nil {
		log.Printf("Can't get the character %d", id)
	}

	episodes, err := character.GetEpisodes()
	if err != nil {
		log.Printf("Can't get the episodes %v", err)
	}

	for i, episode := range *episodes {
		if i == 0 {
			fmt.Printf("The episodes of %s are %s", character.Name, episode.Name)
			continue
		}

		if i == len(*episodes)-1 {
			fmt.Printf("%s.\n")
			break
		}

		fmt.Printf(", %s", episode.Name)
	}
}

func printOriginFromCharacter(id int) {
	character, err := rnm.GetCharacter(id)
	if err != nil {
		log.Printf("Can't get the character %d", id)
	}

	origin, err := character.GetOrigin()
	if err != nil {
		log.Printf("Can't get the origin %v", err)
	}

	fmt.Printf("The origin of %s is %s\n", character.Name, origin.Name)
}
