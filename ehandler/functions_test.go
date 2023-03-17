package ehandler

import (
	"errors"
	"fmt"
)

func ExampleWrap() {

	errFirst := errors.New("first")
	errSecond := errors.New("second")
	errThird := errors.New("third")

	alpha := Wrap("cause", errFirst)
	fmt.Println("1: ", alpha)
	fmt.Println("2: ", errors.Is(alpha, errFirst))

	beta := Wrap(errSecond, alpha)
	fmt.Println("3: ", beta)
	fmt.Println("4: ", errors.Is(beta, errFirst))
	fmt.Println("5: ", errors.Is(beta, errSecond))

	gamma := Wrap(1234, errThird)
	fmt.Println("6: ", gamma)
	fmt.Println("7: ", errors.Is(gamma, errThird))

	//Output:
	// 1:  [cause] -> first
	// 2:  true
	// 3:  [second] -> [cause] -> first
	// 4:  true
	// 5:  false
	// 6:  third
	// 7:  true
}
