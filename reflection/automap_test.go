package reflection

import (
	"fmt"
	"testing"
	"time"
)

func sp(input string) *string {
	return &input
}

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

func Test_AutoMap_Struct_Nested(t *testing.T) {

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
		NestedOnce Bar    `automap:"once"`
		Info       string `automap:"info"`
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
	}
	if f.Info != a.Info {
		t.Fail()
	}
	if f.NestedOnce.Info != a.NestedOnce.Info {
		t.Fail()
	}
	if f.NestedOnce.NestedTwice.Info != a.NestedOnce.NestedTwice.Info {
		t.Fail()
	}
}

func Test_AutoMap_Error_Bad_Input(t *testing.T) {

	type ValidSource struct {
		Info string `automap:"info"`
	}

	type ValidTarget struct {
		Info string `automap:"info"`
	}

	validSource := ValidSource{Info: "I'm very valid!"}

	var err error
	_, err = AutoMap[ValidTarget](1337)
	if err == nil {
		fmt.Println("AutoMap should not allow non-struct source input.")
		t.Fail()
	}

	_, err = AutoMap[string](validSource)
	if err == nil {
		fmt.Println("AutoMap should not allow non-struct type parameters.")
		t.Fail()
	}

	_, err = AutoMap[*ValidTarget](validSource)
	if err == nil {
		fmt.Println("AutoMap should not allow pointer type parameters.")
		t.Fail()
	}

}

func Test_AutoMap_Equal_Struct_Types(t *testing.T) {

	type Source struct {
		Timestamp time.Time `automap:"timestamp"`
	}

	type Target struct {
		Timestamp time.Time `automap:"timestamp"`
	}

	source := Source{time.Now()}

	var err error
	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Timestamp != target.Timestamp {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}
