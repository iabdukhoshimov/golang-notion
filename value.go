package main

import "fmt"

func Value() {
	var a_1 string = "hello world"
	var a_2 int = 15
	var a_3 float32 = 1.2
	var isDead bool = false

	fmt.Printf("%v it is boolean data type in Go\n", isDead)
	fmt.Printf("%v it is float data type in Go\n", a_3)
	fmt.Printf("%v it is a integer data type in Go \n", a_2)
	fmt.Printf("%v it is a string data type in Go\n", a_1)
}
