package main

// $ go run map.go
// map[oldValue:2 newValue:3]

import (
	"fmt"
)

func modifyMap(m map[string]int) {
	m["oldValue"] = 2
	delete(m, "anotherOldValue")
	m["newValue"] = 3
}

func main() {
	m := map[string]int{}
	m["oldValue"] = 0
	m["anotherOldValue"] = 1
	modifyMap(m)
	fmt.Println(m)
}
