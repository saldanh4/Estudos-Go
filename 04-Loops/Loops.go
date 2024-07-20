package main

import (
	"fmt"
)

func main() {
	//Laço for para exibir os inteiros de à 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Laço for para exibir a letra y infinitamente
	for i := 1; i == 0; i++ {
		fmt.Print("y")
	}

	//Laço sem parâmetros. Legal que isso permite o uso do For como o While, declarando apenas a condicional no loop e a variável de controle fora.
	//Achei isso perigoso em uma aplicação, caso alguém crie o laço para uma iteração futura e esqueça de complementar ou deixar comentado
	for {
		fmt.Print("y")
	}
}
