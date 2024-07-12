package main

import (
	"fmt"
)

func main() {

	//declarando slice de inteiros
	intSlice := []int{1, 2, 5}
	fmt.Println("\nSlice 1")
	fmt.Printf("Tamanho: %d\n", len(intSlice))
	fmt.Printf("Capacidade: %d\n", cap(intSlice))

	//atribuindo novos valores ao slice
	intSlice = append(intSlice, 3, 4)

	//exibindo tamanho e capacidade após atribuir dois novos valores ao slice
	//a capacidade do slice sempre dobra em relação à capacidade original ao inserirmos novos itens mesmo que a quantidade não seja o dobro?
	fmt.Printf("Tamanho: %d\n", len(intSlice))
	fmt.Printf("Capacidade: %d\n", cap(intSlice))

	//criando slice com a função make e atribuindo o tamanho
	intSlice2 := make([]int, 3)
	fmt.Println("\nSlice 2")
	//atribuindo os valores do intSlice no intSlice2, exibindo o slice, tamanho e capacidade
	intSlice2 = append(intSlice2, intSlice...)
	fmt.Println(intSlice2)
	fmt.Printf("Tamanho: %d\n", len(intSlice2))
	fmt.Printf("Capacidade: %d\n", cap(intSlice2))

	//criando intSlice3 sem tamanho definido
	intSlice3 := []int{}
	//exibindo o slice, tamanho e capacidade
	fmt.Println(intSlice3)
	fmt.Println("\nSlice 3")
	fmt.Printf("Tamanho: %d\n", len(intSlice3))
	fmt.Printf("Capacidade: %d\n", cap(intSlice3))
	//atribuindo valores ao intSlice3, exibindo novos valores, tamanho e capacidade
	intSlice3 = append(intSlice3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(intSlice3)
	fmt.Printf("Tamanho: %d\n", len(intSlice3))
	fmt.Printf("Capacidade: %d\n", cap(intSlice3))
	//exibindo apenas os 3 ultimos numeros da intSlice3
	fmt.Println(intSlice3[7:10])

}
