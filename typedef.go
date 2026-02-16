package main

import "fmt"

type mp map[[2]int]string

func (t mp) output() {
	fmt.Println(t)
}
func main() {
	x := make(mp, 23)
	x[[2]int{1, 2}] = "subham"
	x[[2]int{3, 4}] = "sai"
	x.output()
}
