package main

import "fmt"

func f1() {
	fmt.Println("hello world")
}
func f2() {
	fmt.Println("wow")
}
func main() {
	//using the go keyword we are starting f1 and f2 as goroutines which will run concurrently with the main function and with each other
	//this program dont prints anything because the main function exits before the goroutines have a chance to execute and print their output
	go f1()
	go f2()
}
