package main

import "fmt"

type bucket map[[2]int]string

func (t bucket) output() {
	fmt.Println(t)
}
func main() {
	x := make(bucket, 23)
	x[[2]int{1, 2}] = "subham"
	x[[2]int{3, 4}] = "sai"
	x.output()

	for k, v := range x {
		fmt.Println("key", k, "value", v)
	}
	t := make([]string, 5)
	t[2] = "subham"
	t[3] = "sai"
	for i, v := range t {
		fmt.Println("index", i, "value", v)
	}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}
