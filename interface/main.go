package main

import "log"

type Animal interface {
	Say() string
	Legs() int
}

type Cat struct {
	name  string
	breed string
}

func main() {
	garfield := Cat{
		name:  "grr",
		breed: "alyx",
	}

	printMe(&garfield)
}

func printMe(p Animal) {
	log.Println("help--", p.Say())
}

func (c *Cat) Say() string {
	return "Meow"
}
func (c *Cat) Legs() int {
	return 4
}
