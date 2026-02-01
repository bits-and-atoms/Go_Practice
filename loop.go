package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		//infinite loop
		x := 0
		fmt.Scan(&x)
		if x == 0 {
			break
		} else if x == 1 {
			println("hjello")
		} else {
			continue
		}
	}
	println("hi")
}
