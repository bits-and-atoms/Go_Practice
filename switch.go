package main

import "fmt"

func main() {
	x := 0
	fmt.Scan(&x)
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}
	fmt.Println(x)
}
