package main

import (
	"fmt"
)

func add(x int, y int, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(add(2, 2, 6))
}
