package main

import (
	"fmt"
	"math"
)

func main() {
	var x, z float64 = 7, 9
	const pi = 3.1415926
	fmt.Scan(&x)
	//take x as input from user
	y := 10.05 //:= is shorthand here you can specify type it auto selects type
	t := math.Pow(x, y) * z
	fmt.Println(t)
	fmt.Printf("%.2f\n", t)
	fmt.Println(`this is first line
this is new line here no need to do \n as backticks took care of it`)
}
