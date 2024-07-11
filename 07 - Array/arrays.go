package main

import "fmt"

func main() {
	var myArray [2]int

	myArray[0] = 10
	myArray[1] = 20

	fmt.Println(myArray)
	//aqui retorna erro Index 2 out of bounds por estar fora do tamanho do array
	//myArray[2] = 30

	for i, j := range myArray {
		fmt.Printf("Posição %d = %v\n", i, j)
	}
}
