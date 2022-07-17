package main

import (
	"fmt"
	"test/HW-Test/do"
)

func main() {
	resoult, error := do.Do("a", 1, true)
	fmt.Println(resoult, error)
}
