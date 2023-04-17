package types

import (
	"fmt"
)

func ExampleNewUpdate() {

	type User struct {
		ID   string  `update:"id" json:"id" binding:"required"`
		Name string  `update:"name" json:"name"`
		Mail *string `update:"mail" json:"mail,omitempty"`
	}

	user := User{
		ID:   "1337",
		Name: "Jeff",
		Mail: nil,
	}

	update, err := NewUpdate(user)
	if err != nil {
		// Handle error...
	}

	for k, v := range update {
		fmt.Println(k, v)
	}
	// Output:
	// id 1337
	// name Jeff
}
