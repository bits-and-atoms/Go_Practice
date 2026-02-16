package main

import (
	"fmt"
	"os"
)

func rr(arr *[]int, i int, done chan bool, errchan chan error) {
	data, err := os.ReadFile("notes.txt")
	//change notes.txt to note.txt to see error
	_ = data
	if err != nil {
		errchan <- err
		return
	}
	fmt.Println("element is ", (*arr)[i])
	done <- true
}
func main() {
	arr := []int{1, 2, 3}
	dones := make([]chan bool, len(arr))
	errs := make([]chan error, len(arr))
	for i, _ := range arr {
		dones[i] = make(chan bool)
		errs[i] = make(chan error)
		go rr(&arr, i, dones[i], errs[i])
	}
	for i := range arr {
		select {
		case <-dones[i]:
			fmt.Println("done")
		case err := <-errs[i]:
			fmt.Println("error is ", err)
		}
	}
}
