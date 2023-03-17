package reflection

import (
	"fmt"
	"time"
)

func ExampleMapFieldToType() {

	type User struct {
		ID       string    `json:"id"`
		Name     string    `json:"name,omitempty"`
		Age      int       `json:"age"`
		Created  time.Time `json:"created"`
		Active   bool      `json:"active,omitempty"`
		Password string    `json:"-"`
		Email    string
	}

	m := MapFieldToType("json", User{})

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
