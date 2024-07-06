package main

import (
	"fmt"
)

func main() {
	//Declarando variável no contexto da função main
	//var number int = 3

	//Declarando variável dentro do contexto IF
	if number := 30; number%2 == 0 {
		fmt.Println("O numero", number, " é par")
	} else {
		fmt.Println("O numero ", number, " é impar")
	}
}
