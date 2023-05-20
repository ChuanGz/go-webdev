package main

import (
	"fmt"
	"math"
)

func multipleSquare(list [4]int) (a float64, b float64, c float64, d float64) {
	a = math.Sqrt(float64(list[0]))
	b = math.Sqrt(float64(list[1]))
	c = math.Sqrt(float64(list[2]))
	d = math.Sqrt(float64(list[3]))
	return
}

func main2() {
	var myinput [4]int
	myinput = [4]int{10, 20, 30, 40}
	fmt.Println(multipleSquare(myinput))
}
