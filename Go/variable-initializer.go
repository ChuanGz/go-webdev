package main

import (
	"fmt"
	"time"
)

var globCheck bool = true
var globAuthName string = "ChuanGz"
var globDate time.Time = time.Now()

func main4() {
	var pi float64 = 2.33445566
	fmt.Println("function level")
	fmt.Println(pi)
	fmt.Println(globCheck)
	fmt.Println(globAuthName)
	fmt.Println(globDate)
}
