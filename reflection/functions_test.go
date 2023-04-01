package reflection

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func ExampleGetTypeName() {

	foo := "bar"
	result := GetTypeName(foo)
	fmt.Println("1: foo:", result)

	// Output:
	// 1: foo: string
}

func ExampleGetFunctionName() {

	// This code runs inside 'func ExampleGetFunctionName()'
	fmt.Println("1: Result:", GetFunctionName())
	// Output:
	// 1: Result: ExampleGetFunctionName

}

func ExampleMapTagToType() {

	type User struct {
		ID       string    `json:"id" xml:"x_id"`
		Name     string    `json:"name,omitempty" bson:"b_name"`
		Age      int       `json:"age" yaml:"y_age"`
		Created  time.Time `json:"created" firestore:"created_time"`
		Active   bool      `json:"active,omitempty"`
		Password string    `json:"-"`
		Email    string
	}

	m := MapTagToType("json", User{})

	fmt.Printf("[Field:%v] -> [Type:%v]\n", "id", m["id"])
	fmt.Printf("[Field:%v] -> [Type:%v]\n", "name", m["name"])
	fmt.Printf("[Field:%v] -> [Type:%v]\n", "age", m["age"])
	fmt.Printf("[Field:%v] -> [Type:%v]\n", "created", m["created"])
	fmt.Printf("[Field:%v] -> [Type:%v]\n", "active", m["active"])

	_, found := m["password"]
	fmt.Printf("[Field:Password] -> [Found:%v]\n", found)
	_, found = m["email"]
	fmt.Printf("[Field:Email] -> [Found:%v]\n", found)

	// Output:
	// [Field:id] -> [Type:string]
	// [Field:name] -> [Type:string]
	// [Field:age] -> [Type:int]
	// [Field:created] -> [Type:Time]
	// [Field:active] -> [Type:bool]
	// [Field:Password] -> [Found:false]
	// [Field:Email] -> [Found:false]
}

func Test_AutoMapper_Value_And_Pointer(t *testing.T) {

	sp := func(input string) *string {
		return &input
	}

	var sourceValue = struct {
		Name string `mapper:"name"`
	}{
		Name: "Name Value",
	}

	var sourcePointer = struct {
		Name *string `mapper:"name"`
	}{
		Name: sp("Name Pointer"),
	}

	var targetValue = struct {
		Name string `mapper:"name"`
	}{}

	var targetPointer = struct {
		Name *string `mapper:"name"`
	}{}

	err := AutoMapper(sourceValue, &targetPointer)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = AutoMapper(sourceValue, &targetValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = AutoMapper(sourcePointer, &targetValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = AutoMapper(sourcePointer, &targetPointer)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

type Source struct {
	Name *string `mapper:"name"`
	//Age    int    `mapper:"age"`
	//Active *bool  `mapper:"active"`
	//Nested NestedSource `mapper:"nested"`
}

type NestedSource struct {
	Secret string `mapper:"secret"`
}

type Target struct {
	//Age    *int    `mapper:"age"`
	//Active bool    `mapper:"active"`
	Name *string `mapper:"name"`
	//Nested NestedTarget `mapper:"nested"`
}

type NestedTarget struct {
	Secret string `mapper:"secret"`
}

func AutoMapper(source, target any) error {

	// If multiple fields have the same tag name, then only the index of the last field will be included in the map.
	targetMap := mapTagToFieldIndex("mapper", target)
	sourceMap := mapTagToFieldIndex("mapper", source)

	// Peel off pointer and interface wrapper from 'any'.
	targetField := reflect.ValueOf(&target).Elem().Elem()
	sourceField := reflect.ValueOf(&source).Elem().Elem()
	describe("1.1: targetField:", targetField)
	describe("1.2: sourceField:", sourceField)

	// Check that target struct is passed by pointer.
	if targetField.Kind() != reflect.Pointer {
		return fmt.Errorf("target must be passed by pointer")
	}

	// Peel off pointer and check that it is a struct.
	targetField = targetField.Elem()
	describe("1.3: targetField:", targetField)
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

		// Peel off pointer and interface wrapper from 'any', then pick the struct field.
		sourceField = sourceField.Field(sourceIndex)
		describe("2.2: sourceField:", sourceField)

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
	if t.Kind() == reflect.Pointer {
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
