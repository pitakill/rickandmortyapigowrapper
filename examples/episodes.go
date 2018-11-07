package main

import (
	"fmt"
	"log"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

func main() {
	//printCharactersFromEpisode(30)
	//printCharactersFromMultipleEpisodes([]int{30, 24, 1, 2})
	printNextPageEpidodes()
}

func printCharactersFromEpisode(id int) {
	// Get the Episode from the Wrapper
	episode, err := rnm.GetEpisode(id)
	if err != nil {
		log.Printf("Can't get the episode %d, %v", id, err)
	}

	// Get the Characters from the episode instance with the Wrapper
	characters, err := episode.GetCharacters()
	if err != nil {
		log.Printf("Can't get the characters, %v", err)
	}

	print(characters, episode.ID)
}

func printCharactersFromMultipleEpisodes(ids []int) {
	// The naive approach is use the printCharactersFromEpisode
	// but that functions make 2 petitions per episode, we can do better with

	// Get all episodes in one petition with the Wrapper
	episodes, err := rnm.GetEpisodesArray(ids)
	if err != nil {
		log.Printf("Can't get the episodes %v, %v", ids, err)
	}

	// Get all the characters from the episodes in as many as len(episodes)
	// petitions
	characters, err := episodes.GetCharacters()
	if err != nil {
		log.Printf("Can't get the characters, %v", err)
	}

	for index, charactersFromOneEpisode := range characters {
		// We have to make this because the rickandmorty api returns an ordered
		// array by ID ascending order ie.
		// make a petition with the episodes ids: 30, 5, 3, 8
		// the api returns an array with the info of the episodes [3, 5, 8, 30]
		print(&charactersFromOneEpisode, (*episodes)[index].ID)
	}
}

func printNextPageEpidodes() {
	options := map[string]interface{}{"page": 1}

	episodes, err := rnm.GetEpisodes(options)
	if err != nil {
		log.Printf("Can't get the episodes %v", err)
	}

	nextEpisodes, err := episodes.GetNextPage()
	if err != nil {
		log.Printf("Can't get the episodes %v", err)
	}

	for _, episode := range nextEpisodes.Results {
		fmt.Printf("The episode %s was first broadcasted in %s\n", episode.Name, episode.Air_Date)
	}
}

func print(characters *rnm.MultipleCharacters, episode int) {
	for i, character := range *characters {
		name := character.Name

		if i == 0 {
			fmt.Printf("The characters from the episode %d are: %s", episode, name)
			continue
		}

		if i == len(*characters)-1 {
			fmt.Printf(" and %s.\n", name)
			break
		}

		fmt.Printf(", %s", name)
	}
}
