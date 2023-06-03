package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("weekday", time.Now().Weekday())
	defer fmt.Println("abc")
	fmt.Print("Hi 2, ")
	defer fmt.Println("now", time.Now())
	fmt.Print("Hi 1, ")
}
