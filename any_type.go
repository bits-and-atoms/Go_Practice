package main

import "fmt"

func call(value interface{}) {
	//interfaces{} means any type is allowed
	switch value.(type) {
	case int64:
		fmt.Println(value.(int64))
	case int:
		fmt.Println(value.(int))
	case float64:
		fmt.Println(value.(float64))
	default:
		fmt.Println("unknown type")
	}
}
func anotherform(a, b interface{}) interface{} {
	ai, isai := a.(int)
	//ai holds value if isai is true which will be true if a is int
	bf, isbf := b.(float64)

	if isai && isbf {
		return float64(ai) + bf
		//returning float64 type
	}
	return "subham"
	//returning string type
}

// in above function i am accepting any type and returning any type
func main() {
	call(13)
	call("subham")
}
