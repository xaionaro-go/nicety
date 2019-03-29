package main

import (
	"fmt"
)

type slice []int

func (s slice) WrongReset() {
	s = s[:0]
}

func (s *slice) Reset() {
	(*s) = (*s)[:0]
}

func main() {
	s := slice{1, 2, 3}
	s.WrongReset()
	fmt.Println(len(s))
	s.Reset()
	fmt.Println(len(s))
}
