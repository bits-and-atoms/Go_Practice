package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstname  string    `json:"pehla_name"`
	lastname   string    `json:"aakhri_name"`
	birthdate  string    `json:"janam_tithi"`
	createdate time.Time `json:"create_date"`
}

//in above struct  we have defined tags for json serialization, anything inside `` will be used by some other package that expects those
//here i did json so json package while using these will name the response like "pehla_name" instead of "firstname"

// constructor function for user struct but not with validator
func NewUser(firstname, lastname, birthdate string, u *user) {
	u.firstname = firstname
	(*u).lastname = lastname
	//both syntax is correct
	u.birthdate = birthdate
	u.createdate = time.Now()
}

// constructor function with validator for user struct
func createUser(firstname, lastname, birthdate string) (*user, error) {
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("some fields are empty")
	}
	return &user{firstname: firstname, lastname: lastname, birthdate: birthdate, createdate: time.Now()}, nil
}

// a function attacked to struct is called method
func (u user) outputUser() {
	// u user is receiver type
	// it is pass by value not reference
	fmt.Println("User:", u.firstname, u.lastname, u.birthdate, u.createdate)
}
func (u *user) deleteUser() {
	// u *user is receiver type
	// it is pass by reference
	u.firstname = ""
	u.lastname = ""
	u.birthdate = ""
	u.createdate = time.Time{}
}

type Acc struct {
	accountNumber string
	u             user
}

func NewAcc(accountNumber string, u user) *Acc {
	return &Acc{accountNumber: accountNumber, u: u}
}
func main() {
	a := user{}
	NewUser("John", "Doe", "1990-01-01", &a)
	a.outputUser()
	//outputuser is method of struct user see we dont need to pass any parameter
	a.deleteUser()
	a.outputUser()
	b, err := createUser("subham", "majumder", "1995-05-05")
	if err != nil {
		fmt.Println(err)
	}
	b.outputUser()
	c := NewAcc("1234567890", *b)
	fmt.Println("Account Number:", c.accountNumber)
	c.u.outputUser()
	//nesting or embedding structs
}
