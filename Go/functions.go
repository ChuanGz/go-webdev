package main

import "fmt"

const Truth = true

func add(x, y, z int, f1, f2 float32) float32 {
	const PI = 3.14
	return float32(x+y+z) + f1 + f2
}

func main1() {
	fmt.Println(add(2, 2, 6, 3.2, 0.7))
}
