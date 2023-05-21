package main

import "fmt"

func main() {

	var defaultS int64 = 13
	if 13 == defaultS {
		fmt.Println("13 equal defaultS")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	number := 200

	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is old")
	}
}
