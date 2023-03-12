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

func ListStructTags(key string, input any) []string {

	t := reflect.TypeOf(input)
	if t == nil {
		return nil
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	var tags []string
	for i := 0; i < t.NumField(); i++ {
		full := t.Field(i).Tag.Get(key)
		split := strings.Split(full, ",")
		if split[0] == "" {
			continue
		}
		tags = append(tags, split[0])
	}
	return tags
}
