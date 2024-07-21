package main

import (
	"fmt"
)

// Criando struct
type Sayajin struct {
	id   int
	name string
}

func returnEmployee(sayajin *Sayajin) {
	(sayajin).id = 1
	(sayajin).name = "Goku"
}

func printEmployee(sayajin Sayajin) {
	sayajin.id = 2
	sayajin.name = "Vegeta"
	fmt.Println("ID: ", sayajin.id)
	fmt.Println("Name: ", sayajin.name)

}

func main() {

	var sayajin, sayajin2 Sayajin

	returnEmployee(&sayajin)

	fmt.Println("ID: ", sayajin.id)
	fmt.Println("Name: ", sayajin.name)

	printEmployee(sayajin2)
}
