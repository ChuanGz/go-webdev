package main

import "fmt"

func main() {
	sum100()
}

func sum100() {
	sum := 0
	for i := -10; i < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
