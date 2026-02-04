package main

import "fmt"

func main() {
	var a [10]int = [10]int{4}
	p := [4]float64{1.0, 2.0, 3.0}
	a[2] = 100
	fmt.Println(a)
	fmt.Println(p)
	sl := a[1:6]
	fmt.Println(sl)
	//sl will have a from index 1 to 5
	sl2 := a[1:]
	//sl2 will have a from index 1 to end
	sl3 := a[:6]
	//sl3 will have a from start to index 5
	fmt.Println(sl2)
	fmt.Println(sl3)
	sl[0] = 500
	//0 of sl is 1 of a
	fmt.Println(a)
	//so slices are just pointer to part of an array they are not copy, if you change in slice it is reflected in array
	fmt.Println(len(sl), cap(sl))
	// length is how many elements are in slice, i.e. ind 1 to 5
	// capacity is how many elements can be stored from starting index to end of array i.e. 1 to 9,
	// keep appending elements to slice until capacity is reached, after that a new array will be created and slice will point to that new array
}
