package main

import "log"

func main()  {
	var myString string
	myString = "Green"
	log.Println("my color is",myString)
	changeUsingPointer(&myString)
	log.Println("after change pointer",myString)
}

func changeUsingPointer(s *string){
	log.Println("s is ",s)

	newValue := "red"
	*s = newValue
}