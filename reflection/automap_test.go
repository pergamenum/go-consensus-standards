package reflection

import (
	"fmt"
	"testing"
)

func Test_AutoMap_Field_Value_Value(t *testing.T) {

	type Source struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Info string `automap:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Field_Value_Pointer(t *testing.T) {

	type Source struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Info *string `automap:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != *target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Field_Pointer_Value(t *testing.T) {

	type Source struct {
		Info *string `automap:"info"`
	}

	type Target struct {
		Info string `automap:"info"`
	}

	source := Source{
		Info: sp("Source.Info"),
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if *source.Info != target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Field_Pointer_Pointer(t *testing.T) {

	type Source struct {
		Info *string `automap:"info"`
	}

	type Target struct {
		Info *string `automap:"info"`
	}

	source := Source{
		Info: sp("Source.Info"),
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if *source.Info != *target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Field_NilPointer_Value(t *testing.T) {

	type Source struct {
		Info *string `automap:"info"`
	}

	type Target struct {
		Info string `automap:"info"`
	}

	source := Source{}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != nil {
		fmt.Println("AutoMap should not set source values.")
		t.Fail()
	}
	if target.Info != "" {
		fmt.Println("Unexpected value set.")
		t.Fail()
	}
}

func Test_AutoMap_Field_NilPointer_Pointer(t *testing.T) {

	type Source struct {
		Info *string `automap:"info"`
	}

	type Target struct {
		Info *string `automap:"info"`
	}

	source := Source{}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != nil {
		fmt.Println("AutoMap should not set source values.")
		t.Fail()
	}
	if target.Info != nil {
		fmt.Println("Unexpected value set.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_Value_Value(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested Bar `automap:"nested"`
	}

	foo := Foo{
		Info: "Hello from Foo!",
	}

	source := Source{
		Nested: foo,
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_Value_Pointer(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested *Bar `automap:"nested"`
	}

	foo := Foo{
		Info: "Hello from Foo!",
	}

	source := Source{
		Nested: foo,
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_Pointer_Value(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested *Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested Bar `automap:"nested"`
	}

	foo := Foo{
		Info: "Hello from Foo!",
	}

	source := Source{
		Nested: &foo,
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_Pointer_Pointer(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested *Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested *Bar `automap:"nested"`
	}

	foo := Foo{
		Info: "Hello from Foo!",
	}

	source := Source{
		Nested: &foo,
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_NilPointer_Value(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested *Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested Bar `automap:"nested"`
	}

	source := Source{}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested != nil {
		fmt.Println("AutoMap should not set source values.")
		t.Fail()
	}

	empty := Bar{}
	if target.Nested != empty {
		fmt.Println("Unexpected value set.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_NilPointer_Pointer(t *testing.T) {

	type Foo struct {
		Info string `automap:"info"`
	}

	type Source struct {
		Nested *Foo `automap:"nested"`
	}

	type Bar struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Nested *Bar `automap:"nested"`
	}

	source := Source{}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Nested != nil {
		fmt.Println("AutoMap should not set source values.")
		t.Fail()
	}

	if target.Nested != nil {
		fmt.Println("Unexpected value set.")
		t.Fail()
	}
}

func Test_AutoMap_Struct_Recursive(t *testing.T) {

	type Source struct {
		Nested *Source `automap:"nested"`
		Info   string  `automap:"info"`
	}

	type Target struct {
		Nested *Target `automap:"nested"`
		Info   string  `automap:"info"`
	}

	bottom := Source{
		Nested: nil,
		Info:   "Bottom",
	}

	middle := Source{
		Nested: &bottom,
		Info:   "Middle",
	}

	source := Source{
		Nested: &middle,
		Info:   "Top",
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
	if source.Nested.Info != target.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
	if source.Nested.Nested.Info != target.Nested.Nested.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}
