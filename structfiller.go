package semprit

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	te "github.com/karincake/tempe/error"
)

func JsonFromIOReader(container any, input io.Reader) error {
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&container)
	if err != nil {
		cV := reflect.ValueOf(container)
		for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
			cV = cV.Elem()
		}
		structName := cV.Type().Name()
		return te.XError{
			Code:        "parse-fail",
			Message:     "failed to parse input, error:" + err.Error(),
			ExpectedVal: fmt.Sprintf("value of %v", structName),
		}
	}

	return nil
}
