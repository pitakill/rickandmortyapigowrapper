package rickandmorty

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func GetCharacters(options map[string]interface{}) (*AllCharacters, error) {
	endpoint := "character/"

	if options == nil {
		options = map[string]interface{}{
			"endpoint": endpoint,
		}
	}

	for k, v := range options {
		switch v.(type) {
		case int:
			if k == "page" {
				options = map[string]interface{}{
					"endpoint": endpoint,
					"params": map[string]string{
						k: strconv.FormatInt(int64(v.(int)), 10),
					},
				}
			}
		}
	}

	data, err := makePetition(options)
	if err != nil {
		return &AllCharacters{}, err
	}

	characters := new(AllCharacters)

	err = mapstructure.Decode(data, &characters)
	if err != nil {
		return &AllCharacters{}, err
	}

	return characters, nil
}
