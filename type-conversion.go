package main

import (
	"fmt"
	"math"
)

func main5() {
	var x, y int = 12, 16
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
