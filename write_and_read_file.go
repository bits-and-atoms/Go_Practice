package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	amt := 0
	amt = 5
	fmt.Println("Amount is:", amt)
	val := fmt.Sprint(amt)
	os.WriteFile("amount.txt", []byte(val), 0644)
	fmt.Println("written to file successfully")
	data, _ := os.ReadFile("amount.txt") // dont want to use the err so did _
	temp := string(data)
	balance, _ := strconv.ParseInt(temp, 10, 64)
	fmt.Println("balance is:", balance)
}
