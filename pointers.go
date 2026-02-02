package main

import "fmt"

func ChangeAge(t *int) {
	*t = *t + 5
}
func main() {
	age := 32
	fmt.Println("My age is", age)
	ChangeAge(&age)
	fmt.Println("My new age is", age)
}
