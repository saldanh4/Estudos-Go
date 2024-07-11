package main

import "fmt"

func stringArrays(text string) {

	//Definindo uma array de string
	var strArray = [5]string{"Ana", "Bruno", "Pedro"}

	//Exibindo os valores do array
	fmt.Println("Valores do Array antes da alteração aplicada com base no parâmetro\n", strArray)

	//Atribuindo valores ao array
	strArray[3] = "Maria"
	strArray[4] = "Julia"

	//atribuindo o valor recebido na função para a posição 0 do array
	strArray[0] = text

	//exibindo os valores do array e seus respectivos indices
	for i, j := range strArray {
		fmt.Printf("\nPosição %v - %v", i, j)
	}

}
