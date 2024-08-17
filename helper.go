package semprit

import (
	"reflect"
	"strconv"

	d "github.com/karincake/dodol"
	p "github.com/karincake/pentol"
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
	case vk >= reflect.Int && vk <= reflect.Int64:
		if rvs != "" {
			rvsVal, err := strconv.Atoi(rvs)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			var fvtemp reflect.Value
			if vk == reflect.Int8 {
				fvtemp = reflect.ValueOf(int8(0))
			} else if vk == reflect.Int16 {
				fvtemp = reflect.ValueOf(int16(0))
			} else if vk == reflect.Int32 {
				fvtemp = reflect.ValueOf(int32(0))
			} else if vk == reflect.Int64 {
				fvtemp = reflect.ValueOf(int64(0))
			} else {
				fvtemp = reflect.ValueOf(int(0))
			}
			if fvtemp.OverflowInt(int64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			}
			if vk == reflect.Int8 {
				fv.Set(reflect.ValueOf(p.Int8(int8(rvsVal))))
			} else if vk == reflect.Int16 {
				fv.Set(reflect.ValueOf(p.Int16(int16(rvsVal))))
			} else if vk == reflect.Int32 {
				fv.Set(reflect.ValueOf(p.Int32(int32(rvsVal))))
			} else if vk == reflect.Int64 {
				fv.Set(reflect.ValueOf(p.Int64(int64(rvsVal))))
			} else {
				fv.Set(reflect.ValueOf(p.Int(rvsVal)))
			}
		}
	case vk >= reflect.Uint && vk <= reflect.Uint64:
		if rvs != "" {
			rvsVal, err := strconv.ParseUint(rvs, 10, 64)
			if err != nil {
				return d.FieldError{Code: "convert-fail", Message: "can not convert \"" + ftName + "\" (value: " + rvs + ") into number"}
			}
			var fvtemp reflect.Value
			if vk == reflect.Uint8 {
				fvtemp = reflect.ValueOf(uint8(0))
			} else if vk == reflect.Uint16 {
				fvtemp = reflect.ValueOf(uint16(0))
			} else if vk == reflect.Uint32 {
				fvtemp = reflect.ValueOf(uint32(0))
			} else if vk == reflect.Uint64 {
				fvtemp = reflect.ValueOf(uint64(0))
			} else {
				fvtemp = reflect.ValueOf(uint(0))
			}
			if fvtemp.OverflowUint(uint64(rvsVal)) {
				return d.FieldError{Code: "value-overflow", Message: "value overflow for \"" + ftName + "\" (value: " + rvs + ")"}
			}
			if vk == reflect.Uint8 {
				fv.Set(reflect.ValueOf(p.Uint8(uint8(rvsVal))))
			} else if vk == reflect.Uint16 {
				fv.Set(reflect.ValueOf(p.Uint16(uint16(rvsVal))))
			} else if vk == reflect.Uint32 {
				fv.Set(reflect.ValueOf(p.Uint32(uint32(rvsVal))))
			} else if vk == reflect.Uint64 {
				fv.Set(reflect.ValueOf(p.Uint64(uint64(rvsVal))))
			} else {
				fv.Set(reflect.ValueOf(p.Uint(uint(rvsVal))))
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

// func caster(typeCode string)
