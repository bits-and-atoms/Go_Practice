package main

import "fmt"

func ops(arr *[]int, whattodo func(int) int) []int {
	newA := []int{}
	for _, v := range *arr {
		newA = append(newA, whattodo(v))
	}
	return newA
}
func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	arr2 := ops(&arr, func(i int) int {
		return i * 2
	})
	//here i didnt need to write a double function separately i wrote it inplace and its called anonymous function because it has no name and its only used once and its not reusable
	fmt.Println(arr2)
}
