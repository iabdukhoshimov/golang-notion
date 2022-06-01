package main

func Swap() []int {
	a, b := 15, 20
	b, a = a, b
	return []int{a, b}
}
