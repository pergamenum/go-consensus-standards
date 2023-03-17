package reflection

import (
	"reflect"
	"runtime"
	"strings"
)

// GetTypeName returns the input type's name.
func GetTypeName(input any) string {
	return reflect.TypeOf(input).Name()
}

// GetFunctionName returns the name of the calling function.
//
// Returns an empty string upon failure.
func GetFunctionName() string {

	c, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	full := runtime.FuncForPC(c).Name()
	// Generic functions have [...] appended to their name.
	trimmed := strings.ReplaceAll(full, "[...]", "")
	last := trimmed[strings.LastIndex(trimmed, ".")+1:]

	return last
}

// MapFieldToType extracts the tag name and type belonging to each field marked with a given struct tag key.
func MapFieldToType(tagKey string, inputStruct any) map[string]string {

	t := reflect.TypeOf(inputStruct)
	if t == nil {
		return nil
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	m := map[string]string{}
	for i := 0; i < t.NumField(); i++ {

		full := t.Field(i).Tag.Get(tagKey)
		split := strings.Split(full, ",")
		tag := split[0]
		if strings.TrimSpace(tag) == "" {
			continue
		}

		tn := t.Field(i).Type.Name()
		m[tag] = tn
	}

	return m
}
