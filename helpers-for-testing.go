package rickandmorty

import (
	"io/ioutil"
	"os"
	"reflect"

	"github.com/google/go-cmp/cmp"
)

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

func sliceEmptyNullReturnTrue() cmp.Option {
	alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })

	// This option handles slices and maps of any type.
	opt := cmp.FilterValues(func(x, y interface{}) bool {
		vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
		return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
			(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
			(vx.Len() == 0 && vy.Len() == 0)
	}, alwaysEqual)

	return opt
}
