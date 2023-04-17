package types

import (
	"fmt"
	"reflect"

	r "github.com/pergamenum/go-consensus-standards/reflection"
)

type Update map[string]any

func NewUpdate(input any) (Update, error) {

	v := reflect.ValueOf(input)
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}

	if r.Nillable(v) && v.IsNil() {
		cause := fmt.Errorf("(invalid: 'input was nil pointer')")
		return nil, cause
	}

	if v.Kind() != reflect.Struct {
		cause := fmt.Errorf("(invalid: 'input was not a struct')")
		return nil, cause
	}

	t := v.Type()
	update := Update{}
	for i := 0; i < t.NumField(); i++ {

		key := t.Field(i).Tag.Get("update")
		if key == "" {
			continue
		}

		val := v.Field(i)

		for val.Kind() == reflect.Pointer && !val.IsNil() {
			val = val.Elem()
		}

		if r.Nillable(val) && val.IsNil() {
			continue
		}

		if !val.CanInterface() {
			continue
		}

		update[key] = val.Interface()
	}

	return update, nil
}
