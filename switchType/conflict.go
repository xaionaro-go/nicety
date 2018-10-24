package main

import (
	"fmt"
)

type someBadInterface interface {
	NonExistantMethod()
}
type someAnotherInterface interface {}
type someInterface interface {}
type someType int

func newSomeType() someInterface {
	return someType(0)
}

func main() {
	switch newSomeType().(type) {
	case someBadInterface:
		fmt.Println("won't happened")
	case someInterface:
		fmt.Println("hey!")
	case someType:
		fmt.Println("won't happened")
	}

	switch newSomeType().(type) {
	case someAnotherInterface:
		fmt.Println("hey, again")
	case someInterface:
		fmt.Println("won't happened")
	default:
		fmt.Println("won't happened")
	}

	switch newSomeType().(type) {
	case someType:
		fmt.Println("hey, once more...")
	case someInterface:
		fmt.Println("won't happened")
	}

	fmt.Println("end")
}
