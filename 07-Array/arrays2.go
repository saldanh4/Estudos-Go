package main

import "fmt"

func myStringArray() {
	var nomes [4]string

	nomes[0] = "Ana"
	nomes[1] = "Julia"
	nomes[2] = "José"
	nomes[3] = "Pedro"

	fmt.Println(nomes)
}

func main() {

	myStringArray()
	//nomes1 := [4]string{"João", "Maria", "Mario", "Alberto"}

	//nomes2 := [...]string{"Alberto", "Joaquim", "Zélia"}
}
