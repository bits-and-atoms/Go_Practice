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
}
