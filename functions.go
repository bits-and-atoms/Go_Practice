package main

import "fmt"

func calc(a int, b int) (int, int) {
	return a + b, a - b
}
func modify(x *int) {
	*x = 100 // modifies the original
}
func main() {
	t1, t2 := calc(5, 10)
	fmt.Println(t1, t2)
	modify(&t1)
	fmt.Println(t1)
}
