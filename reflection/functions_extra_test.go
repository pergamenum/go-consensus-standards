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

type Alpha struct {
	NestedOnce Beta `automap:"nested"`
	//Info string `automap:"info"`
}

type Beta struct {
	NestedTwice Gamma
	//Info string `automap:"info"`
}

type Gamma struct {
	Info string `automap:"info"`
}

type Foo struct {
	NestedOnce Bar `automap:"nested"`
	//Info   string `automap:"info"`
}

type Bar struct {
	NestedTwice Baz
	//Info string `automap:"info"`
}

type Baz struct {
	Info string `automap:"info"`
}

func TestAutoMap(t *testing.T) {

	g := Gamma{"Hello from Gamma!"}
	b := Beta{g}
	a := Alpha{b}

	f, err := AutoMap[Foo](a)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
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

	sourceField := s.Field(0)
	targetField := t.Field(0)
	describe("2.1: sourceField", sourceField)
	describe("2.2: targetField", targetField)

	if sourceField.Kind() != targetField.Kind() {
		cause := fmt.Errorf(
			"(source and target kind mismatch - source: '%s', target: '%s')",
			sourceField.Kind(), targetField.Kind(),
		)
		return cause
	}

	if sourceField.Kind() == reflect.Struct {
		err := autoMap(sourceField, targetField)
		if err != nil {
			return err
		}
		return nil
	}

	targetField.Set(sourceField)

	return nil
}
