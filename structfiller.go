package semprit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	te "github.com/karincake/tempe/error"
	"gorm.io/datatypes"
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

func FDFromHttp(container any, r *http.Request) error {
	// identiy value and loop if its pointer until reaches non pointer
	contV := reflect.ValueOf(container)

	// loop until we get what kind lays behind the input
	for contV.Kind() == reflect.Pointer || contV.Kind() == reflect.Interface {
		contV = contV.Elem()
	}

	// non struct cant be validated
	if contV.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	// check each field
	// contT := reflect.TypeOf(contV.Interface()) // keep this for now
	contT := contV.Type()
	for i := 0; i < contV.NumField(); i++ {
		// identify field type and value of the field
		ft := contT.Field(i)
		fv := contV.Field(i)
		if !fv.CanSet() {
			continue
		}

		key := ft.Tag.Get("json")
		if key != "" {
			keys := strings.Split(key, ",")
			if keys[0] != "" {
				key = keys[0]
			} else {
				key = ft.Name
			}
		} else {
			key = ft.Name
		}

		fName := ft.Name
		rv := r.FormValue(key)
		ftName := ft.Type.String()
		ftNameClean := strings.Trim(ftName, "*")
		switch {
		case ftName == "string":
			fv.SetString(rv)
		case ftName == "*string" && !fv.IsNil():
			reflect.Indirect(fv).SetString(rv)
		case ftName == "bool":
			if rv == "true" {
				fv.SetBool(true)
			} else {
				fv.SetBool(false)
			}
		case ftName == "*bool" && !fv.IsNil():
			if rv == "true" {
				reflect.Indirect(fv).SetBool(true)
			} else if rv == "false" {
				reflect.Indirect(fv).SetBool(false)
			}
		case len(ftNameClean) >= 3 && ftNameClean[:3] == "int": // bundle in one
			if rv != "" {
				rvVal, err := strconv.Atoi(rv)
				if err != nil {
					return fmt.Errorf("can not convert %s into number", fName)
				} else {
					if ftName[:1] != "*" {
						if fv.OverflowInt(int64(rvVal)) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							fv.SetInt(int64(rvVal))
						}
					} else if !fv.IsNil() {
						if reflect.Indirect(fv).OverflowInt(int64(rvVal)) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							reflect.Indirect(fv).SetInt(int64(rvVal))
						}
					}
				}
			} else if ftName[:1] != "*" {
				fv.SetInt(0)
			}
		case len(ftNameClean) >= 5 && ftNameClean[:5] == "float": // bundle in one
			if rv != "" {
				floatType := 32
				if ftName == "float64" {
					floatType = 64
				}
				rvVal, err := strconv.ParseFloat(rv, floatType)
				if err != nil {
					return fmt.Errorf("can not convert %s into number", fName)
				} else {
					if ftName[:1] != "*" {
						if fv.OverflowFloat(rvVal) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							fv.SetFloat(rvVal)
						}
					} else if !fv.IsNil() {
						if reflect.Indirect(fv).OverflowFloat(rvVal) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							reflect.Indirect(fv).SetFloat(rvVal)
						}
					}
				}
			} else if ftName[:1] != "*" {
				fv.SetFloat(0)
			}
		}
	}
	return nil
}

func QPFromHttp(container any, url url.URL) error {
	cV := reflect.ValueOf(container).Elem()
	for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
		cV = cV.Elem()
	}

	cT := cV.Type()
	values := url.Query()
	for i := 0; i < cV.NumField(); i++ {
		fieldV := cV.Field(i)
		fieldT := cT.Field(i)

		if !fieldV.CanSet() {
			continue
		}

		key := keyOrJsonTag(fieldT.Name, fieldT.Tag.Get("json"))

		vals, ok := values[key]
		if !ok {
			continue
		}

		switch fieldV.Interface().(type) {
		case bool, *bool:
			var v bool
			fieldVS := strings.ToLower(vals[0])
			if fieldVS == "true" || fieldVS == "yes" || fieldVS == "1" {
				v = true
			} else if fieldVS == "false" || fieldVS == "no" || fieldVS == "0" {
				v = false
			} else {
				return te.XError{Code: "bool-parse-fail", Message: "failed to parse bool value into field " + key}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&v))
			} else {
				fieldV.Set(reflect.ValueOf(v))
			}
		case string, *string:
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&vals[0]))
			} else {
				fieldV.Set(reflect.ValueOf(vals[0]))
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
			if valInt, err := strconv.Atoi(vals[0]); err != nil {
				return te.XError{Code: "int-parse-fail", Message: "failed to parse int value into field " + key}
			} else {
				fieldV.Set(intToVal(valInt, fieldV))
			}
		case float64, *float64:
			strFloat, err := strconv.ParseFloat(vals[0], 64)
			if err != nil {
				return te.XError{Code: "float32-parse-fail", Message: "failed to parse float32 value into field " + key}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat))
			}
		case float32, *float32:
			strFloat, err := strconv.ParseFloat(vals[0], 32)
			if err != nil {
				return te.XError{Code: "float64-parse-fail", Message: "failed to parse float64 value into field " + key}
			}
			strFloat32 := float32(strFloat)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat32))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat32))
			}
		case []string, *[]string:
			fieldV.Set(reflect.ValueOf(&vals))
		case datatypes.Date, *datatypes.Date:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				return te.XError{Code: "gormDate-parse-fail", Message: "failed to gorm-date value into field " + key}
			}
			date := datatypes.Date(time)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&date))
			} else {
				fieldV.Set(reflect.ValueOf(date))
			}
		case time.Time, *time.Time:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				return te.XError{Code: "time-parse-fail", Message: "failed to time value into field " + key}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&time))
			} else {
				fieldV.Set(reflect.ValueOf(time))
			}
		// TODO: make any *[]int as a function
		case *[]int8:
			failed := false
			valX := []int8{}
			for idx, val := range vals {
				if valInt, err := strconv.Atoi(val); err != nil {
					failed = true
					return te.XError{Code: "[]int8-parse-fail", Message: "failed to parse []uint8 value for field " + key + " at index " + strconv.Itoa(idx)}
				} else {
					valX = append(valX, int8(valInt))
				}
			}
			if !failed {
				fieldV.Set(reflect.ValueOf(valX))
			}

		}
	}

	return nil
}
