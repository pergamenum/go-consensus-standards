package reflection

import (
	"fmt"
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

// MapTagToType extracts the field tag and type belonging to each field marked with a given struct tag key.
func MapTagToType(tagKey string, inputStruct any) map[string]string {

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
		// Typically used to show that a field is supposed to be ignored.
		// For example: Age int `json:"-"`
		if strings.TrimSpace(tag) == "-" {
			continue
		}

		tn := t.Field(i).Type.Name()
		m[tag] = tn
	}

	return m
}

func AutoMapper(source, target any) error {
	fmt.Printf("target: %p\n", &target)
	println("target pointer:", reflect.ValueOf(&target).Pointer())
	println("target pointer:", reflect.ValueOf(&target).Elem().Pointer())
	println("target pointer:", reflect.ValueOf(&target).Elem().Elem().Pointer())
	// Check that target struct is passed by pointer.
	if reflect.ValueOf(&target).Elem().Elem().Kind() != reflect.Pointer {
		return fmt.Errorf("target must be passed by pointer")
	}

	target = reflect.New(reflect.ValueOf(&target).Elem().Elem().Type().Elem()).Interface()

	// If multiple fields have the same tag name, then only the index of the last field will be included in the map.
	//targetMap := mapTagToFieldIndex("mapper", target)
	//sourceMap := mapTagToFieldIndex("mapper", source)

	// Peel off pointer and interface wrapper from 'any'.
	targetField := reflect.ValueOf(&target).Elem().Elem()
	sourceField := reflect.ValueOf(&source).Elem().Elem()
	describe("1.1: targetField:", targetField)
	describe("1.2: sourceField:", sourceField)

	// If multiple fields have the same tag name, then only the index of the last field will be included in the map.
	targetMap := mapTagToFieldIndex("mapper", target)
	sourceMap := mapTagToFieldIndex("mapper", source)

	// Check that target struct is passed by pointer.
	//if targetField.Kind() != reflect.Pointer {
	//	return fmt.Errorf("target must be passed by pointer")
	//}
	// Overwrite whatever the old pointer was with a fresh one.
	// This is supposed to simplify cases where target is a nil pointer to a struct.
	//targetField = reflect.New(targetField.Type().Elem())

	//println("targetField.IsValid():", targetField.IsValid())
	//println("targetField.IsNil():", targetField.IsNil())

	// Peel off pointer and check that it is a struct.
	targetField = targetField.Elem()
	println("targetField.IsValid():", targetField.IsValid())

	describe("1.3: targetField:", targetField)

	if targetField.Kind() == reflect.Interface {
		targetField = targetField.Elem()
	}

	if targetField.Kind() != reflect.Struct {
		return fmt.Errorf("target must be a struct")
	}

	for k := range sourceMap {

		sourceIndex := sourceMap[k]
		targetIndex, found := targetMap[k]
		if !found {
			continue
		}

		// Pick the relevant field of the struct.
		targetField = targetField.Field(targetIndex)
		describe("2.1: targetField:", targetField)

		if targetField.Kind() == reflect.Pointer {
			println(targetField.IsNil())
		}

		// Peel off pointer and interface wrapper from 'any', then pick the struct field.

		// Peel off any pointers.
		// Consider the need for checking against reflect.Interface.
		for sourceField.Kind() == reflect.Pointer {
			sourceField = sourceField.Elem()
		}
		sourceField = sourceField.Field(sourceIndex)
		describe("2.2: sourceField:", sourceField)

		if sourceField.Kind() == reflect.Struct {

			//if targetField.CanAddr() {
			//	Only a Value that has a pre-existing pointer can be Addressed...?
			//	That is, only a Value that has been previously peeled with .Elem()
			//targetField = targetField.Addr()
			//} else {
			//	newPointer := reflect.New(reflect.TypeOf(targetField.Interface()))
			//	newPointer.Elem().Set(reflect.ValueOf(targetField.Interface()))
			//	targetField = newPointer
			//}

			describe("X1:", targetField)

			if targetField.Kind() != reflect.Pointer {
				targetField = targetField.Addr()
				describe("X2:", targetField)
			}

			//err := AutoMapper(sourceField.Interface(), targetField.Addr().Interface())
			err := AutoMapper(sourceField.Interface(), targetField.Interface())
			if err != nil {
				fmt.Println(err)
			} else {
				continue
			}
		}

		// Peel off any pointers.
		// Consider the need for checking against reflect.Interface.
		for sourceField.Kind() == reflect.Pointer {
			sourceField = sourceField.Elem()
		}
		describe("2.3: sourceField:", sourceField)

		//if sourceField.Kind() == reflect.Struct {
		//	aa := sourceField.Type()
		//	x := reflect.ValueOf(&source).Elem().Field(sourceIndex)
		//	y := reflect.TypeOf(&source).Elem().Field(sourceIndex)
		//	xI := x.Interface()
		//	println(xI)
		//	println(x.String())
		//	println(y.Type.String())
		//	tempy(xI, aa)
		//	//a, _ := AutoMapper[aa, xI](xI)
		//	//println(a)
		//	sourceField = sourceField.Elem()
		//}

		// Match target kind by adding pointers.
		if targetField.Kind() == reflect.Pointer {
			if sourceField.CanAddr() {
				// Only a Value that has a pre-existing pointer can be Addressed...?
				// That is, only a Value that has been previously peeled with .Elem()
				sourceField = sourceField.Addr()
			} else {
				newPointer := reflect.New(reflect.TypeOf(sourceField.Interface()))
				newPointer.Elem().Set(reflect.ValueOf(sourceField.Interface()))
				sourceField = newPointer
			}
		}

		// Set the field
		targetField.Set(sourceField)
	}
	return nil
}

func describe(name string, input reflect.Value) {
	description := fmt.Sprintf("Type: %v, Kind: %v", input.Type(), input.Kind())
	fmt.Printf("[%s] -> %s\n", name, description)
}

func mapTagToFieldIndex(tagKey string, inputStruct any) map[string]int {

	t := reflect.TypeOf(inputStruct)
	if t == nil {
		return nil
	}
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	m := map[string]int{}
	for i := 0; i < t.NumField(); i++ {

		full := t.Field(i).Tag.Get(tagKey)
		split := strings.Split(full, ",")
		tag := split[0]
		if strings.TrimSpace(tag) == "" {
			continue
		}
		// Typically used to show that a field is supposed to be ignored.
		// For example: Age int `json:"-"`
		if strings.TrimSpace(tag) == "-" {
			continue
		}

		m[tag] = i
	}

	return m
}
