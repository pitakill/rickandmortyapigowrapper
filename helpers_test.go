package rickandmorty

import "testing"

var slice = []string{
	"hello",
	"world",
	"hola",
	"mundo",
	"Allo",
	"Welt",
}

var sliceIntegers = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func Test_containsStringElementPresent(t *testing.T) {
	sliceToIterate := []string{
		"hello",
		"hola",
		"Allo",
		"world",
		"mundo",
		"Welt",
	}

	for _, element := range sliceToIterate {
		if !containsString(slice, element) {
			t.Errorf("The %s (%T) is not present in %v", element, element, slice)
		}
	}
}

func Test_containsStringElementNotPresent(t *testing.T) {
	sliceToIterate := []string{
		"These",
		"strings",
		"are",
		"not",
		"present",
		"in",
		"the",
		"variable",
		"slice",
	}

	for _, element := range sliceToIterate {
		if containsString(slice, element) {
			t.Errorf("The %s (%T) is present in %v", element, element, slice)
		}
	}
}

func Test_sliceIntToString(t *testing.T) {
	validResponse := "1,2,3,4,5,6,7,8,9,10"

	response := sliceIntToString(sliceIntegers, ",")

	if response != validResponse {
		t.Errorf("The valid string (%s) and the obtained (%v) differ", validResponse, sliceIntegers)
	}

	validResponse = "1-2-3-4-5-6-7-8-9-10"

	response = sliceIntToString(sliceIntegers, "-")

	if response != validResponse {
		t.Errorf("The valid string (%s) and the obtained (%v) differ", validResponse, sliceIntegers)
	}
}
