package types

import (
	"fmt"

	"github.com/pergamenum/go-consensus-standards/constants"
	"github.com/pergamenum/go-consensus-standards/reflection"
)

func ExampleQuery_Validate() {

	type User struct {
		ID int `json:"id"`
	}

	ttt := reflection.MapTagToType("json", User{})
	otb := constants.ValidRelationalOperators

	q := Query{
		Key:      "id",
		Operator: "EQ",
		// Note that value is any/string.
		Value: "1337",
	}

	err := q.Validate(ttt, otb)
	if err != nil {
		fmt.Println("1: Error:", err)
	} else {
		fmt.Printf("2: Value Type: %T", q.Value)
	}
	// Output:
	// 2: Value Type: int
}
