package main

import (
	"fmt"
)

func main() {
	//declaração das variáveis
	texto := "White"
	var num int = 18

	//definição do switch baseado em string, requer a variável/valor esperado após o "switch"
	switch texto {
	case "BLUE":
		fmt.Println(texto)
	case "RED":
		fmt.Println(texto)
	default:
		fmt.Println("O texto não corresponde ao esperado\n")
	}

	//Aqui demorei a entender que não precisava informar a variável após o "switch" assim como é no caso da string
	//Isso ocorre por causa da condição dentro do case? Ou está relacionado à algo que deixei passar nos tutoriais da w3c?
	switch {
	//No primeiro case testo se o numero atende ambas as condições e nos seguintes trato de forma individual
	case num%2 == 0 && num%3 == 0:
		fmt.Printf("%v é divisível por 2 e por 3", num)
	case num%2 == 0:
		fmt.Printf("%v é divisível por 2", num)
	case num%3 == 0:
		fmt.Printf("%v é divisível por 3", num)
	default:
		fmt.Printf("%v não é divisível por 2 ou 3", num)
	}
}
