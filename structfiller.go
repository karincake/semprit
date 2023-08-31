package semprit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	te "github.com/karincake/tempe/error"
)

// Fill struct with io.reader content, desired format is json
func IOReaderJson(container any, input io.Reader) error {
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&container)
	if err != nil {
		cv := reflect.ValueOf(container)
		for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
			cv = cv.Elem()
		}
		structName := cv.Type().Name()
		return te.XError{
			Code:        "parse-fail",
			Message:     "failed to parse input, error: " + err.Error(),
			ExpectedVal: "value of " + structName,
		}
	}

	return nil
}

// Fill struct with form-data content, desired format is key-val pairs
func HttpFormData(container any, r *http.Request) error {
	// identiy value and loop if its pointer until reaches non pointer
	cv := reflect.ValueOf(container)

	// loop until we get what kind lays behind the input
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	// non struct cant be filled
	if cv.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	// check each field
	ct := cv.Type()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		rv := r.PostFormValue(key)
		if rv == "" {
			// try once more if fail, mostly not called tho
			r.ParseForm()
			rv = r.FormValue(key)
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, rv)
		if err != nil {
			return err
		}
	}
	return nil
}

// Fill struct with url encoded content, desired format is url key-val pairs
func UrlQueryParam(container any, url url.URL) error {
	// identiy value and loop if its pointer until reaches non pointer
	cv := reflect.ValueOf(container)

	// loop until we get what kind lays behind the input
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	// non struct cant be filled
	if cv.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	ct := cv.Type()
	values := url.Query()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		vals, ok := values[key]
		if !ok {
			continue
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, vals[0])
		if err != nil {
			return err
		}
	}

	return nil
}
