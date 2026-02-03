package main

import "fmt"

// using interfaces you can do polymorphism in Go
// which ever struct implements the methods of an interface will be considered to implement that interface
// if an interface has a,b,c methods then struct which uses this interface must implement all three methods a,b,c
type human struct {
	name string
}
type animal struct {
	name string
}

func (h human) Speak() string {
	return "Hello, my name is " + h.name
}
func (a animal) Speak() string {
	return "Roar! I am " + a.name
}
func (h human) Read() {
	// implementation of Read method for human
	fmt.Println("i can read")
}
func (h human) Write() {
	// implementation of Write method for human
	fmt.Println("i can write")
}
func (a animal) Read() {
	// implementation of Read method for animal
	fmt.Println("i cant read")
}
func (a animal) Write() {
	// implementation of Write method for animal
	fmt.Println("i cant write")
}

type reader interface {
	Read()
}
type writer interface {
	Write()
}
type operations interface {
	reader
	writer
	Speak() string
	//operations interface is embedding reader and writer interfaces and adding Speak method
}

func main() {
	h := human{name: "Alice"}
	a := animal{name: "Leo"}
	var o operations = h
	fmt.Println(o.Speak())
	o.Read()
	o.Write()
	o = a
	fmt.Println(o.Speak())
	o.Read()
	o.Write()
	// using o i am calling methods of operations interface implemented by human and animal structs
}
