package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func (m* User) getFirstName() string {
	return m.FirstName
}
 
func main(){
	user := User {
		FirstName: "hello",
		LastName: "World",
	}
	log.Println("Print",user.getFirstName())
}