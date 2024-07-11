package variables

import (
	"fmt"
)

func variables() {
	fmt.Println("Definição de Variáveis")
	//Definição de Variáveis
	var number int = 10
	var floatNumber float64 = 25.5
	var text string = "André"
	var booleano bool = false
	const LIGHTSPEED uint32 = 299792458

	//Exibindo os valores das variáveis
	fmt.Println("Inteiro: ", number)
	fmt.Println("Decimal: ", floatNumber)
	fmt.Println("Texto: ", text)
	fmt.Println("Booleano: ", booleano)
	fmt.Println("Velocidade da Luz em km/s:", LIGHTSPEED)
}
