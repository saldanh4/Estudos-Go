package main

import "fmt"

//Definindo interface animal com os métodos a serem utilizados
type Animal interface {
	Comer(comida string) string
	Dormir(horas int) int
}

//Definindo struct para receber os dados do urso
type Urso struct {
	tipoDeUrso string
	peso       int
}

//Implementando método Comer para o tipo criado na struct retornando uma string
func (u Urso) Comer(comida string) string {
	return "- A comida favorita do " + u.tipoDeUrso + " é " + comida
}

//Implementando o método Dormir para o tipo criado na struct sem retorno
func (u Urso) Dormir(horas int) {
	fmt.Printf("- O %v dorme aproximadamente %v horas por dia", u.tipoDeUrso, horas)
}

//Função para chamada dos dados do animal criado
func DadosAnimal(u Urso, comida string, sono int) {
	//imprimindo dados do objeto Urso
	fmt.Println("Animal: ", u.tipoDeUrso)
	fmt.Println("Peso: ", u.peso)

	fmt.Println("\nCuriosidades: ")
	//chamada dos métodos da interface animal implementados pela struct
	fmt.Println(u.Comer(comida))
	u.Dormir(sono)
}

func main() {

	//declarando objeto e definindo os valores dos campos
	u := Urso{
		tipoDeUrso: "Urso Pardo",
		peso:       200,
	}

	//Criando variáveis que serão consumidas nos métodos da interface
	comida := "mel"
	sono := 22
	//chamando função para apresentar os dados do animal com o objeto e as variáveis como parâmetro
	DadosAnimal(u, comida, sono)
}
