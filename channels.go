package main

import "fmt"

func f1(done chan bool) {
	fmt.Println("hello world")
	done <- false
}
func main() {
	done := make(chan bool)
	go f1(done)
	<-done
	//here main waits for the go routine to exit
}
