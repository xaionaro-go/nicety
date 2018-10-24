package main

import (
	"fmt"
)

type c struct {
	wontChange int
	s          []int
}

func main() {
	a := []c{
		c{
			wontChange: 1,
			s:          []int{0, 1, 2},
		},
	}

	for _, e := range a {
		e.wontChange = e.wontChange + 1
		for idx, _ := range e.s {
			e.s[idx] = e.s[idx] + 1
		}
	}

	fmt.Println(a)
}
