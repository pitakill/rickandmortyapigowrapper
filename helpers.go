package rickandmorty

import (
	"fmt"
	"strings"
)

func containsString(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func sliceIntToString(slice []int, join string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), join), "[]")
}
