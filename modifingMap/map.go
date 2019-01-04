package main

// $ go run map.go
// map[oldValue:1 newValue:2]

import (
	"fmt"
)

func modifyMap(m map[string]int) {
	m["oldValue"] = 1
	m["newValue"] = 2
}

func main() {
	m := map[string]int{}
	m["oldValue"] = 0
	modifyMap(m)
	fmt.Println(m)
}
