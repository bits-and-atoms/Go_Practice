package main

import "fmt"

func main() {
	pt := []int{}
	for i := 0; i < 10; i++ {
		pt = append(pt, i*10)
		//here pt was a nil slice initially, when we append first element a new array is created and pt points to it
		//when we keep appending elements, once capacity is reached a new array is created with double capacity and elements are copied
		//and pt points to that new array
		// here you have to assign the result of append back to pt, because append may return a new slice if the underlying array changes
	}
	fmt.Println(pt)
	//to remove an element from start do
	pt = pt[1:]
	fmt.Println(pt)
	//here also assign it back to pt because pt is now pointing to a new slice that starts from index 1 of the original slice

	temp := pt[2:]
	temp[2] = 999
	fmt.Println(pt)
	//this changes pt as well, what if we need a copy not reference
	//do this
	temp = append([]int(nil), temp...)
	temp[2] = 555
	fmt.Println(pt)
	fmt.Println(temp)
}
