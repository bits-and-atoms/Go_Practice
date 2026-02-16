package main

import "fmt"

func transform(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
	//this is anonymous function which is returned by transform and it has access to the variable factor which is defined in the outer scope of transform
}

func ops(arr *[]int, whattodo func(int) int) []int {
	newA := []int{}
	for _, v := range *arr {
		newA = append(newA, whattodo(v))
	}
	return newA
}
func main() {
	arr := []int{1, 2, 3}
	arr2 := ops(&arr, transform(2))
	arr3 := ops(&arr2, transform(3))
	fmt.Println(arr, arr2, arr3)
}
