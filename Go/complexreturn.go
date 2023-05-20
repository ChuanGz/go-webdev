package main

import (
	"fmt"
	"math"
)

func multipleSquare(list [4]int) (float64, float64, float64, float64) {
	return math.Sqrt(float64(list[0])), math.Sqrt(float64(list[1])), math.Sqrt(float64(list[2])), math.Sqrt(float64(list[3]))
}

func main() {
	var myinput [4]int
	myinput = [4]int{10, 20, 30, 40}
	fmt.Println(multipleSquare(myinput))
}
