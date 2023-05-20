package main

import "fmt"

func add(x, y, z int, f1, f2 float32) float32 {
	return float32(x+y+z) + f1 + f2
}

func main() {
	fmt.Println(add(2, 2, 6, 3.2, 0.7))
}
