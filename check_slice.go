package main

import "fmt"

func Check() {
	r := []int{1, 3, 4}

	if len(r) == 0 {
		fmt.Println("it is empty")
	} else {
		fmt.Println("it is not empty")
	}

}
