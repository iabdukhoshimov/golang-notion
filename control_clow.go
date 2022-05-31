package main

import "fmt"

func Control() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	a := 8
	b := 4
	if a%b == 0 {
		fmt.Printf("%v is divisible by %v\n", a, b)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Printf("%v had 1 digit\n", num)
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
