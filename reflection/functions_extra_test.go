package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func sp(input string) *string {
	return &input
}

func Test_AutoMapper_Field_Value_NilPointer(t *testing.T) {

	type Source struct {
		Info string `mapper:"info"`
	}

	type Target struct {
		Info string `mapper:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	var target Target
	fmt.Printf("target: %p\n", &target)

	err := AutoMapper(source, target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != target.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Field_Value_Value(t *testing.T) {

	type Source struct {
		Info string `mapper:"info"`
	}

	type Target struct {
		Info string `mapper:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != target.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Field_Value_Pointer(t *testing.T) {

	type Source struct {
		Info string `mapper:"info"`
	}

	type Target struct {
		Info *string `mapper:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != *target.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Field_Pointer_Value(t *testing.T) {

	type Source struct {
		Info *string `mapper:"info"`
	}

	type Target struct {
		Info string `mapper:"info"`
	}

	source := Source{
		Info: sp("Source.Info"),
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if *source.Info != target.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Field_Pointer_Pointer(t *testing.T) {

	type Source struct {
		Info *string `mapper:"info"`
	}

	type Target struct {
		Info *string `mapper:"info"`
	}

	source := Source{
		Info: sp("Source.Info"),
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if *source.Info != *target.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Nested_Value_Value(t *testing.T) {

	type NestedSource struct {
		Info string `mapper:"info"`
	}

	type NestedTarget struct {
		Info string `mapper:"info"`
	}

	type Source struct {
		Nested NestedSource `mapper:"nested"`
	}

	type Target struct {
		Nested NestedTarget `mapper:"nested"`
	}

	nested := NestedSource{
		Info: "Source.NestedSource.Info",
	}

	source := Source{
		Nested: nested,
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMapper_Nested_Value_Pointer(t *testing.T) {

	type NestedSource struct {
		Info string `mapper:"info"`
	}

	type NestedTarget struct {
		Info string `mapper:"info"`
	}

	type Source struct {
		Nested NestedSource `mapper:"nested"`
	}

	type Target struct {
		Nested *NestedTarget `mapper:"nested"`
	}

	nested := NestedSource{
		Info: "Source.NestedSource.Info",
	}

	source := Source{
		Nested: nested,
	}

	var target Target

	err := AutoMapper(source, &target)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMapper failed to set value while returning no error.")
		t.Fail()
	}
}

func TestAutoMap_Nested(t *testing.T) {

	type Gamma struct {
		Info string `automap:"info"`
	}

	type Beta struct {
		NestedTwice Gamma  `automap:"twice"`
		Info        string `automap:"info"`
	}

	type Alpha struct {
		NestedOnce Beta   `automap:"once"`
		Info       string `automap:"info"`
	}

	type Baz struct {
		Info string `automap:"info"`
	}

	type Bar struct {
		NestedTwice Baz    `automap:"twice"`
		Info        string `automap:"info"`
	}

	type Foo struct {
		NestedOnce Bar `automap:"once"`
		Info       int `automap:"info"`
	}

	g := Gamma{"Hello from Gamma!"}
	b := Beta{
		NestedTwice: g,
		Info:        "Hello from Beta!",
	}
	a := Alpha{
		NestedOnce: b,
		Info:       "Hello from Alpha!",
	}

	f, err := AutoMap[Foo](a)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println("Foo.Info:", f.Info)
		fmt.Println("Foo.NestedOnce.Info:", f.NestedOnce.Info)
		fmt.Println("Foo.NestedOnce.NestedTwice.Info:", f.NestedOnce.NestedTwice.Info)
	}
}

func AutoMap[T any](source any) (target T, err error) {

	// Peel the source until we reach the struct.
	sourceStruct := reflect.ValueOf(&source)
	for sourceStruct.Kind() == reflect.Pointer || sourceStruct.Kind() == reflect.Interface {
		sourceStruct = sourceStruct.Elem()
	}

	// Validate the input.
	if reflect.ValueOf(target).Kind() == reflect.Pointer {
		cause := fmt.Errorf("(target pointers are not supported)")
		return target, cause
	}
	if sourceStruct.Kind() != reflect.Struct {
		cause := fmt.Errorf("(source must be a struct, got: '%s')", sourceStruct.Kind().String())
		return target, cause
	}
	if reflect.ValueOf(target).Kind() != reflect.Struct {
		cause := fmt.Errorf("(target must be a struct, got: '%s')", reflect.ValueOf(target).Kind().String())
		return target, cause
	}

	targetStruct := reflect.ValueOf(&target).Elem()

	// Enter the recursive part.
	err = autoMap(sourceStruct, targetStruct)
	return target, err
}

func autoMap(s, t reflect.Value) error {

	describe("1.1: sourceStruct", s)
	describe("1.2: targetStruct", t)

	sourceMap := mapTagToFieldIndex("automap", s.Interface())
	targetMap := mapTagToFieldIndex("automap", t.Interface())

	for k := range sourceMap {

		sourceIndex := sourceMap[k]
		targetIndex, found := targetMap[k]
		if !found {
			continue
		}

		sourceField := s.Field(sourceIndex)
		targetField := t.Field(targetIndex)
		describe("2.1: sourceField", sourceField)
		describe("2.2: targetField", targetField)

		// Nothing to do when source is nil.
		if sourceField.Kind() == reflect.Pointer && sourceField.IsNil() {
			continue
		}

		err := compareKind(sourceField, targetField)
		if err != nil {
			return err
		}
		describe("7: target", sourceField)
		describe("8: target", targetField)

		// Match target kind by adding pointers.
		if targetField.Kind() == reflect.Pointer && sourceField.Kind() != reflect.Pointer {
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

		if sourceField.Kind() == reflect.Struct {
			err := autoMap(sourceField, targetField)
			if err != nil {
				return err
			}
			continue
		}

		if sourceField.Kind() == reflect.Pointer && targetField.Kind() != reflect.Pointer {
			sourceField = sourceField.Elem()
		}

		targetField.Set(sourceField)
	}

	return nil
}

func compareKind(source, target reflect.Value) error {

	describe("5.1: source", source)
	for source.Kind() == reflect.Pointer {
		source = source.Elem()
	}
	describe("5.2: source", source)

	// TODO: Extract setting nil targets into seperate function called before compareKind.
	describe("6.1: target", target)
	for target.Kind() == reflect.Pointer {
		if target.IsNil() {
			temp := reflect.New(reflect.TypeOf(source.Interface()))
			temp.Elem().Set(reflect.ValueOf(source.Interface()))
			describe("6.3: target", temp)
			target.Set(temp)
		} else {
			target = target.Elem()
		}
	}
	describe("6.2: target", target)

	if source.Kind() != target.Kind() {
		cause := fmt.Errorf(
			"(source and target kind mismatch - source: '%s', target: '%s')",
			source.Kind(), target.Kind(),
		)
		return cause
	}

	return nil
}
