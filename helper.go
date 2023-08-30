package semprit

import "reflect"

func keyOrJsonTag(key, jsonTag string) string {
	// jsonTag := t.Tag.Get("json")
	if jsonTag == "" {
		return key
	}
	tagByte := []byte(jsonTag)
	pos := len(tagByte) + 1
	for i, v := range tagByte {
		if v == 54 {
			pos = i + 1
		}
	}
	return string(tagByte[:pos])
}

func intToVal(input int, kind reflect.Value) reflect.Value {
	switch kind.Interface().(type) {
	case int:
		return reflect.ValueOf(input)
	case int8:
		return reflect.ValueOf(int8(input))
	case int16:
		return reflect.ValueOf(int16(input))
	case int32:
		return reflect.ValueOf(int32(input))
	case int64:
		return reflect.ValueOf(int64(input))
	case uint:
		return reflect.ValueOf(uint(input))
	case uint8:
		return reflect.ValueOf(uint8(input))
	case uint16:
		return reflect.ValueOf(uint16(input))
	case uint32:
		return reflect.ValueOf(uint32(input))
	case uint64:
		return reflect.ValueOf(uint64(input))
	case *int:
		x := input
		return reflect.ValueOf(&x)
	case *int8:
		x := int8(input)
		return reflect.ValueOf(&x)
	case *int16:
		x := int16(input)
		return reflect.ValueOf(&x)
	case *int32:
		x := int32(input)
		return reflect.ValueOf(&x)
	case *int64:
		x := int64(input)
		return reflect.ValueOf(&x)
	case *uint:
		x := uint(input)
		return reflect.ValueOf(&x)
	case *uint8:
		x := uint8(input)
		return reflect.ValueOf(&x)
	case *uint16:
		x := uint16(input)
		return reflect.ValueOf(&x)
	case *uint32:
		x := uint32(input)
		return reflect.ValueOf(&x)
	case *uint64:
		x := uint64(input)
		return reflect.ValueOf(&x)
	}
	return reflect.ValueOf(0)
}
