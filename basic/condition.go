package main

import (
	"fmt"
	"time"
)

func main1() {
	//	ifelse()
	switchf()
}

func ifelse() {
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

func switchf() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	currWD := time.Now().Weekday()
	fmt.Println("today is", currWD)
	switch currWD {
	case time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is weekday")
	}

	currTime := time.Now()
	fmt.Println("now is", currTime)
	switch {
	case currTime.Hour() < 12:
		fmt.Println("before noon")
	case currTime.Hour() >= 12:
		fmt.Println("after noon")
	}

	foundType := func(i interface{}) {
		switch ft := i.(type) {
		case bool:
			fmt.Println(ft, "is bool")
		case int:
			fmt.Println(ft, "is int")
		case string:
			fmt.Println(ft, "is string")
		default:
			fmt.Printf("unknow type %T\n", ft)
		}
	}

	foundType(123)
	foundType("123")
	foundType(true)
	foundType(false)
	foundType(2022 - 1 - 1)
	foundType(currWD)   // time.Weekday
	foundType(currTime) // time.Time
}
