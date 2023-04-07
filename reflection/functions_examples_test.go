package reflection

import (
	"fmt"
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

func ExampleAutoMap() {

	type Alpha struct {
		ThisName string `automap:"name"`
	}
	type Beta struct {
		ThatName string `automap:"name"`
	}
	a := Alpha{"Jeff"}
	b, err := AutoMap[Beta](a)
	if err != nil {
		// Handle err
	}
	fmt.Println(b.ThatName)

	// Output:
	// Jeff
}
