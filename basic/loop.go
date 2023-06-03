package main

import "fmt"

func main() {
	// sum100()
	forcontinue()
}

func forcontinue() {
	sum := 1
	for sum < 20 { // while of Go
		sum += 1
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

func sum100() {
	sum := 0
	for i := -10; i < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
