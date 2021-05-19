package main

import (
	"log"
	"sort"
)

func main(){
	var mySlice []int
	number := []int{1,23,44}

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 35)
	sort.Ints(mySlice)
	log.Println(mySlice,number)
}