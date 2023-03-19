package setup

import (
	"fmt"
	"os"
)

func ExampleValidateEnvironment() {

	_ = os.Setenv("ALPHA", "foo")

	expected := []string{
		"ALPHA",
	}

	err := ValidateEnvironment(expected)
	fmt.Println("1: Error:", err)

	expected = append(expected, "BETA", "GAMMA")

	err = ValidateEnvironment(expected)
	fmt.Println("2: Error:", err)

	// Output:
	// 1: Error: <nil>
	// 2: Error: missing environment variables: BETA, GAMMA
}
