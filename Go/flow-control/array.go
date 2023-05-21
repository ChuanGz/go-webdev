package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a) // [0 0 0 0 0]

	a[0] = 5
	a[2] = 15
	a[3] = 20
	a[4] = 25
	fmt.Println("set ", a)
	fmt.Println("get ", a[4])

	fmt.Println("len ", len(a))

	// intit with predefined
	b := [3]int{1, 2, 3}
	fmt.Println(b)

	// multi demesion
	c33 := [3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, -1}}
	fmt.Println(c33)
	c25 := [2][5]int{{1, 1, 1, -1, -1}, {0, 1, 2, 3, 4}}
	fmt.Println(c25)
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			c25[i][j] = -99 - i - j
		}
	}
	// changed
	fmt.Println(c25)
}
