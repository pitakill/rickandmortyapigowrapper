package rickandmorty

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func makePetition(options map[string]interface{}) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	endpoint := ""
	hasParams := false

	for k, v := range options {
		switch v.(type) {
		case string:
			if k == "endpoint" {
				endpoint = v.(string)
			}
		case map[string]string:
			if k == "params" {
				hasParams = true
			}
		default:
			err := fmt.Sprintf("The option %s is type: %T (value: %v), must be a string.", k, v, v)
			return nil, errors.New(err)
		}
	}

	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	if hasParams {
		for key, value := range options["params"].(map[string]string) {
			q.Add(key, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := make(map[string]interface{})

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
