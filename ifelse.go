package main

func main() {
	choice := 2 != 1
	if choice {
		println("Choice is true")
		return
	} else if choice == false {
		println("choice is false")
	} else {
		println("it should not come here")
	}
	println("helo")
}
