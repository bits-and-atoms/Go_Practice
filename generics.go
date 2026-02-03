package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	return a + b
}
func main() {
	rel1 := add("hello, ", "world!")
	fmt.Println(rel1)
	rel := add(10, 20)
	fmt.Println(rel)
	rel2 := add(10.5, 20.3)
	fmt.Println(rel2)
}
