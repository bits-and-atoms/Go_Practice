package main

import "fmt"

func main() {
	// := expects a type and value in rhs so
	//x := []string is wrong as []string is only type but []string{} is correct as it has type and value which is empty slice of string
	//x := []string{} is an empty slice and to add elements to it we can use append function
	x := []string{}
	x = append(x, "subham")
	x = append(x, "sai")
	println(x[0])
	println(x[1])
	// initially go allocates a slice with some capacity and length 0, when we append elements to it, it automatically resizes the underlying array if needed.
	//So we can keep appending elements to the slice without worrying about its capacity.
	y := make([]string, 2)
	y[0] = "subham"
	y[1] = "sai"
	y = append(y, "kumar") //y[2] = "kumar" will give error as y is of length 2 and we cannot assign value to index 2,
	// but when we append "kumar" to it, it automatically resizes the underlying array to accommodate the new element.
	fmt.Println(y)
	println(y[0])
	println(y[1])
	println(y[2])
	//here make accepts three parameters: type of elements in slice, length of slice and capacity of slice.
	//y is of 2 length and capacity of 2, when we append "kumar" to it, it automatically resizes the underlying array to accommodate the new element.

	z := make([]string, 2, 5)
	z[0] = "subham"
	z[1] = "sai"
	// z[2] = "kumar"  will give error as z is of length 2 and we cannot assign value to index 2,
	z = append(z, "kumar")
	// no new allocation as still capacity left
	fmt.Println(z)
	//here z is of 2 length and 5 capacity so till 5 elements go will not reallocate new memory rather it will wait till len becomes 5
	//after that it will reallocate new memory
	//to accommodate new elements.

	t := map[string]int{} // t is an empty map so go has to reallocate if more keys are inserted
	t["sam"] = 21
	tt := make(map[string]int, 33) // here 33 is hint , meaning it is allocating 33 size after that it will reallocate
	tt["ram"] = 25
}
