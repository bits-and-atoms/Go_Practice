package main

import "fmt"

func sum(x int, y ...int) int {
	//... is used to indicate that y is a variadic parameter which can take any number of arguments and it will be treated as a slice of int
	tot := x
	for _, v := range y {
		tot += v
	}
	return tot
}
func main() {
	fmt.Println(sum(0, 1, 2, 3))
	//so it basically does x+all element in y

	// but what if you want to pass a slice to the variadic parameter
	arr := []int{1, 2, 3}
	fmt.Println(sum(0, arr...))
	//so you can use ... to unpack the slice and pass it as individual arguments to the variadic parameter
}
