package semprit

import (
	"reflect"
	"strconv"

	d "github.com/karincake/dodol"
)

func keyOrJsonTag(key, jsonTag string) string {
	if jsonTag == "" {
		return key
	}
	tagByte := []byte(jsonTag)
	pos := len(tagByte)
	for i, v := range tagByte {
		if v == 44 {
			pos = i
		}
	}
	return string(tagByte[:pos])
}

func reflectValueFiller(fv reflect.Value, vk reflect.Kind, ftName, rvs string) error {
	switch {
	case vk == reflect.String:
		fv.SetString(rvs)
	case vk == reflect.Bool:
		if rvs == "true" || rvs == "yes" || rvs == "1" {
			fv.SetBool(true)
		} else if rvs == "false" || rvs == "no" || rvs == "0" {
			fv.SetBool(false)
		}
	case vk >= reflect.Uint && vk <= reflect.Uint64:
		if rvs != "" {
			rvsVal, err := strconv.ParseUint(rvs, 10, 64)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowUint(uint64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				fv.SetUint(uint64(rvsVal))
			}
		}
	case vk >= reflect.Int && vk <= reflect.Int64:
		if rvs != "" {
			rvsVal, err := strconv.Atoi(rvs)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowInt(int64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				fv.SetInt(int64(rvsVal))
			}
		}
	case vk >= reflect.Float32 && vk <= reflect.Float64:
		if rvs != "" {
			floatType := 32
			if ftName == "float64" {
				floatType = 64
			}
			rvsVal, err := strconv.ParseFloat(rvs, floatType)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowFloat(rvsVal) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				fv.SetFloat(rvsVal)
			}
		}
	}

	return nil
}

func reflectPointerValueFiller(fv reflect.Value, vk reflect.Kind, ftName, rvs string) error {
	switch {
	case vk == reflect.String:
		fv.Set(reflect.ValueOf(&rvs))
	case vk == reflect.Bool:
		if rvs == "true" || rvs == "yes" || rvs == "1" {
			rvsx := true
			fv.Set(reflect.ValueOf(&rvsx))
		} else if rvs == "false" || rvs == "no" || rvs == "0" {
			rvsx := false
			fv.Set(reflect.ValueOf(&rvsx))
		}
	case vk >= reflect.Uint && vk <= reflect.Uint64:
		if rvs != "" {
			rvsVal, err := strconv.ParseUint(rvs, 10, 64)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowUint(uint64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				rvsx := uint64(rvsVal)
				fv.Set(reflect.ValueOf(&rvsx))
			}
		}
	case vk >= reflect.Int && vk <= reflect.Int64:
		if rvs != "" {
			rvsVal, err := strconv.Atoi(rvs)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowInt(int64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				rvsx := int64(rvsVal)
				fv.Set(reflect.ValueOf(&rvsx))
			}
		}
	case vk >= reflect.Float32 && vk <= reflect.Float64:
		if rvs != "" {
			floatType := 32
			if ftName == "float64" {
				floatType = 64
			}
			rvsVal, err := strconv.ParseFloat(rvs, floatType)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			if fv.OverflowFloat(rvsVal) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			} else {
				fv.Set(reflect.ValueOf(&rvsVal))
			}
		}
	}

	return nil
}
