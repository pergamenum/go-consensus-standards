package types

import (
	"fmt"
	"strings"
	"time"
)

func AssertAny(value *any, t string) error {

	if value == nil {
		return fmt.Errorf("(pointer was nil)")
	}

	v := *value

	if v == nil {
		return fmt.Errorf("(derefence points to nil)")
	}

	f := func(result any, ok bool) error {
		if !ok {
			return fmt.Errorf("(value '%v' is not a valid '%s')", v, t)
		} else {
			*value = result
			return nil
		}
	}

	switch strings.ToLower(t) {

	case "bool":
		result, ok := v.(bool)
		return f(result, ok)

	case "string":
		result, ok := v.(string)
		return f(result, ok)

	case "int":
		result, ok := v.(int)
		return f(result, ok)

	case "int8":
		result, ok := v.(int8)
		return f(result, ok)

	case "int16":
		result, ok := v.(int16)
		return f(result, ok)

	case "int32":
		result, ok := v.(int32)
		return f(result, ok)

	case "int64":
		result, ok := v.(int64)
		return f(result, ok)

	case "uint":
		result, ok := v.(uint)
		return f(result, ok)

	case "uint8":
		result, ok := v.(uint8)
		return f(result, ok)

	case "uint16":
		result, ok := v.(uint16)
		return f(result, ok)

	case "uint32":
		result, ok := v.(uint32)
		return f(result, ok)

	case "uint64":
		result, ok := v.(uint64)
		return f(result, ok)

	case "byte":
		result, ok := v.(byte)
		return f(result, ok)

	case "rune":
		result, ok := v.(rune)
		return f(result, ok)

	case "float32":
		result, ok := v.(float32)
		return f(result, ok)

	case "float64":
		result, ok := v.(float64)
		return f(result, ok)

	case "complex64":
		result, ok := v.(complex64)
		return f(result, ok)

	case "complex128":
		result, ok := v.(complex128)
		return f(result, ok)

	case "time":
		result, ok := v.(time.Time)
		return f(result, ok)

	default:
		return fmt.Errorf("(unsupported type '%s' with value '%v')", t, value)
	}
}
