package main

import "fmt"

//função soma
func soma(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

//função para dobrar o valor da variável
func dobraValor(num *int) {
	(*num) *= 2
}

func main() {

	//chama a função soma e exibe o retorno com o resultado
	println(soma(5, 7))

	//Declara variável
	var x int = 5
	//Ponteiro apontando para a variável
	//p := &x

	println(x)
	//chama a função que dobra o valor e exibe o resultado.
	dobraValor(&x)
	fmt.Println(x)

	//atribuí à x o valor de retorno da função subtrai
	x = subtrai(20, 5)
	println(x)

	//Como a função doubleValue não tem retorno, não consegui replicar o valor na variável x e exibir o valor atualizado.
	//Optei por usar uma variável nomeada apenas na segunda função
}

//defini a função subtrai após o main para ver se a ordem de declaração teria influência no momento de compilar e rodar o código.
func subtrai(num1 int, num2 int) int {
	subtrai := num1 - num2
	return subtrai
}
