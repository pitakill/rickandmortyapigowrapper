package rickandmorty

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
)

func Test_makePetition(t *testing.T) {
	response, err := makePetition(nil)
	if err != nil {
		t.Error(err)
	}

	endpointsStructured := new(Endpoints)

	_ = mapstructure.Decode(response, &endpointsStructured)

	data, err := readFile("test-data/api-info.json")
	if err != nil {
		t.Error(err)
	}

	endpoints := new(Endpoints)

	json.Unmarshal(data, &endpoints)

	comparation := cmp.Equal(endpoints, endpointsStructured)

	if !comparation {
		t.Error("The response from makePetition was:")
		t.Error(endpointsStructured)
		t.Error("The data against is being run this test is:")
		t.Error(endpoints)
	}
}

func readFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}
