package kit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ToBool converts an interface to a bool type.
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

func ToBoolE(i interface{}) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		return b != 0, nil
	case int8:
		return b != 0, nil
	case int16:
		return b != 0, nil
	case int32:
		return b != 0, nil
	case int64:
		return b != 0, nil
	case uint:
		return b != 0, nil
	case uint8:
		return b != 0, nil
	case uint16:
		return b != 0, nil
	case uint32:
		return b != 0, nil
	case uint64:
		return b != 0, nil
	case float32:
		return b != 0, nil
	case float64:
		return b != 0, nil
	case time.Duration:
		return b != 0, nil
	case string:
		return strconv.ParseBool(i.(string))
	case json.Number:
		v, err := ToInt64E(b)
		if err == nil {
			return v != 0, nil
		}
		return false, fmt.Errorf("unable to convert %#v of type %T to bool", i, i)
	default:
		return false, fmt.Errorf("unable to convert %#v of type %T to bool", i, i)
	}
}

// ToInt converts an interface to an int type.
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

func ToIntE(i interface{}) (int, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return int(t), nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case time.Weekday:
		return int(t), nil
	case time.Month:
		return int(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to int64", i, i)
	case json.Number:
		return ToIntE(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to int", i, i)
	}
}

// ToInt8 converts an interface to an int8 type.
func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

func ToInt8E(i interface{}) (int8, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return int8(t), nil
	case int8:
		return int8(t), nil
	case int16:
		return int8(t), nil
	case int32:
		return int8(t), nil
	case int64:
		return int8(t), nil
	case uint:
		return int8(t), nil
	case uint8:
		return int8(t), nil
	case uint16:
		return int8(t), nil
	case uint32:
		return int8(t), nil
	case uint64:
		return int8(t), nil
	case float32:
		return int8(t), nil
	case float64:
		return int8(t), nil
	case time.Weekday:
		return int8(t), nil
	case time.Month:
		return int8(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to int8", i, i)
	case json.Number:
		return ToInt8E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to int8", i, i)
	}
}

// ToInt16 converts an interface to an int16 type.
func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

func ToInt16E(i interface{}) (int16, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return int16(t), nil
	case int8:
		return int16(t), nil
	case int16:
		return int16(t), nil
	case int32:
		return int16(t), nil
	case int64:
		return int16(t), nil
	case uint:
		return int16(t), nil
	case uint8:
		return int16(t), nil
	case uint16:
		return int16(t), nil
	case uint32:
		return int16(t), nil
	case uint64:
		return int16(t), nil
	case float32:
		return int16(t), nil
	case float64:
		return int16(t), nil
	case time.Weekday:
		return int16(t), nil
	case time.Month:
		return int16(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to int16", i, i)
	case json.Number:
		return ToInt16E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to int16", i, i)
	}
}

// ToInt32 converts an interface to an int32 type.
func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

func ToInt32E(i interface{}) (int32, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return int32(t), nil
	case int8:
		return int32(t), nil
	case int16:
		return int32(t), nil
	case int32:
		return int32(t), nil
	case int64:
		return int32(t), nil
	case uint:
		return int32(t), nil
	case uint8:
		return int32(t), nil
	case uint16:
		return int32(t), nil
	case uint32:
		return int32(t), nil
	case uint64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case time.Weekday:
		return int32(t), nil
	case time.Month:
		return int32(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to int32", i, i)
	case json.Number:
		return ToInt32E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to int32", i, i)
	}
}

// ToInt64 converts an interface to an int64 type.
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case uint8:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case time.Weekday:
		return int64(t), nil
	case time.Month:
		return int64(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return int64(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to int64", i, i)
	case json.Number:
		return ToInt64E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to int64", i, i)
	}
}

// ToUint converts an interface to an uint type.
func ToUint(i interface{}) uint {
	v, _ := ToUintE(i)
	return v
}

func ToUintE(i interface{}) (uint, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return uint(t), nil
	case int8:
		return uint(t), nil
	case int16:
		return uint(t), nil
	case int32:
		return uint(t), nil
	case int64:
		return uint(t), nil
	case uint:
		return uint(t), nil
	case uint8:
		return uint(t), nil
	case uint16:
		return uint(t), nil
	case uint32:
		return uint(t), nil
	case uint64:
		return uint(t), nil
	case float32:
		return uint(t), nil
	case float64:
		return uint(t), nil
	case time.Weekday:
		return uint(t), nil
	case time.Month:
		return uint(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint", i, i)
	case json.Number:
		return ToUintE(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint", i, i)
	}
}

// ToUint8 converts an interface to an uint8 type.
func ToUint8(i interface{}) uint8 {
	v, _ := ToUint8E(i)
	return v
}

func ToUint8E(i interface{}) (uint8, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return uint8(t), nil
	case int8:
		return uint8(t), nil
	case int16:
		return uint8(t), nil
	case int32:
		return uint8(t), nil
	case int64:
		return uint8(t), nil
	case uint:
		return uint8(t), nil
	case uint8:
		return uint8(t), nil
	case uint16:
		return uint8(t), nil
	case uint32:
		return uint8(t), nil
	case uint64:
		return uint8(t), nil
	case float32:
		return uint8(t), nil
	case float64:
		return uint8(t), nil
	case time.Weekday:
		return uint8(t), nil
	case time.Month:
		return uint8(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint8", i, i)
	case json.Number:
		return ToUint8E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint8", i, i)
	}
}

// ToUint16 converts an interface to an uint16 type.
func ToUint16(i interface{}) uint16 {
	v, _ := ToUint16E(i)
	return v
}

func ToUint16E(i interface{}) (uint16, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return uint16(t), nil
	case int8:
		return uint16(t), nil
	case int16:
		return uint16(t), nil
	case int32:
		return uint16(t), nil
	case int64:
		return uint16(t), nil
	case uint:
		return uint16(t), nil
	case uint8:
		return uint16(t), nil
	case uint16:
		return uint16(t), nil
	case uint32:
		return uint16(t), nil
	case uint64:
		return uint16(t), nil
	case float32:
		return uint16(t), nil
	case float64:
		return uint16(t), nil
	case time.Weekday:
		return uint16(t), nil
	case time.Month:
		return uint16(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint16", i, i)
	case json.Number:
		return ToUint16E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint16", i, i)
	}
}

// ToUint32 converts an interface to an uint32 type.
func ToUint32(i interface{}) uint32 {
	v, _ := ToUint32E(i)
	return v
}

func ToUint32E(i interface{}) (uint32, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return uint32(t), nil
	case int8:
		return uint32(t), nil
	case int16:
		return uint32(t), nil
	case int32:
		return uint32(t), nil
	case int64:
		return uint32(t), nil
	case uint:
		return uint32(t), nil
	case uint8:
		return uint32(t), nil
	case uint16:
		return uint32(t), nil
	case uint32:
		return uint32(t), nil
	case uint64:
		return uint32(t), nil
	case float32:
		return uint32(t), nil
	case float64:
		return uint32(t), nil
	case time.Weekday:
		return uint32(t), nil
	case time.Month:
		return uint32(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint32", i, i)
	case json.Number:
		return ToUint32E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint32", i, i)
	}
}

// ToUint64 converts an interface to an uint64 type.
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

func ToUint64E(i interface{}) (uint64, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return uint64(t), nil
	case int8:
		return uint64(t), nil
	case int16:
		return uint64(t), nil
	case int32:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case uint:
		return uint64(t), nil
	case uint8:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint32:
		return uint64(t), nil
	case uint64:
		return uint64(t), nil
	case float32:
		return uint64(t), nil
	case float64:
		return uint64(t), nil
	case time.Weekday:
		return uint64(t), nil
	case time.Month:
		return uint64(t), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(t), 0, 0)
		if err == nil {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint64", i, i)
	case json.Number:
		return ToUint64E(string(t))
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to uint64", i, i)
	}
}

// ToFloat32 converts an interface to a float32 type.
func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

func ToFloat32E(i interface{}) (float32, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return float32(t), nil
	case int8:
		return float32(t), nil
	case int16:
		return float32(t), nil
	case int32:
		return float32(t), nil
	case int64:
		return float32(t), nil
	case uint:
		return float32(t), nil
	case uint8:
		return float32(t), nil
	case uint16:
		return float32(t), nil
	case uint32:
		return float32(t), nil
	case uint64:
		return float32(t), nil
	case float32:
		return float32(t), nil
	case float64:
		return float32(t), nil
	case time.Weekday:
		return float32(t), nil
	case time.Month:
		return float32(t), nil
	case string:
		v, err := strconv.ParseFloat(t, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to float32", i, i)
	case json.Number:
		if v, err := t.Float64(); err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to float32", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to float32", i, i)
	}
}

// ToFloat64 converts an interface to a float64 type.
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

func ToFloat64E(i interface{}) (float64, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return float64(t), nil
	case time.Weekday:
		return float64(t), nil
	case time.Month:
		return float64(t), nil
	case string:
		v, err := strconv.ParseFloat(t, 32)
		if err == nil {
			return float64(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to float64", i, i)
	case json.Number:
		if v, err := t.Float64(); err == nil {
			return float64(v), nil
		}
		return 0, fmt.Errorf("unable to convert %#v of type %T to float64", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to convert %#v of type %T to float64", i, i)
	}
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
