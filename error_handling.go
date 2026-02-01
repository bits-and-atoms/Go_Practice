package main

import (
	"errors"
	"os"
	"strconv"
)

func readfile() (int64, error) {
	val, err := os.ReadFile("notexist.txt") //this file is not in repo intentionally to generate error
	if err != nil {
		return -1, errors.New("No such file or directory")
	}
	temp := string(val)
	balance, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		return -1, errors.New("Couldn't convert to int")
	}
	return balance, nil
}
func main() {
	balance, err := readfile()
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Balance is:", balance)
}
