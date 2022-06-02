package main

import "fmt"

func Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 4

	fmt.Println("map:", m)
}
