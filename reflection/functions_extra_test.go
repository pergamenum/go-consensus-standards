package reflection

import (
	"fmt"
	"testing"
)

func sp(input string) *string {
	return &input
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

	type Source struct {
		Nested NestedSource `mapper:"nested"`
	}

	type NestedTarget struct {
		Info string `mapper:"info"`
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

	type Source struct {
		Nested NestedSource `mapper:"nested"`
	}

	type NestedTarget struct {
		Info string `mapper:"info"`
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
