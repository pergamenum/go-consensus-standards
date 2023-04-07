package reflection

import (
	"fmt"
	"reflect"
)

func AutoMap[T any](source any) (target T, err error) {

	// Peel the source until we reach the struct.
	sourceStruct := reflect.ValueOf(&source)
	for sourceStruct.Kind() == reflect.Pointer || sourceStruct.Kind() == reflect.Interface {
		sourceStruct = sourceStruct.Elem()
	}

	targetKind := reflect.ValueOf(target).Kind()

	// Validate the input.
	if targetKind == reflect.Pointer {
		cause := fmt.Errorf("(pointers are not allowed as a type parameter)")
		return target, cause
	}
	if sourceStruct.Kind() != reflect.Struct {
		cause := fmt.Errorf("(source must be a struct, got: '%s')", sourceStruct.Kind().String())
		return target, cause
	}
	if targetKind != reflect.Struct {
		cause := fmt.Errorf("(target must be a struct, got: '%s')", targetKind.String())
		return target, cause
	}

	targetStruct := reflect.ValueOf(&target).Elem()

	// Enter the recursive part.
	err = autoMap(sourceStruct, targetStruct)
	return target, err
}

func autoMap(s, t reflect.Value) error {

	sourceMap := mapTagToFieldIndex("automap", s.Interface())
	targetMap := mapTagToFieldIndex("automap", t.Interface())

	for k := range sourceMap {

		sourceIndex := sourceMap[k]
		targetIndex, found := targetMap[k]
		if !found {
			continue
		}

		// Pick out the matching fields from the structs.
		sourceField := s.Field(sourceIndex)
		targetField := t.Field(targetIndex)

		// Peel off pointers to the source struct's current field.
		for sourceField.Kind() == reflect.Pointer && !sourceField.IsNil() {
			sourceField = sourceField.Elem()
		}

		// Nothing to do when the source is nil.
		if nillable(sourceField) && sourceField.IsNil() {
			continue
		}

		// Prepare the target field when it's nil and needs to be used.
		if nillable(targetField) && targetField.IsNil() {
			targetField.Set(reflect.New(targetField.Type().Elem()))
		}

		// Peel off pointers to the target struct's current field.
		for targetField.Kind() == reflect.Pointer && !targetField.IsNil() {
			targetField = targetField.Elem()
		}

		// Assert that the fields match.
		if sourceField.Kind() != targetField.Kind() {
			cause := fmt.Errorf(
				"(source and target kind mismatch - source: '%s', target: '%s')",
				sourceField.Kind(), targetField.Kind(),
			)
			return cause
		}

		// Recurse into nested structs.
		if sourceField.Kind() == reflect.Struct {
			err := autoMap(sourceField, targetField)
			if err != nil {
				return err
			}
			continue
		}

		targetField.Set(sourceField)
	}

	return nil
}

func nillable(input reflect.Value) (nillable bool) {

	switch input.Kind() {
	case reflect.Chan:
		nillable = true
	case reflect.Func:
		nillable = true
	case reflect.Interface:
		nillable = true
	case reflect.Map:
		nillable = true
	case reflect.Pointer:
		nillable = true
	case reflect.Slice:
		nillable = true
	}

	return nillable
}
