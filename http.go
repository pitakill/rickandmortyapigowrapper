package rickandmorty

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func makePetition(options map[string]interface{}) (interface{}, error) {
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
		case map[string]int:
			if k == "params" {
				// Modify the URL, i.e. the endpoint
				integer := strconv.FormatInt(int64(v.(map[string]int)["integer"]), 10)
				endpoint = endpoint + integer
				// Delete the parameters
				delete(options, "params")
			}
		case []int:
			params := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.([]int))), ","), "[]")
			endpoint = endpoint + params
		default:
			err := sliceIntToString(v.([]int), ",")
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

	var response interface{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
