package main

import "fmt"

func Lifo() {
	var stack []string

	stack = append(stack, "world!")
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		//print top
		n := len(stack) - 1

		fmt.Println(stack[n])

		//Pop
		stack = stack[:n]
	}
	//output Hello world!
}
