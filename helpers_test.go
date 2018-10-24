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
