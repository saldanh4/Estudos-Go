package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(fibonacci(x))
	}
}

func fibonacci(x int) int {
	if x <= 1 {
		return x
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}
