package main

import "fmt"

type op func(k int) int
type ch func(arr *[]int) op

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println(arr1, arr2)
	arr3 := opsonArr(&arr1, double)
	arr4 := opsonArr(&arr2, triple)
	fmt.Println(arr3, arr4)
	arr5 := opsonArr2(&arr1)
	arr6 := opsonArr2(&arr2)
	fmt.Println(arr5, arr6)
}
func double(k int) int {
	return 2 * k
}
func triple(k int) int {
	return 3 * k
}
func opsonArr(arr *[]int, whattoDo op) []int {
	newArr := []int{}
	for _, v := range *arr {
		newArr = append(newArr, whattoDo(v))
	}
	return newArr
}
func opsonArr2(arr *[]int) []int {
	newArr := []int{}
	for _, v := range *arr {
		newArr = append(newArr, chooseOps(arr)(v))
	}
	return newArr
}
func chooseOps(arr *[]int) op {
	if (*arr)[0] == 1 {
		return double
	} else {
		return triple
	}
}
