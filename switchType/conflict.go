package main

import (
	"fmt"
)

type someInterface interface {}
type someType int

func newSomeType() someInterface {
	return someType(0)
}

func main() {
	switch newSomeType().(type) {
	case someInterface:
		fmt.Println("hey!")
	case someType:
		fmt.Println("won't happened")
	}

	switch newSomeType().(type) {
	case someType:
		fmt.Println("hey, again...")
	case someInterface:
		fmt.Println("won't happened")
	}

	fmt.Println("end")
}
