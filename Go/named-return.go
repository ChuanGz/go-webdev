package main

import (
	"fmt"
)

func Inductive(perfected int) (x, y, z int) {
	for i := 1; i < perfected; i++ {
		if perfected%i == 0 {
			x = i
		}
	}
	for i := 1; i < perfected; i++ {
		if perfected%i == 0 && x != i {
			z = i
		}
	}
	for j := perfected; j > 0; j-- {
		if perfected%j == 0 {
			y = j
		}
	}
	return
}

func main3() {
	fmt.Println(Inductive(6))
	fmt.Println(Inductive(36))
	fmt.Println(Inductive(16))
	fmt.Println(Inductive(1388))
	fmt.Println(694 * 347)
}
