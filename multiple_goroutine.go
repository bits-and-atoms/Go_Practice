package main

import (
	"fmt"
	"time"
)

func a1(done chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("a1")
	done <- true
}
func a2(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("a2")
	done <- true
}
func a3(done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("a3")
	done <- true
}
func main() {
	donechan := make([]chan bool, 3)
	for i := 0; i < 3; i++ {
		donechan[i] = make(chan bool)
		if i == 0 {
			go a1(donechan[i])
		} else if i == 1 {
			go a2(donechan[i])
		} else {
			go a3(donechan[i])
		}
	}
	for _, v := range donechan {
		<-v
	}
}

//this code runs for 3sec not 6 as all are running in parallel and main is waiting for all of them to finish and then it exits
