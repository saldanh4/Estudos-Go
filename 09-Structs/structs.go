package main

import (
	"fmt"
)

// Criando struct
type Sayajin struct {
	id   int
	name string
}

func returnEmployee(sayajin Sayajin) Sayajin {
	sayajin.id = 0001
	sayajin.name = "Goku"
	return sayajin
}

func printEmployee(sayajin Sayajin) {
	sayajin.id = 0002
	sayajin.name = "Vegeta"
	fmt.Println("ID: ", sayajin.id)
	fmt.Println("Name: ", sayajin.name)

}

func main() {

	var sayajin, sayajin2 Sayajin

	sayajin = returnEmployee(sayajin)

	fmt.Println("ID: ", sayajin.id)
	fmt.Println("Name: ", sayajin.name)

	printEmployee(sayajin2)
}
