package main

import (
	"fmt"
)

func main() {
	//Declarando variável no contexto da função main
	var number int = 3

	if number%2 == 0 {
		fmt.Println("O número ", number, " é par")
	} else {
		fmt.Println("O número ", number, " é impar")
	}

	//Declarando variável dentro do contexto IF
	if number1 := 30; number1%2 == 0 {
		fmt.Println("O numero", number1, " é par")
	} else {
		fmt.Println("O numero ", number1, " é impar")
	}

	/*
		Notei que o código aceita duas variáveis com o mesmo nome ao Declarar a mesma no IF
		Isso é normal dentro da linguagem? No C# teria acusado erro de sintaxe.

		Fiz os dois IFs apenas para ver o comportamento
	*/
}
