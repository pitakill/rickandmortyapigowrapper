package rickandmorty

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const endpoint = "character/"

func GetCharacters(options map[string]interface{}) (*AllCharacters, error) {
	hasParams := false
	params := make(map[string]string)

	if options == nil {
		options = map[string]interface{}{
			"endpoint": endpoint,
		}
	}

	for k, v := range options {
		switch v.(type) {
		case int:
			if k == "page" {
				hasParams = true
				params[k] = strconv.FormatInt(int64(v.(int)), 10)
			}
			delete(options, k)
		case string:
			// Skip endpoint in options
			if k == "endpoint" {
				continue
			}
			// Valid parameters to be passed to the parameters map
			validParams := []string{"name", "status", "species", "type", "gender"}
			exists := containsString(validParams, k)
			if exists {
				hasParams = true
				params[k] = v.(string)
			}
			// Cleanup the options map
			delete(options, k)
		default:
			// Cleanup the options map
			delete(options, k)
			// Set the endpoint
			options["endpoint"] = endpoint
		}
	}

	if hasParams {
		options["endpoint"] = endpoint
		options["params"] = params
	}

	data, err := makePetition(options)
	if err != nil {
		return &AllCharacters{}, err
	}

	characters := new(AllCharacters)

	if err := mapstructure.Decode(data, &characters); err != nil {
		return &AllCharacters{}, err
	}

	return characters, nil
}

func GetCharacter(integer int) (*Character, error) {
	options := map[string]interface{}{
		"endpoint": endpoint,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := makePetition(options)
	if err != nil {
		return &Character{}, err
	}

	character := new(Character)

	if err := mapstructure.Decode(data, &character); err != nil {
		fmt.Println(err)
		return &Character{}, err
	}

	return character, nil
}
