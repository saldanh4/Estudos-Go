package main

import (
	"fmt"
)

func main() {
	//Criando array com valores inicializados
	var myArray = [2]int{2, 3}

	//exibindo array
	fmt.Println(myArray)

	//Atribuindo novos valores
	myArray[0] = 10
	myArray[1] = 20

	//exibindo um elemento do array
	fmt.Println(myArray[0])

	//aqui retorna erro Index 2 out of bounds por estar fora do tamanho do array
	//myArray[2] = 30

	//exibindo a posição e respectivo valor no array usando o for e range
	for i, j := range myArray {
		fmt.Printf("Posição %d = %v\n", i, j)
	}
	//chamando a função stringArrays com parâmetro string
	stringArrays("André")
}
